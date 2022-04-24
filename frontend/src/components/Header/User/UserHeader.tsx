import React from "react";
import Header from "components/Header";
import { TitlesProps } from "components/Header/Titles/Titles";
import { useParams } from "react-router-dom";


const UserHeader: React.FC<TitlesProps> = (props) => {
    let {
        title = props.title,
        ...rest
    } = props

    const qParams = useParams()
    title = qParams.login ? qParams.login : ""

    return (
        <Header title={title} {...rest}/>
    );
}

export default React.memo(UserHeader);
