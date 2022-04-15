import React from "react";
import {
    Box,
    Link,
    Text
} from "@chakra-ui/react";

import styles from "./Category.module.scss";

interface CategoryProps {
    name: string
}

const Category: React.FC<CategoryProps> = (props) => {
    const name = props.name[0].toUpperCase() + props.name.slice(1)
    return (
        <Link className={styles.category}> 
            {name}
        </Link>
    )
}

export default React.memo(Category);
