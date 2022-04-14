import React from "react";
import { Box, Text, Link } from "@chakra-ui/react";

import styles from "../Navbar.module.scss";

import LoginIcon from 'components/Icons/Login'

export interface NoauthActionsProps {}
const NoauthActions: React.FC<NoauthActionsProps> = () => {
    return (
        <Link className={styles['user-act']} href="/auth/signin">
            <Box><LoginIcon/></Box>
            <Text>Войти</Text>
        </Link>
    )
}

export default React.memo(NoauthActions);
