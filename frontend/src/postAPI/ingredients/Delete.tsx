import axios from "axios";
import { backUrl } from "..";

interface resp {
    status: number
}

const DeleteIngredient = async function(id: number, title: string): Promise<resp> {
    const response = await axios.delete(backUrl + `/recipes/${id}/ingredients/${title}`, );
    return {
        status: response.status
    };
}
export default DeleteIngredient;