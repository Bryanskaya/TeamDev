import React from "react";
import { Routes, Route, useParams } from "react-router";
import { BrowserRouter } from "react-router-dom";
import SearchHeader from "components/Header/Search";
import Header from ".";
import CategoryHeader from "./Category/CategoryHeader";
import RecipeHeader from "./Recipe/RecipeHeader";


export const HeaderRouter: React.FC<{}> = () => {
    const params = useParams()
    return <BrowserRouter>
        <Routes>
            <Route path="/" element={<SearchHeader title="Все рецепты"/>}/>
            <Route path="/users" element={<SearchHeader title="Все пользователи"/>}/>
            <Route path="/me/likes" element={<Header subtitle="Понравилось" title="Мне" />}/>
            <Route path="/me/recipes" element={<Header subtitle="Автор" title="Я" />}/>

            <Route path="/accounts/:login/recipes" element={<RecipeHeader subtitle="Автор" title="" />}/>
            <Route path="/accounts/:login/likes" element={<RecipeHeader subtitle="Понравилось" title="" />}/>

            <Route path="/auth/signin" element={<Header title="Вход" undertitle="Добро пожаловать. Снова." />}/>
            <Route path="/auth/signup" element={<Header title="Регистрация" undertitle="Чтобы получить доступ к тысячам новых возможностей в Вашем кулинарном самовыражении!" />}/>
            
            <Route path="/recipes/:id" element={<Header title=""/>}/>
            <Route path="/categories/:title" element={<CategoryHeader subtitle="Категория" title=""/>}/>

            <Route path="*" element={<Header title="Страница не найдена"/>}/>
        </Routes>
    </BrowserRouter>
}
