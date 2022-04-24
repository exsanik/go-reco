import { useState } from "react";

import { recognizeFace } from "src/api/recognizeFace";
import { createUser } from "src/api/createUser";
import { updateUser } from "src/api/updateUser";

import s from "./styles.module.css";

export const Face = ({ image, onNewUserAdded }) => {
  const [face, setFace] = useState();
  const [saved, setSaved] = useState(false);
  const [userName, setUserName] = useState("");

  if (!image) return null;

  const base64Image = image.split(",", 2)[1];

  const handleRecognize = async () => {
    const { data } = await recognizeFace({ photo: base64Image });
    setFace(data);
  };

  const handleInputChange = (event) => {
    setUserName(event.target.value);
  };

  const handleSaveToDatabase = async () => {
    const isNewUser = face === null;

    if (isNewUser) {
      const user = await createUser({ name: userName, photo: base64Image });
      onNewUserAdded(user);
    } else {
      await updateUser(face?.user.ID, { photo: base64Image });
    }
    setSaved(true);
  };

  return (
    <div className={s.item}>
      <img className={s.face} src={image} alt="face" />
      <button onClick={handleRecognize}>Recognize</button>

      {face === null ? (
        <>
          <span>User was not recognized</span>
          <input
            type="text"
            name="name"
            value={userName}
            onChange={handleInputChange}
          />

          {!saved ? (
            <button onClick={handleSaveToDatabase} disabled={!userName}>
              Save to database
            </button>
          ) : (
            <span>saved</span>
          )}
        </>
      ) : (
        face && (
          <>
            {face?.user && face?.user.Name}
            {face?.descriptor && (
              <img
                className={s.face}
                src={`data:image/jpeg;base64,${face?.descriptor.Photo}`}
                alt="face"
              />
            )}
            {!saved ? (
              <button onClick={handleSaveToDatabase}>Save to database</button>
            ) : (
              <span>saved</span>
            )}
          </>
        )
      )}
    </div>
  );
};
