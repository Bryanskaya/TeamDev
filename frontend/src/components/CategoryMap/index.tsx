
import React from "react";

import {default as Elem } from "./CategoryMap";
import { AllCategoriesResp } from "postAPI";

export interface CategoryMapProps {
    getCall: () => Promise<AllCategoriesResp>
}

const CategoryMap = (props: CategoryMapProps) => {
    return (
        <Elem {...props}/>
    )
}

export default CategoryMap;
