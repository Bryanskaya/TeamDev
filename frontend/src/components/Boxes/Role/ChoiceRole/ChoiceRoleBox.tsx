import React from "react";
import { Box, Link } from "@chakra-ui/react";

import styles from "./ChoiceRoleBox.module.scss";
import { ChangeRole } from "postAPI/accounts/ChangeRole";

export interface AuthActionsProps {
  login: string;
  role: string;
  collapse: (string) => void;
}

/**
 * Expanding list for changing user's role
 */
const ChoiceRoleBox: React.FC<AuthActionsProps> = (props) => {
  const roles = { admin: "Администратор", user: "Пользователь" };
  const pares = { admin: "user", user: "admin" };

  async function changeRole() {
    var data = await ChangeRole(props.login, pares[props.role]);
    if (data.status === 200) {
      props.collapse(pares[props.role]);
    }
  }

  async function close() {
    props.collapse(props.role);
  }

  return (
    <Box className={styles["user-expanded"]}>
      <Link onClick={close}>{roles[props.role]}</Link>

      <Link onClick={changeRole}>{roles[pares[props.role]]}</Link>
    </Box>
  );
};

export default React.memo(ChoiceRoleBox);
