import React from "react";
import {
    InputProps as IProps, Input as I
} from "@chakra-ui/react";

import theme from "styles/extendTheme"
import RoundBox from "components/RoundBox";

interface InputProps extends IProps {
    name?: string
    type?: string
    placeholder?: string
    focusBorderColor?: string
    placeholderColor?: string
}

const Input: React.FC<InputProps> = (props) => {
    const {
        name = "", type = "",
        placeholder = "Введите данные",
        textColor = theme.colors["text"],
        placeholderColor = theme.colors["text"],
        focusBorderColor = theme.colors["accent-3"],
        onInput = () => {},
         ...rest 
     } = props

     const style = {
        color: textColor,
        boxShadow: '0 0 5px 1px ' + focusBorderColor,
      };

    return (
    <RoundBox {...rest}>
        <I 
            name={name} type={type}
            textColor={textColor}
            borderWidth={0} borderRadius={10} 
            width="100%"
            placeholder={placeholder}
            onInput={onInput}
            _focus={style}
            _placeholder={{
                color: placeholderColor,
                opacity: 0.6
            }}
            {...rest}
        />
    </RoundBox>
    )
}

export default React.memo(Input);
