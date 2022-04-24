import { useState } from "react";
import { DetectedFaces } from "src/components/DetectedFaces";
import { UsersControl } from "src/components/UsersControl";
import { FaceDetector } from "src/components/FaceDetector";

import s from "./styles.module.css";

const App = () => {
  const [detectedFaces, setDetectedFaces] = useState([]);
  const [newUser, setNewUser] = useState({});

  return (
    <div className={s.webcamContainer}>
      <FaceDetector onDetectedFaces={setDetectedFaces} />

      <DetectedFaces faces={detectedFaces} onNewUserAdded={setNewUser} />

      <UsersControl newUser={newUser} />
    </div>
  );
};

export default App;
