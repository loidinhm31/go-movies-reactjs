import {Stack} from "@mui/material";
import {useRouter} from "next/router";
import React, {useState} from "react";
import WatchMovie from "src/components/Movie/WatchMovie";
import WatchEpisode from "src/components/Movie/WatchEpisode";
import {useHasUsername} from "src/hooks/auth/useHasUsername";

function Episode() {
    const router = useRouter();
    let {id, movieId} = router.query

    const author = useHasUsername();

    const [mutateView, setMutateView] = useState(false);

    return (
        <Stack spacing={2}>
            <WatchMovie
                mutateView={mutateView}
                setMutateView={setMutateView}
                author={author}
                movieId={parseInt(movieId as string)}
                episodeId={parseInt(movieId as string)}
            />
            <WatchEpisode
                setMutateView={setMutateView}
                author={author}
                movieId={parseInt(movieId as string)}
                episodeId={parseInt(id as string)}
            />
        </Stack>
    )
}

export default Episode;