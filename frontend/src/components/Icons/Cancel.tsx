import theme from "styles/extendTheme";
import React from "react";
import LogoProps from "."

const CancelIcon: React.FC<LogoProps> = (props) => {
    var w = props.width ? props.width: "20px";
    var h = props.height ? props.height: "20px";
    var fill = props.fill ? props.fill : theme.colors.button;
    return (
        <svg width={w} height={h} fill="none" xmlns="http://www.w3.org/2000/svg">
            <path fill={fill} d="M10 0C4.47625 0 0 4.4775 0 10C0 15.5225 4.47625 20 10 20C15.5238 20 20 15.5225 20 10C20 4.4775 15.5238 0 10 0ZM14.6337 12.8663C15.1225 13.355 15.1225 14.145 14.6337 14.6337C14.39 14.8775 14.07 15 13.75 15C13.43 15 13.11 14.8775 12.8663 14.6337L10 11.7675L7.13375 14.6337C6.89 14.8775 6.57 15 6.25 15C5.93 15 5.61 14.8775 5.36625 14.6337C4.8775 14.145 4.8775 13.355 5.36625 12.8663L8.2325 10L5.36625 7.13375C4.8775 6.645 4.8775 5.855 5.36625 5.36625C5.855 4.8775 6.645 4.8775 7.13375 5.36625L10 8.2325L12.8663 5.36625C13.355 4.8775 14.145 4.8775 14.6337 5.36625C15.1225 5.855 15.1225 6.645 14.6337 7.13375L11.7675 10L14.6337 12.8663Z"/>
        </svg>
    );
}

export default CancelIcon
