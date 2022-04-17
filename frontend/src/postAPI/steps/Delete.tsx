import axios from "axios";
import { backUrl } from "..";


const DeleteStep = async function(id_rcp: number, num: number) {
    const response = await axios.delete(backUrl + `/recipes/${id_rcp}/steps/${num}`, );
    return {
        status: response.status
    };
}
export default DeleteStep;