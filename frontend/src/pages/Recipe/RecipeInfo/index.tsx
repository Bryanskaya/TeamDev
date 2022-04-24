import React from "react";
import { useCookies } from "react-cookie";
import { useNavigate, useParams } from "react-router-dom";
import RecipeInfoPage from "./RecipeInfoPage";

const RecipeInfo = () => {
    let [cookie] = useCookies(['role', 'login']);
    let navigate = useNavigate();
    return (
        <RecipeInfoPage match={useParams()} navigate={navigate} cookie={cookie}/>
    )
}

export default RecipeInfo;