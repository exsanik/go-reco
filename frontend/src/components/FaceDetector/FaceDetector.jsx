import { useCallback, useEffect, useRef, useState } from "react";
import Webcam from "react-webcam";

import { detectFace } from "src/api/detectFace";
import { TYPE_FACE } from "src/config";

import { PersonCanvas } from "src/components/PersonCanvas";

import s from "./styles.module.css";

const HIGHT = 384;
const WIDTH = 512;

export const FaceDetector = ({ onDetectedFaces }) => {
  const webcamRef = useRef();

  const [drawingCoords, setDrawingCoords] = useState([]);
  const [pause, setPause] = useState(false);

  const handleDetectFace = useCallback(async (photo) => {
    if (!photo) return;

    const base64 = photo.split(",", 2)[1];

    const resp = await detectFace({ photo: base64 });
    setDrawingCoords(resp?.data || []);

    if (resp?.data?.length === 0) {
      console.log("No faces detected");
    }
  }, []);

  const updatePhoto = useCallback(() => {
    if (webcamRef.current) {
      setPause(false);
      const imageSrc = webcamRef.current.getScreenshot();

      handleDetectFace(imageSrc);
    }
  }, [handleDetectFace]);

  const handleStop = () => {
    setPause(true);
  };

  const handleGetFaces = async () => {
    const facesRects = drawingCoords.flatMap((face) =>
      face.reduce((acc, feature) => {
        if (feature[3] === TYPE_FACE) {
          acc.push(feature);
        }
        return acc;
      }, [])
    );

    const canvas = document.createElement("canvas");
    const context = canvas.getContext("2d");
    canvas.width = WIDTH;
    canvas.height = HIGHT;

    const drawImage = (src, context) =>
      new Promise((resolve) => {
        const image = new Image();
        image.src = src;
        image.onload = () => {
          context.drawImage(image, 0, 0);
          resolve();
        };
      });

    await drawImage(webcamRef.current.getScreenshot(), context);

    const imageFaces = facesRects.map((coords, idx) => {
      const [a, b, length] = coords;
      const x = b - length / 2;
      const y = a - length / 2;

      const imageData = context.getImageData(x, y, length, length);
      const imageCanvas = document.createElement("canvas");
      imageCanvas.width = length;
      imageCanvas.height = length;

      const imageContext = imageCanvas.getContext("2d");
      imageContext.putImageData(imageData, 0, 0);

      return { id: idx, image: imageCanvas.toDataURL("image/jpeg") };
    });

    onDetectedFaces(imageFaces);
  };

  useEffect(() => {
    if (!pause) {
      updatePhoto();
    }
  }, [updatePhoto, drawingCoords, pause]);

  return (
    <div className={s.webcamContainer}>
      <div className={s.webcamCanvas}>
        <Webcam
          screenshotFormat="image/jpeg"
          width={WIDTH}
          height={HIGHT}
          ref={webcamRef}
        />
        <PersonCanvas
          width={WIDTH}
          height={HIGHT}
          drawingCoords={drawingCoords}
          className={s.canvas}
        />
      </div>

      <button onClick={updatePhoto}>Start capturing</button>
      <button onClick={handleStop}>Stop capturing</button>

      <button onClick={handleGetFaces}>Get faces</button>
    </div>
  );
};
