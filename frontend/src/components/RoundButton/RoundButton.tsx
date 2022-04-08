import React from "react";

import {
    ButtonProps, chakra
} from "@chakra-ui/react";

import styles from "./RoundButton.module.scss";

interface RoundButtonProps extends ButtonProps {}

const RoundButton: React.FC<RoundButtonProps> = (props) => {
    return (
        <chakra.button 
            className={styles.round_button} 
            {...props}
        />
    )
}

export default React.memo(RoundButton)
