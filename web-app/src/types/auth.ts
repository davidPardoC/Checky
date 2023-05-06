export type SignUpDto = {
  name: string;
  lastName: string;
  email: string;
  password: string;
};

export type SingupResponse = {
  message: string;
};

export type LoginDto = {
  email: string;
  password: string;
};

export type LoginResponseDto = {
  token: string;
  refreshToken: string;
};
