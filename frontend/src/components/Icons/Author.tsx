import theme from "styles/extendTheme";
import React from "react";
import LogoProps from "."

const AuthorIcon: React.FC<LogoProps> = (props) => {
    const {width="20px", height="20px", fill=theme.colors['accent-3']} = props
    return (
        <svg width={width} height={height} viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path fill={fill} d="M7.84211 0.296296C7.05263 0.592592 6.15789 1.38272 5.84211 2.12346C5.05263 3.90123 5.36842 7.95062 6.47368 10.321C8 13.5802 7.73684 14.2716 4.52632 15.7037C2 16.8395 0 18.5679 0 19.6049C0 19.8519 4.52632 20 10 20C15.5263 20 20 19.8025 20 19.5556C20 18.6173 18.0526 16.9383 15.5263 15.6543C12.4737 14.0741 12.1053 13.1358 13.5789 10.0741C14.2632 8.64198 14.5789 6.91358 14.5263 4.98765C14.4737 2.51852 14.2632 1.97531 13.1579 1.18519C11.4211 -1.49012e-07 9.52632 -0.296296 7.84211 0.296296Z"/>
        </svg>
    );
}

export default AuthorIcon
