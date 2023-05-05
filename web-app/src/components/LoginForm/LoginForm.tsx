import AuthServices from "@/services/auth.services";
import { LoginDto } from "@/types/auth";
import {
  Box,
  Button,
  Flex,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Heading,
  Input,
} from "@chakra-ui/react";
import { yupResolver } from "@hookform/resolvers/yup";
import Link from "next/link";
import { useForm } from "react-hook-form";
import * as yup from "yup";

const schema = yup.object({
  email: yup.string().email().required(),
  password: yup.string().required(),
});

export const LoginForm = () => {
  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm({ resolver: yupResolver(schema) });

  const onSubmit = async (data: { [key: string]: any }) => {
    const responseData = await AuthServices.login(data as LoginDto)
    console.log(responseData)
  }

  return (
    <Box boxShadow={"1px 0px 3px"} padding={3} borderRadius={5}>
      <form onSubmit={handleSubmit(onSubmit)}>
        <Heading color={"primary"} textAlign={"center"}>
          Checky
        </Heading>
        <FormControl isRequired isInvalid={!!errors.email}>
          <FormLabel>Email address</FormLabel>
          <Input type="email" {...register("email")} />
          <FormErrorMessage>
            {errors.email?.message as string}
          </FormErrorMessage>
        </FormControl>
        <FormControl isRequired isInvalid={!!errors.password}>
          <FormLabel>Password</FormLabel>
          <Input type="password" {...register("password")} />
          <FormErrorMessage>
            {errors.password?.message as string}
          </FormErrorMessage>
        </FormControl>
        <Flex
          justifyContent={"space-between"}
          alignItems={"center"}
          marginTop={4}
        >
          <Link style={{ textDecoration: "underline" }} href={"/signup"}>
            Do not have an account?
          </Link>
          <Button bg={"primary"} color={"white"} type="submit">
            Iniciar Sesión
          </Button>
        </Flex>
      </form>
    </Box>
  );
};
