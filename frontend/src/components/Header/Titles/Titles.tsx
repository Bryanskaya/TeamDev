import React from "react";
import {
  Text,
  Box,
} from "@chakra-ui/react";

import styles from "./Titles.module.scss";


export interface TitlesProps {
  subtitle?: string
  title: string
  undertitle?: string
}

const Titles: React.FC<TitlesProps> = (props) => {
  return (

    <Box className={styles.titles}>
        <Text id="subtitle" className={styles.subtitle}>
            {props.subtitle ? props.subtitle : "ㅤ"}
        </Text>
        <Text id="title" className={styles.title}>
            {props.title}
        </Text>
        <Text id="undertitle" className={styles.undertitle}>
            {props.undertitle ? props.undertitle : "ㅤ"}
        </Text>
    </Box>
  );
};

export default React.memo(Titles);
