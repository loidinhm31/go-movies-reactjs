import {Stack} from "@mui/material";
import {useRouter} from "next/router";
import React, {useState} from "react";
import WatchMovie from "../../../components/Movie/WatchMovie";
import {useHasUsername} from "../../../hooks/auth/useHasUsername";

function Movie() {
    const router = useRouter();
    let {id} = router.query

    const author = useHasUsername();

    const [mutateView, setMutateView] = useState(false);

    return (
        <Stack spacing={2}>
            <WatchMovie
                mutateView={mutateView}
                setMutateView={setMutateView}
                author={author}
                movieId={parseInt(id as string)}
            />
        </Stack>
    )
}

export default Movie;