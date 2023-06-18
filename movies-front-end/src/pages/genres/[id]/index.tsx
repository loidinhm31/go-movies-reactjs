import { useRouter } from "next/router";
import useSWR from "swr";
import { get } from "src/libs/api";
import { GridMovies } from "src/components/Tables/GridMoviesTable";
import { useEffect, useState } from "react";
import { Box, Stack, Typography } from "@mui/material";
import Divider from "@mui/material/Divider";
import { PageType } from "src/types/page";
import { MovieType } from "src/types/movies";

function OneGenre() {
    const router = useRouter();

    const [pageIndex, setPageIndex] = useState(0);
    const [pageSize, setPageSize] = useState(9);

    // Get the id from the url
    let { id } = router.query;

    // Need to get the "prop" passed to this component
    const { genreName } = router.query;

    // Get list of Tables
    const { data: page } = useSWR<PageType<MovieType>>(
        `/api/v1/movies/genres/${id}?pageIndex=${pageIndex}&pageSize=${pageSize}`,
        get
    );

    // Ensure the page index has been reset when the page size changes
    useEffect(() => {
        setPageIndex(0);
    }, [pageSize]);

    return (
        <Stack spacing={2}>
            <Box sx={{ p: 1, m: 1 }}>
                <Typography variant="h4">Genre: {genreName}</Typography>
            </Box>
            <Divider />

            {page ? (
                <GridMovies
                    page={page}
                    pageIndex={pageIndex}
                    pageSize={pageSize}
                    setPageIndex={setPageIndex}
                    setPageSize={setPageSize}
                />
            ) : (
                <Typography>No movies in this genre (yet)!</Typography>
            )}
        </Stack>
    );
}

export default OneGenre;
