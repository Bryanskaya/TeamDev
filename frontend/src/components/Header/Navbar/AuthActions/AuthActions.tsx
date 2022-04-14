import React from "react";
import { Box, Text, Link } from "@chakra-ui/react";

import styles from "../Navbar.module.scss";

import AuthorIcon from "components/Icons/Author";
import DownArrowIcon from "components/Icons/DownArrow";
import FullLikeIcon from "components/Icons/FullLike";
import RecipeIcon from "components/Icons/Recipes";
import LogoutIcon from "components/Icons/Logout";

export interface AuthActionsProps {
    login: string
    logout: () => void
}
const AuthActions: React.FC<AuthActionsProps> = (props) => {
    const [expanded, setExpanded] = React.useState(false);

    return (
        <Box className={styles['user-act']}>
            { expanded && <Link href="/me/likes"> <FullLikeIcon/> </Link> }
            { expanded && <Link href="/me/recipes"> <RecipeIcon/> </Link> }
            { expanded && <Link onClick={props.logout}> <LogoutIcon/> </Link> }

            { (!expanded) && <Box> <AuthorIcon/> </Box> }
            { (!expanded) && <Text> {props.login} </Text> }

            <Box onClick={() => setExpanded(!expanded)}> 
                <DownArrowIcon flipped={expanded}/> 
            </Box>
        </Box>
    )
}

export default React.memo(AuthActions);
