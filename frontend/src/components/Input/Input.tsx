import React from "react";
import {
    InputProps as IProps, Input as I, ChakraProvider, chakra
} from "@chakra-ui/react";

import styles from "./RoundInput.module.scss";

interface InputProps {
    name?: string
    type?: string
    placeholder?: string
}

const Input: React.FC<InputProps> = (props) => {
    const {
        name = "", type = "",
        placeholder = "Введите данные",
         ...rest
     } = props

    return (
        <chakra.input
            name={name}
            type={type}
            className={styles.round_input}
            placeholder={placeholder}
            {...rest}
        />
    )
}

export default React.memo(Input);
