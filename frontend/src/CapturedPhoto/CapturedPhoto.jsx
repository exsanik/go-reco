import { useState } from "react";
import { Canvas } from "../Canvas";

import s from "./styles.module.css";

export const CapturedPhoto = ({ photo, height, width }) => {
  const [drawingCoords, setDrawingCoords] = useState([]);
  if (!photo) return null;

  const handleDetectFace = async () => {
    const base64 = photo.split(",", 2)[1];

    const resp = await fetch("http://localhost:8080/api/detect_face", {
      method: "POST",
      body: JSON.stringify({ photo: base64 }),
    });
    const data = await resp.json();
    setDrawingCoords(data?.data || []);

    if (data?.data?.length === 0) {
      alert("No faces detected");
    }
  };

  return (
    <div className={s.wrapper}>
      <Canvas
        base64Image={photo}
        drawingCoords={drawingCoords}
        height={height}
        width={width}
      />
      <button onClick={handleDetectFace}>Detect face</button>
    </div>
  );
};
