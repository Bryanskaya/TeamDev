import React from "react";
import Header from "components/Header";
import { TitlesProps } from "components/Header/Titles/Titles";
import { useParams } from "react-router-dom";
import AuthorBox from "components/Boxes/Author";
import GetRecipe from "postAPI/recipes/Get";


const RecipeHeader: React.FC<TitlesProps> = (props) => {
    let {
        title = props.title,
        ...rest
    } = props
    const qParams = useParams()
    const [author, setAuthor] = React.useState("")

    const getAuthor = async (): Promise<void> => {
        if (!qParams.id)
            return
        
        var data = await GetRecipe(Number(qParams.id))
        if (data.status === 200)
            setAuthor(data.content.author)
    }

    getAuthor()
    const addFiled = <AuthorBox name={author} />
    return (
        <Header addField={addFiled} title={title} {...rest}/>
    );
}

export default React.memo(RecipeHeader);
