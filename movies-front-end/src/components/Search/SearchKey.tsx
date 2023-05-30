import {Box, TextField} from "@mui/material";
import {useState} from "react";

interface SearchKeyProps {
    keyword: string;
    setKeyword: (word: string) => void;
}

export default function SearchKey({keyword, setKeyword}: SearchKeyProps) {
    return (
        <Box sx={{my: 3}}>
            <TextField
                fullWidth
                label="Keyword"
                variant="outlined"
                value={keyword}
                onChange={e => setKeyword(e.target.value)}
            />
        </Box>
    );
}