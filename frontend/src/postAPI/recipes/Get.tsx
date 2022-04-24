import axios from "axios";
import { Recipe } from "types/Recipe";
import { backUrl } from "..";

interface resp {
    status: number
    content: Recipe
}

const GetRecipe = async function(id: number): Promise<resp> {
    const response = await axios.get(backUrl + `/recipes/${id}`);
    return {
        status: response.status,
        content: response.data as Recipe
    };
}

export default GetRecipe
