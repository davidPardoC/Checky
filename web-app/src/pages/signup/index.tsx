import { SignupForm } from "@/components/SignupForm/SignupForm";
import { Container, Heading } from "@chakra-ui/react";

export default function SignupScreen() {
  return (
    <Container mt={3}>
      <Heading textAlign={"center"} color={"primary"}>
        Checky
      </Heading>
      <SignupForm />
    </Container>
  );
}
