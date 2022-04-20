import { useState } from "react";
import Webcam from "react-webcam";

import { CapturedPhoto } from "../CapturedPhoto/CapturedPhoto";

import s from "./styles.module.css";

const HIGHT = 384;
const WIDTH = 512;

const App = () => {
  const [photo, setPhoto] = useState("");

  const handleCapturePhoto = (getScreenshot) => () => {
    const imageSrc = getScreenshot();
    setPhoto(imageSrc);
  };

  return (
    <div className={s.webcamContainer}>
      <Webcam height={HIGHT} screenshotFormat="image/jpeg" width={WIDTH}>
        {({ getScreenshot }) => (
          <button onClick={handleCapturePhoto(getScreenshot)}>
            Capture photo
          </button>
        )}
      </Webcam>

      <CapturedPhoto photo={photo} height={HIGHT} width={WIDTH} />
    </div>
  );
};

export default App;
