import axios from "axios";
import { Step } from "types/Step";
import { backUrl } from "..";


const PatchStep = async function(data: Step) {
    const response = await axios.patch(backUrl + `/recipes/${data.recipe}/steps/${data.num}`, data);
    return {
        status: response.status
    };
}
export default PatchStep;