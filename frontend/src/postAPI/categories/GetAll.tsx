import axios from "axios";
import { AllCategoriesResp, backUrl } from "..";

const GetCategories = async function(title?: string): Promise<AllCategoriesResp> {
    const response = await axios.get(backUrl + "/categories");
    return {
        status: response.status,
        content: response.data
    };
}

export default GetCategories
