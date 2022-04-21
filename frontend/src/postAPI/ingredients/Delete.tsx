import axios from "axios";
import { backUrl } from "..";


const DeleteIngredient = async function(id: number, title: string) {
    const response = await axios.delete(backUrl + `/recipes/${id}/ingredients/${title}`, );
    return {
        status: response.status,
        content: response.data
    };
}
export default DeleteIngredient;