import {useEffect, useState} from "react";
import useSWR from "swr";
import {get} from "../../libs/api";
import {Box, Stack, Typography,} from "@mui/material";
import {MovieType} from "../../types/movies";
import Divider from "@mui/material/Divider";
import {PageType} from "../../types/page";
import {GridMovies} from "../../components/Tables/GridMoviesTable";


function Movies() {
    const [pageIndex, setPageIndex] = useState(0);
    const [pageSize, setPageSize] = useState(9);

    // Get Tables
    const {data: page} =
        useSWR<PageType<MovieType>>(`../api/v1/movies?pageIndex=${pageIndex}&pageSize=${pageSize}`, get);

    // Ensure the page index has been reset when the page size changes
    useEffect(() => {
        setPageIndex(0);
    }, [pageSize])



    return (
        <Stack spacing={2}>
            <Box sx={{p: 1, m: 1}}>
                <Typography variant="h4">Movies</Typography>
            </Box>
            <Divider/>
            {page &&
                <Box component="span"
                     sx={{display: "flex", justifyContent: "center", p: 1, m: 1}}>
                    <GridMovies
                        page={page}
                        pageIndex={pageIndex}
                        pageSize={pageSize}
                        setPageIndex={setPageIndex}
                        setPageSize={setPageSize}
                    />

                </Box>
            }

        </Stack>
    )
}

export default Movies;