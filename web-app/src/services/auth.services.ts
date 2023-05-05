import { LoginDto, SignUpDto, SingupResponse } from "@/types/auth";
import axios from "axios";

const basePath = "/api/v1/auth";

const signup = async (newUser: SignUpDto) => {
  const { data } = await axios.post<SingupResponse>(`${basePath}/signup`, newUser);
  return data
};

const login = async (loginDto: LoginDto) => {
  const { data } = await axios.post<SingupResponse>(`${basePath}/login`, loginDto);
  return data
}

const AuthServices = { signup, login };

export default AuthServices;
