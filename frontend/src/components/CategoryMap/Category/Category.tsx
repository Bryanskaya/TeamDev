import React from "react";
import {
    Link,
    LinkProps
} from "@chakra-ui/react";

import styles from "./Category.module.scss";
import { capitalize } from "functions";

interface CategoryProps extends LinkProps {
    name: string
}

const Category: React.FC<CategoryProps> = (props) => {
    let { name, ...rest } = props

    if (!rest.onClick)
        rest.href = 'categories/' + props.name 

    return (
        <Link className={styles.category} {...rest}> 
            {capitalize(props.name)}
        </Link>
    )
}

export default React.memo(Category);
