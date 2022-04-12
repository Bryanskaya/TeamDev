import React from "react";
import {
  Box,
  Container,
} from "@chakra-ui/react";

import styles from "./Header.module.scss";

import Navbar from "./Navbar"
import Titles, { TitlesProps } from "components/Header/Titles/Titles";
import { Routes, Route } from "react-router";
import { BrowserRouter } from "react-router-dom";

interface HeaderProps extends TitlesProps {
    role?: string
    addField?: React.FC
}

const Header: React.FC<HeaderProps> = (props) => {
    return (
    <Box className={styles.header}>
        <Container>
            <Navbar />
            <Titles {...props}/>
            {props.addField && <props.addField />}
        </Container>
    </Box>
    );
}


export const HeaderRouter: React.FC<{}> = ({}) => {
    return <BrowserRouter>
        <Routes>
            <Route path="/" element={<Header title="Все рецепты"/>}/>
            <Route path="/me/likes" element={<Header subtitle="Понравилось" title="Мне" />}/>
            <Route path="/me/recipes" element={<Header subtitle="Автор" title="Я" />}/>

            <Route path="/auth/signin" element={<Header title="Вход" undertitle="Добро пожаловать. Снова." />}/>
            <Route path="/auth/signup" element={<Header title="Регистрация" undertitle="Чтобы получить доступ к тысячам новых возможностей в Вашем кулинарном самовыражении!" />}/>
            
            <Route path="/recipes/:id" element={<Header title=""/>}/>

            <Route path="*" element={<Header title="Страница не найдена"/>}/>
        </Routes>
    </BrowserRouter>
}

export default React.memo(Header);
