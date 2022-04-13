import {
    Box,
    HStack,
    Image,
    Link,
    Text,
    VStack,
  } from "@chakra-ui/react";
import React, { useState } from "react";

import photoRecipe from "img/photoRecipe.png"
import FullLike from "components/Icons/FullLike";
import {Recipe as RecipeI} from "types/Recipe";

import GetLikes from "postAPI/likes/Get";
import ClockBox from "components/ClockBox";

import styles from "./RecipeCard.module.scss";

interface RecipeProps extends RecipeI {}


const RecipeCard: React.FC<RecipeProps> = (props) => {
    var path = "/recipes/" + props.id;
    const [likes, change] = React.useState(0);

    async function getLikes() {
        var data = await GetLikes(props.id)
        if (data.status == 200) {
            change(data.content)
        }
    }

    getLikes();
    
    return (
        <Link className={styles.link_div}
            href={path}
        >
            <Box className={styles.main_box}>
                <Image src={photoRecipe} className={styles.image_div}
                />

                <Box className={styles.info_box}
                >
                    <VStack>
                        <Box className={styles.title_box}>
                            <Text>
                                {props.title}
                            </Text>
                        </Box>

                        <Box className={styles.description_box}>
                            <Text>
                                {props.description}
                            </Text>
                        </Box>
                    </VStack>

                    <HStack>
                        <ClockBox duration={props.duration}/>

                        <Box className={styles.likes_box}> 
                                <Box> <FullLike /> </Box>
                                <Text> {likes} </Text>
                        </Box>
                    </HStack>
                </Box>
            </Box>
        </Link>
    )
}

export default RecipeCard;