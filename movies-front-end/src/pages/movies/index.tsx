import {useState} from "react";
import useSWR from "swr";
import {get} from "../../libs/api";
import {Box, Container, Stack, Typography,} from "@mui/material";
import EnhancedTable from "../../components/movies/table/MoviesTable";
import {MovieType} from "../../types/movies";
import Divider from "@mui/material/Divider";

function Movies() {
    const [page, setPage] = useState(0);

    // Get all movies
    const {data: movies, isLoading} = useSWR<MovieType[]>(`../api/movies`, get);

    return (

        <Stack spacing={2}>
            <Box sx={{p: 1, m: 1}}>
                <Typography variant="h4">Movies</Typography>
            </Box>
            <Divider/>
            {!isLoading &&
                <Box component="span"
                     sx={{display: "flex", justifyContent: "center", p: 1, m: 1}}>
                    {movies && <EnhancedTable rows={movies}/>}
                </Box>
            }
        </Stack>
    )
}

export default Movies;