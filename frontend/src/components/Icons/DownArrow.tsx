import theme from "styles/extendTheme";
import React from "react";
import LogoProps from "./Logo"

const DownArrowIcon: React.FC<LogoProps> = (props) => {
    const {width="24px", height="14px", fill= theme.colors['title'], ...rest} = props
    return (
        <svg width={width} height={height} fill="none" xmlns="http://www.w3.org/2000/svg">
            <path stroke={fill} d="M0.600098 1L12.0001 13L23.4001 1" strokeWidth="2"/>
        </svg>
    );
}

export default DownArrowIcon
