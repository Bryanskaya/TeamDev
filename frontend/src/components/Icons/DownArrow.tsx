import theme from "styles/extendTheme";
import React from "react";
import LogoProps from "."

const DownArrowIcon: React.FC<LogoProps> = (props) => {
    const {width="24px", height="14px", fill=theme.colors['title'], style={}, flipped=false} = props
    
    if (flipped) {
        style.transform = "rotate(180deg)"; 
    } else {
        style.transform = "rotate(0deg)";
    }
        
    return (
        <svg width={width} height={height} style={style} viewBox='0 0 24 14' fill="none" xmlns="http://www.w3.org/2000/svg">
            <path style={{fill: 'none'}} stroke={fill} d="M0.600098 1L12.0001 13L23.4001 1" strokeWidth="2"/>
        </svg>
    );
}

export default DownArrowIcon
