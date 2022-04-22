import React from "react";
import {
    InputProps as IProps,
    Box, Text
} from "@chakra-ui/react";
import Plate from "components/Icons/Plate";

import styles from "./PortionBox.module.scss";

interface InputProps extends IProps {
    portionNum?: number
}

const PortionBox: React.FC<InputProps> = (props) => {
    return (
    <Box className={styles.round_box_tiny}> 
        <Box> <Plate /> </Box>
        <Text> {props.portionNum} шт </Text>
    </Box>
    )
}

export default PortionBox;