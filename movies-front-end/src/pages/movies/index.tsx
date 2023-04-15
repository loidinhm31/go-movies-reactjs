import {useState} from "react";
import useSWR from "swr";
import {get} from "../../libs/api";
import {Box, Stack,} from "@mui/material";
import EnhancedTable from "../../components/movies/table/MoviesTable";
import {MovieType} from "../../types/movies";

const Movies = () => {
    const [page, setPage] = useState(0);

    // Get all movies
    const {data: movies, isLoading} = useSWR<MovieType[]>(`../api/movies`, get);

    return (
        <Stack>
            <Box sx={{display: "flex", p: 1, m: 1}}>
                <h2>Movies</h2>
            </Box>
            <hr/>
            {!isLoading &&
                <Box component="span"
                     sx={{display: "flex", justifyContent: "center", p: 1, m: 1}}>
                    <EnhancedTable rows={movies}/>
                </Box>
            }
        </Stack>
    )
}

export default Movies;