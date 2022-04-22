import axios from "axios";
import { AllUsersResp, backUrl } from "..";

const GetUsers = async function(login?: string): Promise<AllUsersResp> {
    const params = { login: login ? login : '' }
    const response = await axios.get(backUrl + "/accounts", {params:params});
    return {
        status: response.status,
        content: response.data
    };
}

export default GetUsers
