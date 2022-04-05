import axios from "axios";
import { backUrl } from "../Common";


const DeleteRecipe = async function(id: number) {
    const response = await axios.delete(backUrl + `/recipes/${id}`);
    return {
        status: response.status,
    };
}

export default DeleteRecipe
