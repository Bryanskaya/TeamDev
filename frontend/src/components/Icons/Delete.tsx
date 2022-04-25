import theme from "styles/extendTheme";
import React from "react";
import LogoProps from "."

const DeleteIcon: React.FC<LogoProps> = (props) => {
    var w = props.width ? props.width: "18px";
    var h = props.height ? props.height: "16px";
    var fill = props.fill ? props.fill : theme.colors.button;
    return (
        <svg className={props.className} width={w} height={h} fill="none" xmlns="http://www.w3.org/2000/svg">
            <path fill={fill} d="M32.2674 3.9182C32.2674 1.75773 30.0702 0 27.3697 0H22.6306C19.93 0 17.7328 1.75773 17.7328 3.91828V4.58297H6.70764V12.1495H9.15266L13.3362 40H36.6638L40.8474 12.1493H43.2924V4.58289H32.2674V3.9182ZM20.794 3.91828C20.794 3.10812 21.6179 2.44898 22.6305 2.44898H27.3696C28.3822 2.44898 29.2061 3.10812 29.2061 3.91828V4.58297H20.794V3.91828ZM33.9485 37.551H16.0517L12.2358 12.1493H37.7644L33.9485 37.551ZM40.2313 7.0318V9.70031H9.76887V7.0318H17.7327H32.2674H40.2313Z"/>
            <path fill={fill} d="M19.6988 14.7389L16.642 14.8687L17.9883 35.1628L21.0452 35.033L19.6988 14.7389Z"/>
            <path fill={fill} d="M26.5306 14.8022H23.4694V35.0962H26.5306V14.8022Z"/>
            <path fill={fill} d="M30.301 14.7386L28.9546 35.0327L32.0114 35.1625L33.3578 14.8684L30.301 14.7386Z"/>
        </svg>
    );
}

export default DeleteIcon