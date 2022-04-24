import axios from "axios";
import { backUrl } from "..";

interface resp {
    status: number
    content: number
}

const GetLikes = async function(id: number): Promise<resp> {
    const response = await axios.get(backUrl + `/recipes/${id}/like/amount`);
    return {
        status: response.status,
        content: response.data as number
    };
}
export default GetLikes;
