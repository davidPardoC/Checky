import {
  Button,
  Flex,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Input,
} from "@chakra-ui/react";
import Link from "next/link";
import { useForm } from "react-hook-form";
import * as yup from "yup";
import { yupResolver } from "@hookform/resolvers/yup";
import AuthServices from "@/services/auth.services";
import { SignUpDto } from "@/types/auth";

const schema = yup
  .object({
    name: yup
      .string()
      .min(3, "Field is invalid")
      .required("Field is required."),
    lastName: yup
      .string()
      .min(3, "Field is invalid")
      .required("Field is required."),
    email: yup
      .string()
      .min(3, "Field is invalid")
      .email()
      .required("Field is required."),
    password: yup
      .string()
      .min(8, "Password must have at least 8 characters")
      .required("Field is required."),
    repeatPassword: yup
      .string()
      .min(8, "Password must have at least 8 characters")
      .required("Field is required."),
  })
  .required();

export const SignupForm = () => {
  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm({ resolver: yupResolver(schema) });

  const password = watch("password");
  const repeatPassword = watch("repeatPassword");

  const passwordMatch = password === repeatPassword;

  const onSubmit = (data: { [key: string]: any }) => {
    if (!passwordMatch) {
      return;
    }
    AuthServices.signup(data as SignUpDto)
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <FormControl marginTop={3} isInvalid={!!errors.name} isRequired>
        <FormLabel>Name</FormLabel>
        <Input type="text" placeholder="Name" {...register("name")} />
        <FormErrorMessage _firstLetter={{ textTransform: "capitalize" }}>
          {errors.name?.message as string}
        </FormErrorMessage>
      </FormControl>
      <FormControl marginTop={3} isInvalid={!!errors.lastName} isRequired>
        <FormLabel>Last Name</FormLabel>
        <Input type="text" placeholder="Last Name" {...register("lastName")} />
        <FormErrorMessage>
          {errors.lastName?.message as string}
        </FormErrorMessage>
      </FormControl>
      <FormControl marginTop={3} isInvalid={!!errors.email} isRequired>
        <FormLabel>Email</FormLabel>
        <Input type="email" placeholder="Email" {...register("email")} />
        <FormErrorMessage>{errors.email?.message as string}</FormErrorMessage>
      </FormControl>
      <FormControl marginTop={3} isInvalid={!!errors.password} isRequired>
        <FormLabel>Password</FormLabel>
        <Input
          type="Password"
          placeholder="Password"
          {...register("password")}
        />
        <FormErrorMessage>
          {errors.password?.message as string}
        </FormErrorMessage>
      </FormControl>
      <FormControl
        marginTop={3}
        marginBottom={5}
        isInvalid={!!errors.password || !passwordMatch}
        isRequired
      >
        <FormLabel>Repeat Password</FormLabel>
        <Input
          type="password"
          placeholder="Password"
          {...register("repeatPassword")}
        />
        <FormErrorMessage>
          {errors.repeatPassword?.message as string}
        </FormErrorMessage>
        {!passwordMatch && (
          <FormErrorMessage>The password do not match.</FormErrorMessage>
        )}
      </FormControl>
      <Flex justifyContent={"space-between"} alignItems={"center"}>
        <Link style={{ textDecoration: "underline" }} href={"/login"}>
          Already have an account?
        </Link>
        <Button type="submit" bg={"primary"} color={"white"}>
          Create Account
        </Button>
      </Flex>
    </form>
  );
};
