import {get} from "../../libs/api";
import {MovieType} from "../../types/movies";
import {PageType} from "../../types/page";
import {GridMovies} from "../../components/Tables/GridMoviesTable";
import {useEffect, useState } from "react";
import useSWR from "swr";
import {Box, Divider, Stack, Typography} from "@mui/material";


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