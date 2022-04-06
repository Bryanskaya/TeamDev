import axios from "axios";
import { backUrl } from "../Common";


const GetLikes = async function(id: number) {
    const response = await axios.get(backUrl + `/recipes/${id}/like/amount`);
    return {
        status: response.status,
        content: response.data
    };
}
export default GetLikes;