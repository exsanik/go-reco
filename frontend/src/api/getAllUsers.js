import { apiInstance } from "./baseAxios";

export const getAllUsers = async () => {
  const { data } = await apiInstance.get("/api/users");

  return data;
};
