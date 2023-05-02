import { ChakraProvider, extendTheme } from "@chakra-ui/react";
import type { AppProps } from "next/app";
import { colors } from "../theme";

export default function App({ Component, pageProps }: AppProps) {
  const theme = extendTheme({ colors });

  return (
    <ChakraProvider theme={theme}>
      <Component {...pageProps} />
    </ChakraProvider>
  );
}
