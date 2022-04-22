import React from "react";
import {
    InputProps as IProps,
    Box, Text
} from "@chakra-ui/react";
import Recipes from "components/Icons/Recipes";

import styles from "./RecipeBox.module.scss";

interface InputProps extends IProps {
    data?: string
}

const RecipeBox: React.FC<InputProps> = (props) => {
    return (
    <Box className={styles.recipe_box}> 
        <Box> <Recipes className={styles.icon}/> </Box>
        <Text> {props.data} </Text>
    </Box>
    )
}

export default RecipeBox;