import React from "react";
import { Routes, Route } from "react-router";
import { BrowserRouter } from "react-router-dom";
import SearchHeader from "components/SearchHeader";
import Header from ".";


export const HeaderRouter: React.FC<{}> = () => {
    return <BrowserRouter>
        <Routes>
            <Route path="/" element={<SearchHeader title="Все рецепты"/>}/>
            <Route path="/me/likes" element={<Header subtitle="Понравилось" title="Мне" />}/>
            <Route path="/me/recipes" element={<Header subtitle="Автор" title="Я" />}/>

            <Route path="/auth/signin" element={<Header title="Вход" undertitle="Добро пожаловать. Снова." />}/>
            <Route path="/auth/signup" element={<Header title="Регистрация" undertitle="Чтобы получить доступ к тысячам новых возможностей в Вашем кулинарном самовыражении!" />}/>
            
            <Route path="/recipes/:id" element={<Header title=""/>}/>

            <Route path="*" element={<Header title="Страница не найдена"/>}/>
        </Routes>
    </BrowserRouter>
}
