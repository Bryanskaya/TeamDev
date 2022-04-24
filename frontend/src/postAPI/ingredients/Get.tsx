import axios from "axios";
import { Ingredient } from "types/Ingredient";
import { backUrl } from "..";

interface resp {
    status: number
    content: Ingredient[]
}

const GetIngredient = async function(id: number): Promise<resp> {
    const response = await axios.get(backUrl + `/recipes/${id}/ingredients`);
    return {
        status: response.status,
        content: response.data as Ingredient[]
    };
}
export default GetIngredient;