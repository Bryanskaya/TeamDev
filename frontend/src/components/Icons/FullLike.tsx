import theme from "styles/extendTheme";
import React from "react";
import IconProps from "."

const FullLikeIcon: React.FC<IconProps> = (props) => {
    var w = props.width ? props.width: "18px";
    var h = props.height ? props.height: "16px";
    var fill = props.fill ? props.fill : theme.colors.button;
    var visibility = props.visibility ? props.visibility : "visible";
    return (
        <svg className={props.className} width={w} height={h} viewBox="0 0 18 16" fill="none" visibility={visibility} xmlns="http://www.w3.org/2000/svg">
            <path  fill={fill} d="M2.7132 0.594747C0.700059 1.7513 -0.432336 4.29571 0.154832 6.51629C0.364534 7.44153 1.07752 8.92191 1.66469 9.80089C3.04873 11.7902 8.03965 16 9.00429 16C10.0109 16 14.4146 12.3453 16.0503 10.171C17.8957 7.7191 18.4409 5.406 17.644 3.32421C16.4278 0.132127 12.8209 -0.978163 10.3044 1.10363L9.00429 2.16765L7.87189 1.19615C6.15233 -0.145445 4.43277 -0.376757 2.7132 0.594747Z"/>
        </svg>  
    );
}

export default FullLikeIcon
