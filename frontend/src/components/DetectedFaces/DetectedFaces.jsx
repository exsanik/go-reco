import { Face } from "src/components/Face";

import s from "./styles.module.css";

export const DetectedFaces = ({ faces, onNewUserAdded }) => {
  return (
    <div className={s.wrapper}>
      {faces.map(({ id, image }) => (
        <Face key={id} image={image} onNewUserAdded={onNewUserAdded} />
      ))}
    </div>
  );
};
