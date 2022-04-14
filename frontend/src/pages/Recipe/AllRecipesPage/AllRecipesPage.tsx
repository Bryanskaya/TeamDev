import { Box } from "@chakra-ui/react";
import Search from "components/Search";
import GetRecipes from "postAPI/recipes/GetAll";
import React from "react";
import RecipeMap from "../RecipeMap/RecipeMap";

import styles from "./AllRecipesPage.module.scss";

interface AllRecipesProps {}

const AllRecipesPage: React.FC<AllRecipesProps> = (props) => {
  let [searchQuery, setSearchQuery] = React.useState("")
  return (
    <Box className={styles.main_box}>
      <Search onChange={(event) => setSearchQuery(event.target.value)}/>
      <RecipeMap searchQuery={searchQuery} getCall={GetRecipes}/>
    </Box>
  );
};

export default AllRecipesPage;
