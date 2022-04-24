import axios from "axios";
import { backUrl } from "..";
import { Account } from "types/Account";

interface resp {
    status: number
}

export const Create = async function(data: Account): Promise<resp> {
    const response = await axios.post(backUrl + `/accounts`, data).catch((error) => {
        return {
            status: error.response?.status,
        };
    });

    return {
        status: response?.status,
    };
}
