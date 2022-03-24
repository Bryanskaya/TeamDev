import { extendTheme } from "@chakra-ui/react";

const theme = extendTheme({
  colors: {
    primary: "#d4503c",
    "primary-dark": "#a52818",
    secondary: "#eee5cf",
    third: "#ddbc6d",

    "bg": "#EAE7E0",
    "bg-alt": "#D4CDC4",
    "text": "#000000",
    "title": "#9b4e23",
    "button": "#d5a129",
    "accent-1": "#5f6c11",
    "accent-2": "#29AB26",
    "accent-3": "#4D9699",
    "add-1": "#C8A98C",
    "add-2": "#BA9B00",
  },
  fonts: {
    heading: "Unica One",
    body: "Lora",
  },

  textStyles : {
    headline: {
      fontFamily: 'Lora',
      fontStyle: 'normal',
      fontWeight: 'bold',
      fontSize: '66px',
      lineHeight: '110%',
    },

    header: {
      fontFamily: 'Lora',
      fontStyle: 'normal',
      fontWeight: 'bold',
      fontSize: '41px',
      lineHeight: '140%',
    },

    subtitle: {
      fontFamily: 'Lora',
      fontStyle: 'normal',
      fontWeight: 'bold',
      fontSize: '26px',
      lineHeight: '140%',
    },

    body: {
      fontFamily: 'Lora',
      fontStyle: 'normal',
      fontWeight: 'bold',
      fontSize: '16px',
      lineHeight: '200%',
    },

    'alt-body': {
      fontFamily: 'Lora',
      fontStyle: 'normal',
      fontWeight: 'bold',
      fontSize: '16px',
      lineHeight: '20px',
      letterSpacing: '-0.045em'
    },

    caption: {
      fontFamily: 'Lora',
      fontStyle: 'normal',
      fontWeight: 'bold',
      fontSize: '12px',
      lineHeight: '200%',
      textTransform: 'uppercase',
    },
  }
});

export default theme;
