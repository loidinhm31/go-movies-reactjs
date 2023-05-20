import {useEffect, useState} from "react";
import useSWR from "swr";
import {PageType} from "../../types/page";
import {MovieType} from "../../types/movies";
import {get} from "../../libs/api";
import {Box} from "@mui/material";
import {GridMovies} from "../Tables/GridMoviesTable";

export function TabTvSeries() {
    const [pageIndex, setPageIndex] = useState(0);
    const [pageSize, setPageSize] = useState(9);

    // Get Tables
    const {data: page} =
        useSWR<PageType<MovieType>>(`../api/v1/movies?type=TV&pageIndex=${pageIndex}&pageSize=${pageSize}`, get);

    // Ensure the page index has been reset when the page size changes
    useEffect(() => {
        setPageIndex(0);
    }, [pageSize])
    return (
        <>
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
        </>
    );
}