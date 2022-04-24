import { useEffect, useState } from "react";
import { deleteUser } from "src/api/deleteUser";
import { getAllUsers } from "src/api/getAllUsers";

import s from "./styles.module.css";

export const UsersControl = ({ newUser }) => {
  const [users, setUsers] = useState([]);

  useEffect(() => {
    getAllUsers().then((resp) => setUsers(resp?.data || []));
  }, []);

  useEffect(() => {
    if (newUser?.data?.ID) {
      setUsers((prevUsers) => [...prevUsers, newUser?.data]);
    }
  }, [newUser]);

  const handleDelete = (id) => async () => {
    await deleteUser(id);

    setUsers((prevUsers) => prevUsers.filter(({ ID }) => ID !== id));
  };

  return (
    <div className={s.userWrapper}>
      {users.map(({ ID, Name, UserDescriptors }) => (
        <div key={ID} className={s.user}>
          <p>{Name}</p>
          <img
            src={`data:image/jpeg;base64,${UserDescriptors[0].Photo}`}
            alt="user"
          />
          <button onClick={handleDelete(ID)}>Delete</button>
        </div>
      ))}
    </div>
  );
};
