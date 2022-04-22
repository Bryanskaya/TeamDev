import { Box, HStack, Image, Link, Text, VStack } from "@chakra-ui/react";
import React from "react";

import {Recipe as RecipeI} from "types/Recipe";
import {Account as AccountI} from "types/Account";

import GetLikes from "postAPI/likes/Get";
import ClockBox from "components/ClockBox";
import FullLikeBox from "components/FullLikeBox";
import LoginBox from "components/LoginBox";
import RecipeBox from "components/RecipeBox";
import user from "img/user.png";

import styles from "./UserCard.module.scss";

interface UserProps extends AccountI {}


const UserCard: React.FC<UserProps> = (props) => {
    const [likes, change] = React.useState(0);
    const pathRecipes = "/accounts/" + props.login + "/recipes";
    const pathLiked = "/accounts/" + props.login + "/like";

    async function getLikes() {
        // var data = await GetLikes(props.id)
        // if (data.status === 200) {
        //     change(data.content)
        // }
    }

    getLikes();
    
    return (
        <Box className={styles.main_box}>
            <HStack>
                <VStack>
                    <Image src={user}/>
                    <VStack className={styles.role}>
                        <Text>Роль</Text>
                        <FullLikeBox likesNum={"Понравилось"}/>
                    </VStack>
                </VStack>

                <VStack className={styles.info}>
                    <Text>Логин</Text>
                    <LoginBox login={props.login} className={styles.login}/>

                    <Link href={pathRecipes}>
                        <RecipeBox data={"Рецепты"} className={styles.recipes}/>
                    </Link>

                    <Link href={pathLiked}>
                        <FullLikeBox likesNum={"Понравилось"}/>
                    </Link>
                </VStack> 

            </HStack>
        </Box>
    )
}

export default UserCard;