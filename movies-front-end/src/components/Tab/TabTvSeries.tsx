import {useEffect, useState} from "react";
import useSWR from "swr";
import {PageType} from "src/types/page";
import {MovieType} from "src/types/movies";
import {get} from "src/libs/api";
import {Box, Grid} from "@mui/material";
import {GridMovies} from "src/components/Tables/GridMoviesTable";
import SearchKey from "src/components/Search/SearchKey";

export function TabTvSeries() {
    const [pageIndex, setPageIndex] = useState(0);
    const [pageSize, setPageSize] = useState(9);

    const [searchKey, setSearchKey] = useState<string>("");

    // Get Tables
    const {data: page} =
        useSWR<PageType<MovieType>>(`/api/v1/movies?type=TV&q=${searchKey}&pageIndex=${pageIndex}&pageSize=${pageSize}`, get);

    // Ensure the page index has been reset when the page size changes
    useEffect(() => {
        setPageIndex(0);
    }, [pageSize])
    return (
        <Grid>
            <SearchKey
                keyword={searchKey}
                setKeyword={setSearchKey}
            />

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
        </Grid>
    );
}