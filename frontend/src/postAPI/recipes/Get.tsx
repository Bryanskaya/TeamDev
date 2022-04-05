import axios from "axios";
import { backUrl } from "../Common";


const GetRecipe = async function(id: number) {
    const response = await axios.get(backUrl + `/recipes/${id}`);
    return {
        status: response.status,
        content: response.data
    };
}

export default GetRecipe
