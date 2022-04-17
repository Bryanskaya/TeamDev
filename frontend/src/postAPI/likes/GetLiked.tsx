import axios from "axios";
import { backUrl } from "..";


const GetIsLiked = async function(id: number) {
    const response = await axios.get(backUrl + `/recipes/${id}/isliked`);
    return {
        status: response.status,
        content: response.data
    };
}
export default GetIsLiked;