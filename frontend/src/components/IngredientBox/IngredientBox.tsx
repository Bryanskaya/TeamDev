import { Box, Button, HStack, Text, VStack } from "@chakra-ui/react";
import React from "react";

import CancelIcon from "components/Icons/Cancel";
import {Ingredient as IngredientT} from "types/Ingredient";

import styles from "./IngredientBox.module.scss";

interface IngredientProps extends IngredientT {
    delCallback?: (title: string) => Promise<void>
}

function capitalizeFirstLetter(string) {
    return string.charAt(0).toUpperCase() + string.slice(1);
}

const Ingredient: React.FC<IngredientProps> = (props) => {
    const [hide, show] = React.useState(true);

    return (
        <Box className={styles.main_box} 
            onMouseEnter={() => show(false)}
            onMouseLeave={() => show(true)}> 
            <HStack>
                <VStack className={styles.info}>
                    <Text className={styles.info_title}>
                        {capitalizeFirstLetter(props.title)}
                    </Text>

                    <Text className={styles.info_amount}>
                        {props.amount}
                    </Text>
                </VStack>
                
                { props.delCallback &&
                <Button className={styles.del_btn}
                    visibility={hide ? "hidden" : "visible"} 
                    onClick={() => props.delCallback && props.delCallback(props.title)}
                >
                    <Box>
                        <CancelIcon /> 
                    </Box>
                </Button>
                }
            </HStack>
        </Box>
    )
}

export default Ingredient;