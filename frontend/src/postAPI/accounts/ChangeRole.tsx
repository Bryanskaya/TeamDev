import axios from "axios";
import { backUrl } from "..";

interface resp {
    status: number
}

export const ChangeRole = async function(login: string, role: string): Promise<resp> {
    var body = {'role':role};
    const response = await axios.patch(backUrl + `/accounts/${login}/role`, body).catch((error) => {
        return {
            status: error.response?.status,
        };
    });

    return {
        status: response?.status,
    };
}
