import {Stack} from "@mui/material";
import {useRouter} from "next/router";
import React from "react";
import WatchMovie from "src/components/Movie/WatchMovie";
import {useHasUsername} from "src/hooks/auth/useHasUsername";

function Movie() {
    const router = useRouter();
    let {id} = router.query

    const author = useHasUsername();

    return (
        <Stack spacing={2}>
            <WatchMovie
                author={author}
                movieId={parseInt(id as string)}
            />
        </Stack>
    )
}

export default Movie;