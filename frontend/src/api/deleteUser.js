import { apiInstance } from "./baseAxios";

export const deleteUser = async (id) => {
  const { data } = await apiInstance.delete(`/api/users/${id}`);

  return data;
};
