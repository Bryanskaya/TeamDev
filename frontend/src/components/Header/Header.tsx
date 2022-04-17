import React from "react";
import {
  Box,
  Container,
} from "@chakra-ui/react";

import styles from "./Header.module.scss";

import Navbar from "./Navbar"
import Titles, { TitlesProps } from "components/Header/Titles/Titles";

export interface HeaderProps extends TitlesProps {
    role?: string
    addField?: JSX.Element
}

const Header: React.FC<HeaderProps> = (props) => {
    return (
    <Box className={styles.header}>
        <Container>
            <Navbar />
            <Titles {...props}/>
            {props.addField && props.addField}
        </Container>
    </Box>
    );
}

export default React.memo(Header);
