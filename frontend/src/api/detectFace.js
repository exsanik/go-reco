import { apiInstance } from "./baseAxios";

export const detectFace = async (body) => {
  const { data } = await apiInstance.post("/api/detect", body);

  return data;
};
