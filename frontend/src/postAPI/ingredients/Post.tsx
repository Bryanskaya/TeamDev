import axios from "axios";
import { backUrl } from "../Common";
import {Ingredient as IngredientT} from "types/Ingredient"


const PutIngredient = async function(id: number, data: IngredientT) {
    const response = await axios.put(backUrl + `/recipes/${id}/ingredients`, data);
    return {
        status: response.status
    };
}
export default PutIngredient;