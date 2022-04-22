import { Box } from "@chakra-ui/react";
import CategoryMap from "components/CategoryMap";
import { SearchContext } from "context/Search";
import GetCategories from "postAPI/categories/GetAll";
import GetRecipes from "postAPI/recipes/GetAll";
import React, { useContext } from "react";
import RecipeMap from "../../../components/RecipeMap/RecipeMap";

import styles from "./AllRecipesPage.module.scss";

interface AllRecipesProps {}

const AllRecipesPage: React.FC<AllRecipesProps> = (props) => {
  const searchContext = useContext(SearchContext);

  return (
    <Box className={styles.main_box}>
      <CategoryMap getCall={GetCategories}/>
      <RecipeMap searchQuery={searchContext.query} getCall={GetRecipes}/>
    </Box>
  );
};

export default AllRecipesPage;
