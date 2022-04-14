import { Box } from "@chakra-ui/react";
import GetRecipes from "postAPI/likes/GetRecipes";
import React from "react";
import { useCookies } from "react-cookie";
import RecipeMap from "../RecipeMap/RecipeMap";

import styles from "./LikedRecipesPage.module.scss";

interface LikedRecipesProps {}

const LikedRecipes: React.FC<LikedRecipesProps> = (props) => {
    let [cookie] = useCookies(['role', 'login']);

    return (
    <Box
        d="flex" width="100%"
        flexDir="column"
        alignItems="start"
        justifyContent="space-around"
        //className={styles.main_box}
    >
        <RecipeMap getCall={() => GetRecipes(cookie.login)}/>
    </Box>
    );
};

export default LikedRecipes;
