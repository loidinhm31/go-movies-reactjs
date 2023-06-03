import useSWRMutation from "swr/mutation";
import {get, put} from "src/libs/api";
import LibraryAddIcon from "@mui/icons-material/LibraryAdd";
import {Box, Button, Stack, Typography} from "@mui/material";
import {CollectionType, MovieType} from "src/types/movies";
import {useHasUsername} from "src/hooks/auth/useHasUsername";
import {useEffect, useState} from "react";
import LibraryAddCheckIcon from "@mui/icons-material/LibraryAddCheck";
import {useRouter} from "next/router";

interface CollectionProps {
    movie: MovieType;
}

export function BuyCollection({movie}: CollectionProps) {
    const username = useHasUsername();

    const router = useRouter();

    const {trigger: addMovie} = useSWRMutation("/api/v1/collections", put);

    const {trigger: checkBuy} = useSWRMutation(`/api/v1/collections/check?username=${username}&movieId=${movie.id!}`, get);

    const [wasAdded, setWasAdded] = useState(false);

    useEffect(() => {
        checkBuy().then((result) => {
            if (result.movie_id === movie.id! && result.username === username) {
                setWasAdded(true);
            }
        });
    }, [username]);

    const handleClickAddToCollection = () => {
        addMovie({
            movie_id: movie.id!,
        } as CollectionType).then((result) => {
            if (result.message === "ok") {
                setWasAdded(true);
            }
        });
    }

    const handleBuyCollection = () => {
        router.push(`/checkout?movieId=${movie.id!}`);
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
            ) : movie.price ?
                (
                    <Stack>
                        <Button color="error">
                            <Typography variant="overline">
                                Price: <b>{movie.price}</b> USD
                            </Typography>
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
    )
        ;
}