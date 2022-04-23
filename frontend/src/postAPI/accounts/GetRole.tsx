import axios from "axios";
import { backUrl } from "..";


const GetRole = async function(login: string) {
    const response = await axios.get(backUrl + `/accounts/${login}/role`);
    return {
        status: response.status,
        content: response.data
    };
}
export default GetRole;