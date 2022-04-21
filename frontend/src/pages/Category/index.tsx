import React from "react";
import { useCookies } from "react-cookie";
import { useNavigate, useParams } from "react-router-dom";
import { default as Page} from "./CategoryPage";

const CategoryPage = () => {
    let [cookie] = useCookies(['role', 'login']);
    let navigate = useNavigate();
    return (
        <Page match={useParams()} navigate={navigate} cookie={cookie}/>
    )
}

export default CategoryPage;
