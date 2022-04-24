import { apiInstance } from "./baseAxios";

export const createUser = async (body) => {
  const { data } = await apiInstance.post("/api/users", body);

  return data;
};
