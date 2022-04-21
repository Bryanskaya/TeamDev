import React from "react";
import Header from "components/Header";
import { TitlesProps } from "components/Header/Titles/Titles";
import { useParams } from "react-router-dom";
import { capitalize } from "functions";


const CategoryHeader: React.FC<TitlesProps> = (props) => {
    let {
        title = props.title,
        ...rest
    } = props

    const qParams = useParams()
    title = capitalize(qParams.title ? qParams.title : "")

    return (
        <Header title={title} {...rest}/>
    );
}

export default CategoryHeader;
