import {
  Box,
  Button,
  Flex,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Heading,
  Input,
  Modal,
  ModalBody,
  ModalContent,
  ModalHeader,
  ModalOverlay,
  Select,
} from "@chakra-ui/react";
import Link from "next/link";
import { useForm } from "react-hook-form";
import * as yup from "yup";
import { yupResolver } from "@hookform/resolvers/yup";
import AuthServices from "@/services/auth.services";
import { SignUpDto } from "@/types/auth";
import useUsersStore from "@/store/users.store";

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

type CreateUserModalProps = {
  onClose: () => void;
};

export const CreateUserForm = ({ onClose }: CreateUserModalProps) => {
  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm({ resolver: yupResolver(schema) });

  const roles = useUsersStore((state) => state.roles);

  const password = watch("password");
  const repeatPassword = watch("repeatPassword");

  const passwordMatch = password === repeatPassword;

  const onSubmit = (data: { [key: string]: any }) => {
    if (!passwordMatch) {
      return;
    }
    AuthServices.signup(data as SignUpDto);
  };

  return (
    <Modal isOpen onClose={onClose}>
      <ModalOverlay />

      <ModalContent>
        <ModalHeader>Create User</ModalHeader>
        <ModalBody>
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
              <Input
                type="text"
                placeholder="Last Name"
                {...register("lastName")}
              />
              <FormErrorMessage>
                {errors.lastName?.message as string}
              </FormErrorMessage>
            </FormControl>
            <FormControl marginTop={3} isInvalid={!!errors.email} isRequired>
              <FormLabel>Email</FormLabel>
              <Input type="email" placeholder="Email" {...register("email")} />
              <FormErrorMessage>
                {errors.email?.message as string}
              </FormErrorMessage>
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
            <FormControl marginTop={3}>
              <FormLabel>Role</FormLabel>
              <Select textTransform={"capitalize"} placeholder="Select option">
                {roles.map((role) => (
                  <option className="capitalize" key={role.name} value={role.name}>
                    {role.name}
                  </option>
                ))}
              </Select>
            </FormControl>

            <Flex
              justifyContent={"flex-end"}
              gap={2}
              alignItems={"center"}
              marginTop={3}
            >
              <Button bg={"secondary"} color={"white"} onClick={onClose}>
                Cancel
              </Button>
              <Button type="submit" bg={"primary"} color={"white"}>
                Create Account
              </Button>
            </Flex>
          </form>
        </ModalBody>
      </ModalContent>
    </Modal>
  );
};
