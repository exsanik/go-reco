package services

import (
	"log"
	"reco/models"
	"time"

	"github.com/Kagami/go-face"
)

func FindThresholdUserID(rec *face.Recognizer, users []models.UserDescriptor, input [128]float32, tolerance float32) int {
	var (
		length     = len(users)
		categories = make([]int32, length)
		samples    = make([]face.Descriptor, length)
	)

	for i, f := range users {
		var byteDescriptor [512]byte
		copy(byteDescriptor[:], f.Descriptor[:512])

		descriptor := BytesToDescriptor(byteDescriptor)
		samples[i] = descriptor
		categories[i] = int32(f.ID)
	}

	rec.SetSamples(samples, categories)
	userID := rec.ClassifyThreshold(input, tolerance)

	return userID
}

func GetFaceDescriptor(rec *face.Recognizer, baseString string) *face.Descriptor {
	imageBytes, err := DecodeBaseImageToBytes(baseString)
	if err != nil {
		log.Fatalf("GetFaceDescriptor:DecodeBaseImageToBytes %v", err)
	}

	recognizeStartTime := time.Now()
	var faces, recognizeErr = rec.Recognize(imageBytes)
	log.Printf("recognize faces on by %s", time.Since(recognizeStartTime))

	length := len(faces)
	if length != 1 || recognizeErr != nil {
		log.Printf("Can't recognize face %d", length)

		return &face.Descriptor{}
	}

	return &faces[0].Descriptor
}
