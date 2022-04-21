import { Box } from "@chakra-ui/react";
import React from "react";

import styles from "./CategoryMap.module.scss";
import Category from "./Category";
import { CategoryMapProps } from ".";
import { Category as CategoryT } from "types/Categories";

type State = {
    expanded: boolean
    postContent: CategoryT[]
}

class CategoryMap extends React.Component<CategoryMapProps, State> {
    constructor(props) {
        super(props);
        this.state = {
            expanded: false,
            postContent: []
        }
    }

    async getAll() {
        var data = await this.props.getCall()
        if (data.status === 200)
            this.setState({postContent: data.content})
    }

    expandClick() {
        this.setState({expanded: !this.state.expanded})
    }

    componentDidMount() {
        this.getAll()
    }

    render() {
        return (
            <Box className={this.state.expanded ? styles.category_map_expanded : styles.category_map_collapsed}>
                <Category key='...' name='...' onClick={() => this.expandClick()}/>
                {this.state.postContent.map(item => <Category key={item.title} name={item.title}/>)}
            </Box>
        )
    }
}

export default React.memo(CategoryMap); 
