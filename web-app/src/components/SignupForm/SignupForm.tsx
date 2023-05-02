import {  Button, FormControl, FormLabel, Input } from "@chakra-ui/react";
import { Link } from '@chakra-ui/next-js'

export const SignupForm = () => {
  return (
    <form>
      <FormControl marginTop={3}>
        <FormLabel>Name</FormLabel>
        <Input type="text" placeholder="Name" />
      </FormControl>
      <FormControl marginTop={3}>
        <FormLabel>Last Name</FormLabel>
        <Input type="text" placeholder="Last Name" />
      </FormControl>
      <FormControl marginTop={3}>
        <FormLabel>Email</FormLabel>
        <Input type="email" placeholder="Email" />
      </FormControl>
      <FormControl marginTop={3}>
        <FormLabel>Password</FormLabel>
        <Input type="Password" placeholder="Password" />
      </FormControl>
      <Link href={"/login"}>
        Already have an account?
      </Link>
      <Button>Create Account</Button>
    </form>
  );
};
