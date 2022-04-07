import React from "react";

import {
    Button, ButtonProps, chakra, ChakraProvider, extendTheme
} from "@chakra-ui/react";

import theme from "styles/extendTheme";
import styles from "./RoundButton.module.scss";

interface RoundButtonProps extends ButtonProps {
    focusBorderColor?: string
    focusColor?: string
    focusBgColor?: string
}

const RoundButton: React.FC<RoundButtonProps> = (props) => {
    const { 
        border="1px solid",  borderRadius="10px",
        display="flex",
        focusBorderColor = theme.colors["accent-3"],
        focusColor = theme.colors["text"],
        focusBgColor = theme.colors["bg"],
        ...rest 
    } = props

    return (
        <chakra.button 
            className={styles.round_button} 
            {...rest}
        />
    )
}

export default React.memo(RoundButton)
