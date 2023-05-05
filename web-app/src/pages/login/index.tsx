import { LoginForm } from "@/components/LoginForm/LoginForm";
import { Container } from "@chakra-ui/react";
import Head from "next/head";

export default function LoginPage() {
  return (
    <>
      <Head>
        <title>Login</title>
      </Head>
      <Container mt={3}>
        <LoginForm />
      </Container>
    </>
  );
}
