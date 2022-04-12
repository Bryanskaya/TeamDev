import React from "react";
import { Box, Text, Link } from "@chakra-ui/react";
import { useCookies } from "react-cookie";

import styles from "./Navbar.module.scss";

import LoginIcon from 'components/Icons/Login'
import AuthorIcon from "components/Icons/Author";
import DownArrowIcon from "components/Icons/DownArrow";
import FullLikeIcon from "components/Icons/FullLike";
import RecipeIcon from "components/Icons/Recipes";
import LogoutIcon from "components/Icons/Logout";

import { Logout as LogoutQuery } from "postAPI/accounts/Logout";

import { navItems } from "./items";

export interface NavbarProps {}
const Navbar: React.FC<NavbarProps> = () => {
    let [cookie, , removeCookie] = useCookies(['role', 'login', 'token']);
    let role = cookie.login ? 'admin' : ''

    const [items, ] = React.useState(navItems[role]);
    const [expanded, setExpanded] = React.useState(false);
    
    const RenderNoAuth = () => {
        return (
        <Link href="/auth/signin"> <Box className={styles['user-act']}>
            <Box><LoginIcon/></Box>
            <Text>Войти</Text>
        </Box></Link>
        )
    }
    const logout = () => {
        localStorage.clear();
        LogoutQuery(removeCookie)
        window.location.href = '/';
    }

    const RenderAuth = () => {
        return (
        <Box className={styles['user-act']}>
            { expanded && <Link href="/me/likes"> <FullLikeIcon/> </Link> }
            { expanded && <Link href="/me/recipes"> <RecipeIcon/> </Link> }
            { expanded && <Link onClick={logout}> <LogoutIcon/> </Link> }

            { (!expanded) && <Box> <AuthorIcon/> </Box> }
            { (!expanded) && <Text> {cookie.login} </Text> }

            <Box onClick={() => setExpanded(!expanded)}> 
                <DownArrowIcon flipped={expanded}/> 
            </Box>
        </Box>
        )
    }

    return (
    <Box className={styles.navbar}>
        <Box className={styles.navpages}> {items.map(item =>
            <Link key={item.name} href={item.ref}> {item.name} </Link>
        )} </Box>
        { (role === '' && <RenderNoAuth />) || <RenderAuth />}
    </Box>
    )
}

export default React.memo(Navbar);
