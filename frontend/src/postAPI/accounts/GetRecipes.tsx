import axios from "axios";
import { AllRecipeResp, backUrl } from "..";

const GetRecipes = async function(login: string): Promise<AllRecipeResp> {
    const response = await axios.get(backUrl + `/accounts/${login}/recipes`);
    return {
        status: response.status,
        content: response.data
    };
}

export default GetRecipes