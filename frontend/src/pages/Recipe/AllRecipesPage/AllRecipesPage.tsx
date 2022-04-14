import { Box } from "@chakra-ui/react";
import Search from "components/Search";
import { SearchContext } from "context/Search";
import GetRecipes from "postAPI/recipes/GetAll";
import React, { useContext } from "react";
import RecipeMap from "../RecipeMap/RecipeMap";

import styles from "./AllRecipesPage.module.scss";

interface AllRecipesProps {}

const AllRecipesPage: React.FC<AllRecipesProps> = (props) => {
  const searchContext = useContext(SearchContext);

  return (
    <Box className={styles.main_box}>
      <RecipeMap searchQuery={searchContext.query} getCall={GetRecipes}/>
    </Box>
  );
};

export default AllRecipesPage;
