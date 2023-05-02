import {  Button, FormControl, FormErrorMessage, FormLabel, Input } from "@chakra-ui/react";
import Link from "next/link";

export const SignupForm = () => {
  return (
    <form>
      <FormControl marginTop={3} isRequired isInvalid>
        <FormLabel>Name</FormLabel>
        <Input type="text" placeholder="Name" />
        <FormErrorMessage>Name is required</FormErrorMessage>
      </FormControl>
      <FormControl marginTop={3} isRequired>
        <FormLabel>Last Name</FormLabel>
        <Input type="text" placeholder="Last Name" />
        <FormErrorMessage>Name is required</FormErrorMessage>
      </FormControl>
      <FormControl marginTop={3} isRequired>
        <FormLabel>Email</FormLabel>
        <Input type="email" placeholder="Email" />
        <FormErrorMessage>Name is required</FormErrorMessage>
      </FormControl>
      <FormControl marginTop={3} isRequired>
        <FormLabel>Password</FormLabel>
        <Input type="Password" placeholder="Password" />
        <FormErrorMessage>Name is required</FormErrorMessage>
      </FormControl>
      <Link href={"/login"}>
        Already have an account?
      </Link>
      <Button type="submit" bg={"primary"} color={"white"}>Create Account</Button>
    </form>
  );
};
