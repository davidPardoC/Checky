import {
  Alert,
  AlertDescription,
  AlertIcon,
  AlertTitle,
} from "@chakra-ui/react";
import React from "react";

type AlertProps = {
  title?: string;
  message?: string;
};

const AlertComponent = ({ title = "", message = "" }: AlertProps) => {
  return (
    <Alert status="error" position={"absolute"} zIndex={2}>
      <AlertIcon />
      <AlertTitle>{title}</AlertTitle>
      <AlertDescription>{message}</AlertDescription>
    </Alert>
  );
};

export default AlertComponent;
