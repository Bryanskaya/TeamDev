import axios from "axios";
import { backUrl } from "..";

export const ChangeRole = async function(login: string, role: string) {
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
