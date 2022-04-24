import { apiInstance } from "./baseAxios";

export const updateUser = async (id, body) => {
  const { data } = await apiInstance.put(`/api/users/${id}`, body);

  return data;
};
