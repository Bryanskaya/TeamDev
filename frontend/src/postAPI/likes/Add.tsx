import axios from "axios";
import { backUrl } from "..";

interface resp {
    status: number
}

const AddLike = async function(id: number): Promise<resp> {
    const response = await axios.put(backUrl + `/recipes/${id}/like`, {withCredentials: true});
    return {
        status: response.status,
    };
}
export default AddLike;