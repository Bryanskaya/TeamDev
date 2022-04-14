import React, { useContext } from "react";
import Header from "components/Header";
import Search from "components/Search/Search";
import { SearchContext } from "context/Search";
import { TitlesProps } from "components/Header/Titles/Titles";


const SearchHeader: React.FC<TitlesProps> = (props) => {
    let [searchQuery, setSearchQuery] = React.useState("")
    const searchContext = useContext(SearchContext);

    const searchQueryHandler = (query: string) => {
        console.log(query)
        searchContext.searchHandler(query)
    };

    const searchField = 
    <Search value={searchQuery} onChange={(event) => {
        setSearchQuery(event.target.value)
        searchQueryHandler(event.target.value)
    }}/>

    return (
        <Header addField={searchField} {...props}/>
    );
}

export default SearchHeader;
