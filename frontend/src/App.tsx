import * as React from "react";
import Fonts from "./components/Fonts";
// import Home from "./components/layout/Home";
import theme from "./extendTheme";
import { ChakraProvider, Box, Text, Menu, Container } from "@chakra-ui/react";
import { Routes, Route } from "react-router";
import { BrowserRouter } from "react-router-dom";

interface HomeProps {}
const Home: React.FC<HomeProps> = ({}) => {
  return (
    <Box backgroundColor="bg" minH="100vh" h="auto">
      <Container maxW="1000px" minH="95%"
        display="flex" 
        paddingX="0px" paddingY="30px"  
        alignSelf="center" justifyContent="center"
        textStyle="body"
      >
        <Routing />
      </Container>
    </Box>
  );
};

function Routing() {
  return <BrowserRouter>
    <Routes>
      <Route path="/" element={'All Recipes'}/>
      <Route path="/me/likes" element={'I liked these recipes'}/>
      <Route path="/me/recipes" element={'My recipes'}/>

      <Route path="/auth/signin" element={'Login'}/>
      <Route path="/auth/signup" element={'SignUp'}/>

      <Route path="/recipes/:id" element={'Recipe'}/>

      <Route path="*" element={<NotFound />}/>
    </Routes>
  </BrowserRouter>
}

function NotFound () {
  return <h1>Page not Found</h1>
}

export const App = () => {
  return (
    <ChakraProvider theme={theme}>
      <Fonts />
      <Home />
    </ChakraProvider>
  )
};


