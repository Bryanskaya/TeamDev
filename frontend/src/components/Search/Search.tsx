import React from "react";
import {
    Box,
    chakra
} from "@chakra-ui/react";

import styles from "./Search.module.scss";
import SearchIcon from "components/Icons/Search";

interface SearchProps extends React.InputHTMLAttributes<HTMLInputElement> {
    name?: string
    type?: string
    placeholder?: string
}

const Search: React.FC<SearchProps> = (props) => {
    const {
        name = "", type = "",
        placeholder = "Введите запрос",
         ...rest
     } = props

    return (
        <Box className={styles.search}>
            <SearchIcon/>
            <chakra.input
                name={name}
                type={type}
                placeholder={placeholder}
                {...rest}
            />
        </Box>
    )
}

export default React.memo(Search);