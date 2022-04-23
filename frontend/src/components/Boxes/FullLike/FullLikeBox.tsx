import React from "react";
import {
    InputProps as IProps,
    Box, Text
} from "@chakra-ui/react";
import FullLike from "components/Icons/FullLike";

import styles from "./FullLikeBox.module.scss";

interface InputProps extends IProps {
    likesNum?: number | string
}

const FullLikeBox: React.FC<InputProps> = (props) => {
    return (
    <Box className={styles.likes_box}> 
        <FullLike />
        <Text> {props.likesNum} </Text>
    </Box>
    )
}

export default FullLikeBox;