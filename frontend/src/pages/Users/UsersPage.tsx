import { Box } from "@chakra-ui/react";
import GetUsers from "postAPI/accounts/GetAll";
import React from "react";
import { useCookies } from "react-cookie";
import UserMap from "components/UserMap";

import styles from "./UsersPage.module.scss";

interface UsersProps {}

const UsersPage: React.FC<UsersProps> = (props) => {
    let [cookie] = useCookies(['role', 'login']);

    return (
    <Box className={styles.main_box}>
        <UserMap getCall={() => GetUsers(cookie.login)}/>
    </Box>
    );
};

export default UsersPage;
