import React from "react";
import {
    Box, BoxProps
} from "@chakra-ui/react";

export interface RoundBoxProps extends BoxProps {}

const RoundBox: React.FC<RoundBoxProps> = (props) => {
    const { border="1px solid #000000", borderRadius="10px", display="flex", ...rest } = props
    
    return (
        <Box
        display={display}
        border={border}
        borderRadius={borderRadius}
        {...rest}
        />
    )
}

export default React.memo(RoundBox)
