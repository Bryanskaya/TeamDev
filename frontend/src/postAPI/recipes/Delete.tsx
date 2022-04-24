import axios from "axios";
import { backUrl } from "..";

interface resp {
    status: number
}

const DeleteRecipe = async function(id: number): Promise<resp> {
    const response = await axios.delete(backUrl + `/recipes/${id}`);
    return {
        status: response.status,
    };
}

export default DeleteRecipe
