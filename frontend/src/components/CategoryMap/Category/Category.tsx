import React from "react";
import {
    Link,
    LinkProps
} from "@chakra-ui/react";

import styles from "./Category.module.scss";

interface CategoryProps extends LinkProps {
    name: string
}

const Category: React.FC<CategoryProps> = (props) => {
    let {
        name = props.name[0].toUpperCase() + props.name.slice(1),
         ...rest
     } = props

    // name = props.name[0].toUpperCase() + props.name.slice(1)
    return (
        <Link className={styles.category} {...rest}> 
            {name}
        </Link>
    )
}

export default React.memo(Category);
