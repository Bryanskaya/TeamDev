import { Box } from "@chakra-ui/react";
import GetRecipes from "postAPI/recipes/GetAll";
import React from "react";
import RecipeMap from "../RecipeMap/RecipeMap";

import styles from "./AllRecipesPage.module.scss";

interface AllRecipesProps {}

const AllRecipesPage: React.FC<AllRecipesProps> = (props) => {
  return (
    <Box className={styles.main_box}>
      <RecipeMap getCall={GetRecipes}/>
    </Box>
  );
};

export default AllRecipesPage;
