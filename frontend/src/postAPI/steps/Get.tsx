import axios from "axios";
import { backUrl } from "..";


const GetSteps = async function(id: number) {
    const response = await axios.get(backUrl + `/recipes/${id}/steps`);
    return {
        status: response.status,
        content: response.data
    };
}
export default GetSteps;