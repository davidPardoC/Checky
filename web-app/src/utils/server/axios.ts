import axios from "axios";

export const getServerSideAxiosInstance = () => {
  const axiosInstance = axios.create({ baseURL: "http://localhost:8080" });
  return axiosInstance;
};
