import axios from "axios";
import { backUrl } from "../Common";

export const Logout = async function(removeCookie) {
    const response = await axios.post(backUrl + `/accounts/logout`).catch((error) => {
        return {
            status: error.response?.status,
        };
    });

    removeCookie('token', {path: '/'})
    removeCookie('login', {path: '/'})
    removeCookie('role', {path: '/'})

    return {
        status: response?.status,
    };
}
