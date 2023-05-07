import { Role } from "@/types/roles";
import { AxiosInstance } from "axios";

const getRoles = async (axiosInstance: AxiosInstance):Promise<Role[] | undefined > => {
  try {
    const { data: roles } = await axiosInstance.get<Role[]>("/api/v1/roles");
    return roles;
  } catch (error) {
    console.log(error);
  }
};

const RolesService = { getRoles };

export default RolesService;
