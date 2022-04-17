import axios from "axios";
import { Step } from "types/Step";
import { backUrl } from "..";


const PushStep = async function(id: number, data: Step) {
    const response = await axios.post(backUrl + `/recipes/${id}/steps`, data);
    return {
        status: response.status,
        content: response.data
    };
}
export default PushStep;