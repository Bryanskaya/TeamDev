import { Box } from "@chakra-ui/react";
import CategoryMap from "components/CategoryMap";
import { SearchContext } from "context/Search";
import GetRecipes from "postAPI/recipes/GetAll";
import React, { useContext } from "react";
import RecipeMap from "../RecipeMap/RecipeMap";

import styles from "./AllRecipesPage.module.scss";

function GetCategories() {
  return { 
    status: 200,
    content: ['peepee', 'poopoo', 'Мороженое', 'Пироги', 'Кондитерские изделия', 'Русская кухня', 'Рецепты детства']
  }
}

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
