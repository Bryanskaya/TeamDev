import React, { useContext } from "react";
import Header from "components/Header";
import Search from "components/Search/Search";
import { SearchContext } from "context/Search";
import { TitlesProps } from "components/Header/Titles/Titles";


const SearchHeader: React.FC<TitlesProps> = (props) => {
    let [searchQuery, setSearchQuery] = React.useState("")
    const searchContext = useContext(SearchContext);

    const changeHandler = (event: React.ChangeEvent<HTMLInputElement>) => {
        setSearchQuery(event.target.value)
        searchContext.searchHandler(event.target.value)
    };

    const searchField = <Search value={searchQuery} onChange={(event) => changeHandler(event)}/>

    return (
        <Header addField={searchField} {...props}/>
    );
}

export default SearchHeader;
