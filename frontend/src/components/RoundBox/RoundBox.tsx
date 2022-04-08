import React from "react";
import {
    Box, BoxProps, chakra, ChakraProvider
} from "@chakra-ui/react";

import styles from "./RoundInput.module.scss";

export interface RoundBoxProps extends BoxProps {}

const RoundBox: React.FC<RoundBoxProps> = (props) => {
    return (
        <chakra.div
            className={styles.round_box}
        {...props}
        />
    )
}

export default React.memo(RoundBox)
