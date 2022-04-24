import axios from "axios";
import { Account } from "types/Account";
import { backUrl } from "..";

interface resp {
    status: number
    content: Account
}

const GetRole = async function(login: string): Promise<resp> {
    const response = await axios.get(backUrl + `/accounts/${login}/role`);
    return {
        status: response.status,
        content: response.data as Account
    };
}
export default GetRole;