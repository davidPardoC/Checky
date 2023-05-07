import { ChakraProvider, extendTheme } from "@chakra-ui/react";
import type { AppProps } from "next/app";
import { colors } from "../theme";
import clientUtils from "@/utils/client/client";
import AxiosUtils from "@/utils/client/axios";
import "../styles.css"

if(clientUtils.isClient()){
  AxiosUtils.setupClientSideAxiosInstance()
}

export default function App({ Component, pageProps }: AppProps) {
  const theme = extendTheme({ colors });

  return (
    <ChakraProvider theme={theme}>
      <Component {...pageProps} />
    </ChakraProvider>
  );
}
