import React from "react";
import {
  Box,
  Container,
} from "@chakra-ui/react";

import styles from "./Header.module.scss";

import Navbar from "./Navbar"
import Titles, { TitlesProps } from "components/Header/Titles/Titles";

interface HeaderProps extends TitlesProps {
    role?: string
    addField?: React.FC
}

const Header: React.FC<HeaderProps> = (props) => {
    return (
    <Box className={styles.header}>
        <Container>
            <Navbar />
            <Titles {...props}/>
            {props.addField && <props.addField />}
        </Container>
    </Box>
    );
}

export default React.memo(Header);
