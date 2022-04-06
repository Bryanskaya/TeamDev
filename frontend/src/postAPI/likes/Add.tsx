import axios from "axios";
import { backUrl } from "../Common";


const AddLike = async function(id: number) {
    const response = await axios.put(backUrl + `/recipes/${id}/like`, {withCredentials: true});
    return {
        status: response.status,
    };
}
export default AddLike;