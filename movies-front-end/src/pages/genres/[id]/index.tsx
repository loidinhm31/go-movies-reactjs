import {useRouter} from "next/router";
import useSWR from "swr";
import {get} from "../../../libs/api";
import {GridMovies} from "../../../components/Tables/GridMoviesTable";
import {useState} from "react";
import {Box, Stack, Typography} from "@mui/material";
import Divider from "@mui/material/Divider";


function OneGenre() {
    const router = useRouter();

    const [pageIndex, setPageIndex] = useState(0);
    const [pageSize, setPageSize] = useState(9);

    // Get the id from the url
    let {id} = router.query;

    // Need to get the "prop" passed to this component
    const {genreName} = router.query;

    // Get list of Tables
    const { data: movies } = useSWR(`../api/v1/movies/genres/${id}`, get);

    return (
        <Stack spacing={2}>

            <Box sx={{p: 1, m: 1}}>
                <Typography variant="h4">Genre: {genreName}s</Typography>
            </Box>
            <Divider/>

            {movies ? (
                <GridMovies
                    page={movies}
                    pageIndex={pageIndex}
                    pageSize={pageSize}
                    setPageIndex={setPageIndex}
                    setPageSize={setPageSize}
                />

            ) : (
                <Typography>No movies in this genre (yet)!</Typography>
            )}
        </Stack>
    )
}

export default OneGenre;