import { apiInstance } from "./baseAxios";

export const recognizeFace = async (body) => {
  const { data } = await apiInstance.post("/api/recognize", body);

  return data;
};
