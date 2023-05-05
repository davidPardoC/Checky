import { SignupForm } from "@/components/SignupForm/SignupForm";
import { Container, Heading } from "@chakra-ui/react";
import Head from "next/head";

export default function SignupScreen() {
  return (
    <>
      <Head>
        <title>SignUp</title>
      </Head>
      <Container mt={3}>
        <SignupForm />
      </Container>
    </>
  );
}
