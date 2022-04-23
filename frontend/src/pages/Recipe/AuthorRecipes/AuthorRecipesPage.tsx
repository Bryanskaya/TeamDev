import { Box } from "@chakra-ui/react";
import GetRecipes from "postAPI/accounts/GetRecipes";
import React from "react";
import { useCookies } from "react-cookie";
import { useParams } from "react-router-dom";
import RecipeMap from "../../../components/RecipeMap/RecipeMap";

import styles from "./AuthorRecipesPage.module.scss";

interface AuthorRecipesProps {}

const AuthorRecipes: React.FC<AuthorRecipesProps> = (props) => {
  let [cookie] = useCookies(["role", "login"]);
  const params = useParams();
  let login = params.login ? params.login : cookie.login;

  return (
    <Box className={styles.main_box}>
      <RecipeMap getCall={() => GetRecipes(login)} />
    </Box>
  );
};

export default AuthorRecipes;
