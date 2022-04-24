package services

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	pigo "github.com/esimov/pigo/core"
)

var (
	cascade          []byte
	puplocCascade    []byte
	faceClassifier   *pigo.Pigo
	puplocClassifier *pigo.PuplocCascade
	flpcs            map[string][]*pigo.FlpCascade
	imgParams        *pigo.ImageParams
)

var (
	eyeCascades  = []string{"lp46", "lp44", "lp42", "lp38", "lp312"}
	mouthCascade = []string{"lp93", "lp84", "lp82", "lp81"}
)

func FindFaces(baseString string) [][][]int {
	results := clusterDetection(baseString)
	dets := make([][]int, len(results))

	for i := 0; i < len(results); i++ {
		dets[i] = append(dets[i], results[i].Row, results[i].Col, results[i].Scale, 0)

		puploc := &pigo.Puploc{
			Row:      results[i].Row - int(0.085*float32(results[i].Scale)),
			Col:      results[i].Col - int(0.185*float32(results[i].Scale)),
			Scale:    float32(results[i].Scale) * 0.4,
			Perturbs: 63,
		}
		leftEye := puplocClassifier.RunDetector(*puploc, *imgParams, 0.0, false)
		if leftEye.Row > 0 && leftEye.Col > 0 {
			dets[i] = append(dets[i], leftEye.Row, leftEye.Col, int(leftEye.Scale), 1)
		}

		puploc = &pigo.Puploc{
			Row:      results[i].Row - int(0.085*float32(results[i].Scale)),
			Col:      results[i].Col + int(0.185*float32(results[i].Scale)),
			Scale:    float32(results[i].Scale) * 0.4,
			Perturbs: 63,
		}

		rightEye := puplocClassifier.RunDetector(*puploc, *imgParams, 0.0, false)
		if rightEye.Row > 0 && rightEye.Col > 0 {
			dets[i] = append(dets[i], rightEye.Row, rightEye.Col, int(rightEye.Scale), 1)
		}
		log.Println(dets)

		for _, eye := range eyeCascades {
			for _, flpc := range flpcs[eye] {
				flp := flpc.GetLandmarkPoint(leftEye, rightEye, *imgParams, puploc.Perturbs, false)
				if flp.Row > 0 && flp.Col > 0 {
					dets[i] = append(dets[i], flp.Row, flp.Col, int(flp.Scale), 2)
				}

				flp = flpc.GetLandmarkPoint(leftEye, rightEye, *imgParams, puploc.Perturbs, true)
				if flp.Row > 0 && flp.Col > 0 {
					dets[i] = append(dets[i], flp.Row, flp.Col, int(flp.Scale), 2)
				}
			}
		}

		for _, mouth := range mouthCascade {
			for _, flpc := range flpcs[mouth] {
				flp := flpc.GetLandmarkPoint(leftEye, rightEye, *imgParams, puploc.Perturbs, false)
				if flp.Row > 0 && flp.Col > 0 {
					dets[i] = append(dets[i], flp.Row, flp.Col, int(flp.Scale), 2)
				}
			}
		}
		flp := flpcs["lp84"][0].GetLandmarkPoint(leftEye, rightEye, *imgParams, puploc.Perturbs, true)
		if flp.Row > 0 && flp.Col > 0 {
			dets[i] = append(dets[i], flp.Row, flp.Col, int(flp.Scale), 2)
		}
	}

	persons := make([][][]int, len(dets))

	batchSize := 4
	for personIndex := range dets {
		person := dets[personIndex]

		batches := make([][]int, 0, (len(person)+batchSize-1)/batchSize)

		for batchSize < len(person) {
			person, batches = person[batchSize:], append(batches, person[0:batchSize:batchSize])
		}
		batches = append(batches, person)

		persons[personIndex] = batches
	}

	return persons
}

func clusterDetection(baseString string) []pigo.Detection {
	pwd, _ := os.Getwd()

	imageBytes, err := DecodeBaseImageToBytes(baseString)
	if err != nil {
		log.Fatalf("DecodeBaseImageToBytes %v", err)
	}

	reader := bytes.NewReader(imageBytes)

	src, err := pigo.DecodeImage(reader)
	if err != nil {
		log.Fatalf("pigo.DecodeImage: %v", err)
	}

	pixels := pigo.RgbToGrayscale(src)
	cols, rows := src.Bounds().Max.X, src.Bounds().Max.Y

	imgParams = &pigo.ImageParams{
		Pixels: pixels,
		Rows:   rows,
		Cols:   cols,
		Dim:    cols,
	}

	cParams := pigo.CascadeParams{
		MinSize:     20,
		MaxSize:     1000,
		ShiftFactor: 0.1,
		ScaleFactor: 1.1,

		ImageParams: *imgParams,
	}

	if len(cascade) == 0 {
		cascade, err = ioutil.ReadFile(filepath.Join(pwd, "cascade/facefinder"))
		if err != nil {
			log.Fatalf("Error reading the cascade file: %v", err)
		}
		p := pigo.NewPigo()

		faceClassifier, err = p.Unpack(cascade)
		if err != nil {
			log.Fatalf("Error unpacking the cascade file: %s", err)
		}
	}

	if len(puplocCascade) == 0 {
		puplocCascade, err := ioutil.ReadFile(filepath.Join(pwd, "cascade/puploc"))
		if err != nil {
			log.Fatalf("Error reading the puploc cascade file: %s", err)
		}
		puplocClassifier, err = puplocClassifier.UnpackCascade(puplocCascade)
		if err != nil {
			log.Fatalf("Error unpacking the puploc cascade file: %s", err)
		}

		flpcs, err = puplocClassifier.ReadCascadeDir(filepath.Join(pwd, "cascade/lps"))
		if err != nil {
			log.Fatalf("Error unpacking the facial landmark detection cascades: %s", err)
		}
	}

	dets := faceClassifier.RunCascade(cParams, 0.0)

	dets = faceClassifier.ClusterDetections(dets, 0.0)
	log.Println("Go though cascade", dets)
	return dets
}
