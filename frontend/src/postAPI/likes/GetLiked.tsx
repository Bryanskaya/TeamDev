import axios from "axios";
import { backUrl } from "..";

interface resp {
    status: number
    content: boolean
}

const GetIsLiked = async function(id: number): Promise<resp> {
    const response = await axios.get(backUrl + `/recipes/${id}/isliked`);
    return {
        status: response.status,
        content: response.data as boolean
    };
}
export default GetIsLiked;
