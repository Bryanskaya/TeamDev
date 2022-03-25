import * as React from "react";
import Fonts from "./components/Fonts";
import Home from "./components/layout/Home";
import theme from "./extendTheme";
import { ChakraProvider, Box, Text, Menu, Container } from "@chakra-ui/react";

export const App = () => (
  <ChakraProvider theme={theme}>
    <Fonts />
    Hello worl
    <Home />
  </ChakraProvider>
);
