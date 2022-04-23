import React from "react";
import { Box, Text } from "@chakra-ui/react";

import styles from "./RoleBox.module.scss";

import DownArrowIcon from "components/Icons/DownArrow";
import ChoiceRole from "../ChoiceRole";
import GetRole from "postAPI/accounts/GetRole";
import { useCookies } from "react-cookie";

export interface AuthActionsProps {
  login: string;
}
const RoleBox: React.FC<AuthActionsProps> = (props) => {
  let [cookie] = useCookies(["role"]);
  const [expanded, setExpanded] = React.useState(false);
  const [role, change] = React.useState("");

  async function getRole() {
    var data = await GetRole(props.login);
    if (data.status === 200) {
      change(data.content);
    }
  }

  getRole();

  if (!expanded) {
    return (
      <Box className={styles["user-act"]}>
        {!expanded && role === "admin" && <Text> {"Администратор"} </Text>}
        {!expanded && role === "user" && <Text> {"Пользователь"} </Text>}

        {cookie.role === "admin" && (
          <Box onClick={() => setExpanded(!expanded)}>
            <DownArrowIcon flipped={expanded} />
          </Box>
        )}
      </Box>
    );
  } else {
    return (
      <ChoiceRole
        login={props.login}
        role={role}
        collapse={() => setExpanded(false)}
      />
    );
  }
};

export default React.memo(RoleBox);
