import React from "react";
import {
    Button, ButtonProps
} from "@chakra-ui/react";
import theme from "extendTheme"

interface RoundBoxProps extends ButtonProps {
    focusBorderColor?: string
    focusColor?: string
    focusBgColor?: string
}

const RoundButton: React.FC<RoundBoxProps> = (props) => {
    const { 
        border="1px solid",  borderRadius="10px",
        display="flex",
        focusBorderColor = theme.colors["accent-3"],
        focusColor = theme.colors["text"],
        focusBgColor = theme.colors["bg"],
        ...rest 
    } = props
    const style = {
        color: focusColor,
        bg: focusBgColor,
        boxShadow: '0 0 5px 1px ' + focusBorderColor
    };
    return (
        <Button
            display={display}
            border={border}
            borderRadius={borderRadius}
            _hover={style}
            {...rest}
        />
    )
}

export default React.memo(RoundButton)
