import React, { useState } from "react";

export const SearchContext = React.createContext({
  query: "",
  searchHandler: (query: string) => {},
});

const SearchContextProvider = (props) => {
  const [query, setQuery] = useState("");

  const searchHandler = (query: string) => {
    setQuery(query);
  };

  return (
    <SearchContext.Provider
      value={{ query: query, searchHandler: searchHandler }}
    >
      {props.children}
    </SearchContext.Provider>
  );
};

export default SearchContextProvider;
