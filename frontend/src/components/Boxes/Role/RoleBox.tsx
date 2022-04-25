import React from "react";
import { Box, Text } from "@chakra-ui/react";

import styles from "./RoleBox.module.scss";

import DownArrowIcon from "components/Icons/DownArrow";
import ChoiceRole from "./ChoiceRole";
import { useCookies } from "react-cookie";

export interface RoleBoxProps {
  login: string;
  role: string
}
const RoleBox: React.FC<RoleBoxProps> = (props) => {
  let [cookie] = useCookies(["role"]);
  const [expanded, setExpanded] = React.useState(false);
  const [role, change] = React.useState(props.role);

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
        collapse={(newRole: string) => {setExpanded(false); change(newRole)}}
      />
    );
  }
};

export default React.memo(RoleBox);
