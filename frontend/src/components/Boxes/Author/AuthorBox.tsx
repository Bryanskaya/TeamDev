import React from "react";
import {
    chakra
} from "@chakra-ui/react";

import styles from "./AuthorBox.module.scss";
import AuthorIcon from "components/Icons/Author";

interface AuthorBoxProps extends React.InputHTMLAttributes<HTMLLabelElement> {
    name: string
}

/**
 * Author of recipe link box
 */
const AuthorBox: React.FC<AuthorBoxProps> = (props) => {
    const { name = "", ...rest } = props

    return (
        <chakra.a className={styles.author_box} href={`/accounts/${name}/recipes`}>
            <AuthorIcon/>
            <chakra.label {...rest}> 
                {name}
            </chakra.label>
        </chakra.a>
    )
}

export default React.memo(AuthorBox);