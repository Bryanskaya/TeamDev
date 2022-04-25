import theme from "styles/extendTheme";
import React from "react";
import LogoProps from "."

const PenIcon: React.FC<LogoProps> = (props) => {
    var w = props.width ? props.width: "20px";
    var h = props.height ? props.height: "20px";
    var fill = props.fill ? props.fill : theme.colors.button;
    return (
        <svg width={w} height={h} fill="none" xmlns="http://www.w3.org/2000/svg">
            <path fill={fill} d="M0 19.9999L5.39452 19.4417L0.558452 14.6055L0 19.9999Z"/>
            <path fill={fill} d="M12.0017 3.16246L16.8378 7.99853L6.32444 18.5118L1.48828 13.6758L12.0017 3.16246ZM19.336 5.50018C20.2213 4.61491 20.2217 3.17955 19.336 2.29401L17.706 0.66402C16.8208 -0.221431 15.3852 -0.221249 14.4999 0.66402L12.923 2.24101L17.7592 7.07699L19.336 5.50018Z"/>
        </svg>
    );
}


export default PenIcon