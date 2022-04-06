import axios from "axios";
import { backUrl } from "../Common";
import { Account } from "types/Account";

export const Create = async function(data: Account) {
    const response = await axios.post(backUrl + `/accounts`, data).catch((error) => {
        return {
            status: error.response?.status,
        };
    });

    return {
        status: response?.status,
    };
}
