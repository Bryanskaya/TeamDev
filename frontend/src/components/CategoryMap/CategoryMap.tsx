import { Box } from "@chakra-ui/react";
import React from "react";
import { AllCategoriesResp } from "postAPI"

import styles from "./CategoryMap.module.scss";
import Category from "./Category";

interface CategoryMapProps {
    getCall: (title?: string) => Promise<AllCategoriesResp>
}

type State = {
    postContent: string[]
}

class CategoryMap extends React.Component<CategoryMapProps, State> {
    constructor(props) {
        super(props);
        this.state = {
            postContent: []
        }
    }

    async getAll() {
        var data = await this.props.getCall()
        if (data.status === 200)
            this.setState({postContent: data.content})
    }

    componentDidMount() {
        this.getAll()
    }

    render() {
        return (
            <Box className={styles.category_map}>
                {this.state.postContent.map(item => <Category name={item}/>)}
            </Box>
        )
    }
}

export default React.memo(CategoryMap); 