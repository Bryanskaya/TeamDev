import React from "react";
import { Box, Link } from "@chakra-ui/react";

import styles from "./ChoiceRoleBox.module.scss";
import { ChangeRole } from "postAPI/accounts/ChangeRole";

export interface AuthActionsProps {
  login: string;
  role: string;
  collapse: () => void;
}
const ChoiceRoleBox: React.FC<AuthActionsProps> = (props) => {
  const [role, change] = React.useState(props.role);

  const roles = { admin: "Администратор", user: "Пользователь" };
  const pares = { admin: "user", user: "admin" };

  async function changeRole() {
    var data = await ChangeRole(props.login, pares[role]);
    if (data.status === 200) {
      change(pares[role]);
      props.collapse();
    }
  }

  return (
    <Box className={styles["user-expanded"]}>
      <Link onClick={props.collapse}>{roles[role]}</Link>

      <Link onClick={changeRole}>{roles[pares[role]]}</Link>
    </Box>
  );
};

export default React.memo(ChoiceRoleBox);
