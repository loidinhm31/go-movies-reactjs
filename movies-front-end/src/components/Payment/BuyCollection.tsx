import useSWRMutation from "swr/mutation";
import {get, put} from "src/libs/api";
import LibraryAddIcon from "@mui/icons-material/LibraryAdd";
import {Box, Button, Stack, Typography} from "@mui/material";
import {CollectionType, MovieType} from "src/types/movies";
import {useHasUsername} from "src/hooks/auth/useHasUsername";
import {useEffect, useState} from "react";
import LibraryAddCheckIcon from "@mui/icons-material/LibraryAddCheck";
import {useRouter} from "next/router";
import {EpisodeType} from "src/types/seasons";

interface CollectionProps {
    wasAdded: boolean;
    setWasAdded: (flag1: boolean) => void;
    movie: MovieType;
    episode?: EpisodeType;
}

export function BuyCollection({wasAdded, setWasAdded, movie, episode}: CollectionProps) {
    const username = useHasUsername();

    const router = useRouter();

    const [refId, setRefId] = useState<number>();

    const {trigger: addCollection} = useSWRMutation("/api/v1/collections", put);

    const {trigger: checkBuy} = useSWRMutation(`/api/v1/collections/check?type=${movie.type_code}&refId=${refId}`, get);

    useEffect(() => {
        if (movie.type_code === "MOVIE") {
            setRefId(movie.id!);
        } else if (movie.type_code === "TV") {
            setRefId(episode?.id);
        }
    }, [movie, episode]);

    useEffect(() => {
        if (refId) {
            checkBuy().then((result: CollectionType) => {
                if (movie.type_code === "MOVIE") {
                    if (result.movie_id === refId && result.username === username) {
                        setWasAdded(true);
                    }
                } else if (movie.type_code === "TV") {
                    if (result.episode_id === refId && result.username === username) {
                        setWasAdded(true);
                    }
                }
            });
        }
    }, [username, refId]);

    const handleClickAddToCollection = () => {
        addCollection({
            movie_id: movie.id!,
            episode_id: episode?.id!,
            type_code: movie.type_code,
        } as CollectionType).then((result) => {
            if (result.message === "ok") {
                setWasAdded(true);
            }
        });
    }

    const handleBuyCollection = () => {
        router.push(`/checkout?type=${movie.type_code}&refId=${refId}`);
    }


    return (
        <>
            {wasAdded ? (
                <Box>
                    <Button
                        variant="contained"
                        color="success"
                    >
                        <LibraryAddCheckIcon sx={{m: 1}}/> Collected
                    </Button>
                </Box>
            ) : (movie.type_code === "MOVIE" && movie.price) || (movie.type_code === "TV" && episode?.price) ?
                (
                    <Stack>
                        <Button color="error">
                            {movie.type_code === "MOVIE" ? (
                                <Typography variant="overline">
                                    Price: <b>{movie.price}</b> USD
                                </Typography>
                            ) : (
                                <Typography variant="overline">
                                    Price: <b>{episode?.price}</b> USD
                                </Typography>
                            )
                            }
                        </Button>
                        <Button
                            variant="contained"
                            color="secondary"
                            onClick={handleBuyCollection}
                        >
                            <LibraryAddIcon sx={{m: 1}}/> Buy to Watch
                        </Button>
                    </Stack>

                ) : (
                    <Box>
                        <Button
                            variant="contained"
                            color="secondary"
                            onClick={handleClickAddToCollection}
                        >
                            <LibraryAddIcon sx={{m: 1}}/> Add to Collection
                        </Button>
                    </Box>
                )
            }
        </>
    );
}