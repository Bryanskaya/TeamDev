import { Box } from "@chakra-ui/react";
import CategoryMap from "components/CategoryMap";
import { SearchContext } from "context/Search";
import GetRecipes from "postAPI/recipes/GetAll";
import React, { useContext } from "react";
import RecipeMap from "../RecipeMap/RecipeMap";

import styles from "./AllRecipesPage.module.scss";

async function GetCategories() {
  return { 
    status: 200,
    content: ['Мороженое', 'Пироги', 'Кондитерские изделия', 'Русская кухня', 'Рецепты детства', 
    'Рецепты отротчества', 'Рецепты юности']
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
