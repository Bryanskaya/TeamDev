import { Box } from "@chakra-ui/react";
import React, { useState } from "react";
import { AllRecipeResp } from "postAPI/Common"
import RecipeCard from "../RecipeCard";

import styles from "./RecipeMap.module.scss";

interface RecipeBoxProps {
    getCall: () => Promise<AllRecipeResp>
}

type State = {
    postContent?: any
}

class RecipeMap extends React.Component<RecipeBoxProps, State> {
    constructor(props) {
        super(props);
        this.state = {
            postContent: Array()
        }
    }

    async getAll() {
        var data = await this.props.getCall()
        if (data.status == 200)
            this.setState({postContent: data.content})
    }

    componentDidMount() {
        this.getAll()
    }

    render() {
        return (
            <Box 
                width="100%"
                display='flex'
                flexDirection="row"
                flexWrap="wrap"
                justifyContent="space-between"
                rowGap="40px"
                /* className={styles.map_box} */
            >
                {this.state.postContent.map(item =>
                    <RecipeCard {...item} key={item.id}/>
                )}
            </Box>
        )
    }
}

export default RecipeMap; 