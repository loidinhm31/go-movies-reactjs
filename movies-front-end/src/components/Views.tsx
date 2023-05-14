import useSWR from "swr";
import React from "react";
import {Chip} from "@mui/material";
import MovieIcon from "@mui/icons-material/Movie";
import {get} from "src/libs/api";

interface View {
    message: string;
    views: number;
}

export function Views({movieId}) {
    const {data, mutate} = useSWR<View>(`../../api/v1/views/${movieId}`, get);

    return (
        <>
            <Chip
                icon={<MovieIcon/>}
                label={`${data?.views ? data.views : 0} views`}
                color="info"
            />
        </>
    );
}