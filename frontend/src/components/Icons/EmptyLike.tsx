import theme from "styles/extendTheme";
import React from "react";
import LogoProps from "."

const EmptyLike: React.FC<LogoProps> = (props) => {
    var w = props.width ? props.width: "18px";
    var h = props.height ? props.height: "16px";
    var fill = props.fill ? props.fill : theme.colors.button;
    var visibility = props.visibility ? props.visibility : "visible";
    return (
        <svg className={props.className} width={w} height={h} fill="none" visibility={visibility} xmlns="http://www.w3.org/2000/svg">
            <path fill={fill} d="M6.79104 1.93718C-2.53137 7.45628 -2.24004 17.8284 7.56791 27.7247C12.4233 32.5777 23.1053 40 25.3388 40C26.1156 40 30.1942 37.5259 34.3698 34.4809C48.9361 23.8233 53.5973 13.356 47.1881 5.36283C42.8183 -0.0611229 34.0785 -1.67879 28.1549 1.93718C24.9503 3.84032 24.9503 3.84032 21.8428 1.93718C17.6672 -0.536911 10.8696 -0.536911 6.79104 1.93718ZM17.8614 2.9839C19.318 3.55484 21.5515 4.88704 22.8139 5.93377L24.9503 7.93207L27.1838 5.93377C32.6219 0.890446 40.0021 1.27108 44.8575 6.79018C50.1014 12.8802 46.5084 21.7298 35.438 30.4843C27.1838 36.955 25.4359 37.7162 22.5226 36.1937C17.473 33.5293 8.83032 26.0119 5.81996 21.5395C0.478995 13.927 2.12984 6.59987 9.8014 3.26937C13.4915 1.74686 14.3655 1.6517 17.8614 2.9839Z"/>
        </svg>
    );
}


export default EmptyLike