import useSWRMutation from "swr/mutation";
import { del, get, post } from "@/libs/api";
import LibraryAddIcon from "@mui/icons-material/LibraryAdd";
import { Box, Button, Stack, Typography } from "@mui/material";
import { CollectionType, MovieType, PaymentType } from "@/types/movies";
import { useHasUsername } from "@/hooks/auth/useHasUsername";
import { useEffect, useState } from "react";
import LibraryAddCheckIcon from "@mui/icons-material/LibraryAddCheck";
import { useRouter } from "next/router";
import { EpisodeType } from "@/types/seasons";
import { signIn } from "next-auth/react";

interface CollectionProps {
  movie: MovieType;
  episode?: EpisodeType;
}

export function BuyCollection({ movie, episode }: CollectionProps) {
  const username = useHasUsername();

  const router = useRouter();

  const [refId, setRefId] = useState<number>();

  const [wasCollected, setWasCollected] = useState(false);
  const [wasPaid, setWasPaid] = useState(false);

  const { trigger: addCollection } = useSWRMutation("/api/v1/collections", post);
  const { trigger: deleteCollection } = useSWRMutation(
    `/api/v1/collections?type=${movie.type_code}&refId=${refId}`,
    del
  );

  const { trigger: checkBuy } = useSWRMutation(`/api/v1/payments/check?type=${movie.type_code}&refId=${refId}`, get);
  const { trigger: checkCollect } = useSWRMutation(
    `/api/v1/collections/check?type=${movie.type_code}&refId=${refId}`,
    get
  );

  useEffect(() => {
    if (movie.type_code === "MOVIE") {
      setRefId(movie.id!);
    } else if (movie.type_code === "TV") {
      setRefId(episode?.id);
    }
  }, [movie, episode]);

  useEffect(() => {
    if (refId) {
      if ((movie.price && movie.price > 0) || (episode?.price && episode?.price > 0)) {
        checkBuy().then((payment: PaymentType) => {
          if (payment.ref_id === refId && payment.type_code === movie.type_code) {
            setWasPaid(true);
            checkCollect().then((collection: CollectionType) => {
              if (movie.type_code === "MOVIE") {
                if (collection.movie_id === refId) {
                  setWasCollected(true);
                }
              } else if (movie.type_code === "TV") {
                if (collection.episode_id === refId) {
                  setWasCollected(true);
                }
              }
            });
          }
        });
      } else {
        checkCollect().then((collection: CollectionType) => {
          if (movie.type_code === "MOVIE") {
            if (collection.movie_id === refId) {
              setWasCollected(true);
              setWasPaid(true);
            }
          } else if (movie.type_code === "TV") {
            if (collection.episode_id === refId) {
              setWasCollected(true);
              setWasPaid(true);
            }
          }
        });
      }
    }
  }, [username, refId]);

  const handleClickAddToCollection = () => {
    if (username === "anonymous") {
      return signIn();
    }

    addCollection({
      movie_id: movie.id!,
      episode_id: episode?.id!,
      type_code: movie.type_code,
    } as CollectionType).then((result) => {
      if (result.message === "ok") {
        setWasPaid(true);
        setWasCollected(true);
      }
    });
  };

  const handleRemoveCollection = () => {
    deleteCollection().then((result) => {
      if (result.message === "ok") {
        setWasCollected(false);
      }
    });
  };

  const handleBuy = () => {
    router.push(`/checkout?type=${movie.type_code}&refId=${refId}`);
  };

  return (
    <>
      {wasPaid && wasCollected ? (
        <Box>
          <Button variant="contained" color="success" onClick={handleRemoveCollection}>
            <LibraryAddCheckIcon sx={{ m: 1 }} /> Collected
          </Button>
        </Box>
      ) : ((movie.type_code === "MOVIE" && movie.price) || (movie.type_code === "TV" && episode?.price)) &&
        !wasPaid &&
        !wasCollected ? (
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
            )}
          </Button>
          <Button variant="contained" color="secondary" onClick={handleBuy}>
            <LibraryAddIcon sx={{ m: 1 }} /> Buy to Watch
          </Button>
        </Stack>
      ) : (
        <Box>
          <Button variant="contained" color="secondary" onClick={handleClickAddToCollection}>
            <LibraryAddIcon sx={{ m: 1 }} /> Add to Collection
          </Button>
        </Box>
      )}
    </>
  );
}
