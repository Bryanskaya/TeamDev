import axios from "axios";
import { backUrl } from "..";


const DeleteLike = async function(id: number) {
    const response = await axios.delete(backUrl + `/recipes/${id}/like`, {withCredentials: true});
    return {
        status: response.status,
    };
}
export default DeleteLike;