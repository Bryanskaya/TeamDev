import React from "react";
import {
    InputProps as IProps,
    Box, Text
} from "@chakra-ui/react";
import ClockIcon from "components/Icons/Clock";

import styles from "./ClockBox.module.scss";

interface InputProps extends IProps {
    duration?: number
}

const ClockBox: React.FC<InputProps> = (props) => {

      var stringDuration = ""

        if (!props.duration)
            stringDuration = "---"
        else if (props.duration < 90)
            stringDuration += props.duration + " мин"
        else if (props.duration < 60 * 24)
            stringDuration += Math.round(props.duration / 60) + " ч"
        else
            stringDuration += Math.round(props.duration / (60 * 24)) + " д"

    return (
    <Box className={styles.round_box_tiny}> 
        <Box> <ClockIcon /> </Box>
        <Text> {stringDuration} </Text>
    </Box>
    )
}

export default ClockBox;