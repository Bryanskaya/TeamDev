import axios from "axios";
import { backUrl } from "..";

interface resp {
    status: number
}

export const Logout = async function(removeCookie): Promise<resp> {
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
