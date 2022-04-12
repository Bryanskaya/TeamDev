import React from "react";
import { useNavigate } from "react-router-dom";
import { useCookies } from "react-cookie";

import SignUpPage from "./SignupPage";

const SignUp = () => {
    let navigate = useNavigate();
    let [cookie, setCookie] = useCookies(['token', 'role', 'login']);
    return ( 
        <SignUpPage navigate={navigate} cookie={cookie} setCookie={setCookie}/>
    )
}

export default SignUp;