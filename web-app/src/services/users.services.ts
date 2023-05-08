import { AxiosInstance } from "axios";

const getAllUsers = async (axiosInstance: AxiosInstance) => {
  const { data } = await axiosInstance.get("/api/v1/users");
  return data;
};

const UserServices = { getAllUsers };

export default UserServices;
