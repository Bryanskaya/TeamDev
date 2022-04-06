import React from "react";
import { useNavigate } from "react-router-dom";
import { useCookies } from "react-cookie";

import LoginPage from "./LoginPage";

const Login = () => {
    let navigate = useNavigate();
    let [cookie, setCookie] = useCookies(['token', 'role', 'login']);
    return (
        <LoginPage navigate={navigate} cookie={cookie} setCookie={setCookie}/>
    )
}

export default Login;
