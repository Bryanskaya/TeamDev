import axios from "axios";
import { backUrl } from "..";


const GetIngredient = async function(id: number) {
    const response = await axios.get(backUrl + `/recipes/${id}/ingredients`);
    return {
        status: response.status,
        content: response.data
    };
}
export default GetIngredient;