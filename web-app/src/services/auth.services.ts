import {
  LoginDto,
  LoginResponseDto,
  SignUpDto,
  SingupResponse,
} from "@/types/auth";
import axios from "axios";
import { setCookie } from "cookies-next";

const basePath = "/api/v1/auth";

const signup = async (newUser: SignUpDto) => {
  const { data } = await axios.post<SingupResponse>(
    `${basePath}`,
    newUser
  );
  return data;
};

const login = async (loginDto: LoginDto) => {
  try {
    const { data } = await axios.post<LoginResponseDto>(
      `${basePath}/login`,
      loginDto
    );
    setCookie("token", data.token);
    setCookie("refreshToken", data.refreshToken);
  } catch (error) {}
};

const AuthServices = { signup, login };

export default AuthServices;
