import axios from "axios";
import { backUrl } from "..";
import {Ingredient as IngredientT} from "types/Ingredient"

interface resp {
    status: number
}

const PutIngredient = async function(id: number, data: IngredientT): Promise<resp> {
    const response = await axios.put(backUrl + `/recipes/${id}/ingredients`, data);
    return {
        status: response.status
    };
}
export default PutIngredient;