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
    let {
        name = props.name[0].toUpperCase() + props.name.slice(1),
         ...rest
     } = props

    const path = 'categories/' + props.name

    name = capitalize(props.name)
    return (
        <Link className={styles.category} href={path} {...rest}> 
            {name}
        </Link>
    )
}

export default React.memo(Category);
