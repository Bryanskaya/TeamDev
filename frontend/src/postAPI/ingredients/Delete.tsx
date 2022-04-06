import axios from "axios";
import { backUrl } from "../Common";
import {Ingredient as IngredientT} from "types/Ingredient"


const DeleteIngredient = async function(id: number, title: string) {
    const response = await axios.delete(backUrl + `/recipes/${id}/ingredients/${title}`, );
    return {
        status: response.status,
        content: response.data
    };
}
export default DeleteIngredient;