import axios from "axios";
import { AllRecipeResp, backUrl } from "..";

const GetRecipes = async function(title?: string): Promise<AllRecipeResp> {
    const params = { title: title ? title : '' }
    const response = await axios.get(backUrl + `/categories/${title}/recipes`, {params:params});
    return {
        status: response.status,
        content: response.data
    };
}

export default GetRecipes
