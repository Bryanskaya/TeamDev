import axios from "axios";
import { backUrl } from "..";

interface resp {
    status: number
}

const DeleteLike = async function(id: number): Promise<resp> {
    const response = await axios.delete(backUrl + `/recipes/${id}/like`, {withCredentials: true});
    return {
        status: response.status,
    };
}
export default DeleteLike;