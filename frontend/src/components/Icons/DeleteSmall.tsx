import theme from "styles/extendTheme";
import React from "react";
import LogoProps from "."

const DeleteIcon: React.FC<LogoProps> = (props) => {
    var w = props.width ? props.width: "25px";
    var h = props.height ? props.height: "20px";
    var fill = props.fill ? props.fill : theme.colors.button;
    return (
        <svg width={w} height={h} fill="none" xmlns="http://www.w3.org/2000/svg">
            <path fill={fill} d="M16.1336 1.9591C16.1336 0.878867 15.0351 0 13.6848 0H11.3152C9.96494 0 8.86636 0.878867 8.86636 1.95914V2.29148H3.35376V6.07473H4.57627L6.66802 20H18.3318L20.4236 6.07465H21.6461V2.29145H16.1336V1.9591ZM10.3969 1.95914C10.3969 1.55406 10.8089 1.22449 11.3152 1.22449H13.6847C14.1911 1.22449 14.603 1.55406 14.603 1.95914V2.29148H10.3969V1.95914ZM16.9742 18.7755H8.02578L6.11782 6.07465H18.8821L16.9742 18.7755ZM20.1156 3.5159V4.85016H4.88437V3.5159H8.86631H16.1336H20.1156Z"/>
            <path fill={fill} d="M9.84947 7.36919L8.32104 7.43408L8.99423 17.5811L10.5227 17.5163L9.84947 7.36919Z"/>
            <path fill={fill} d="M13.2652 7.40137H11.7346V17.5483H13.2652V7.40137Z"/>
            <path fill={fill} d="M15.1505 7.36954L14.4773 17.5166L16.0057 17.5815L16.6789 7.43444L15.1505 7.36954Z"/>
        </svg>
    );
}

export default DeleteIcon