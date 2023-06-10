import React, {useState} from "react";
import VideoPlayer, {VideoJsOption} from "src/components/Player/VideoPlayer";
import {get} from "src/libs/api";
import useSWR from "swr";
import {EpisodeType} from "src/types/seasons";
import {MovieType} from "src/types/movies";
import {Box, CardMedia, Chip, Divider, Grid, Stack, Typography} from "@mui/material";
import {format} from "date-fns";
import {BuyCollection} from "src/components/Payment/BuyCollection";
import {Views} from "src/components/Views";
import SeasonInformation from "./SeasonInformation";

interface WatchEpisodeProps {
    author: string;
    movieId: number;
    episodeId: number;
}
export default function WatchEpisode({author, movieId, episodeId}: WatchEpisodeProps) {
    const [mutateView, setMutateView] = useState(false);

    const [videoJsOptions, setVideoJsOptions] = useState<VideoJsOption>({
        autoplay: false,
        controls: true,

    })

    const {data: movie} = useSWR<MovieType>(`/api/v1/movies/${movieId}`, get);

    const {data: episode} = useSWR<EpisodeType>(`/api/v1/episodes/${episodeId}`, get, {
        onSuccess: (result) => {
            if (result.video_path) {
                setVideoJsOptions({
                    ...videoJsOptions,
                    sources: [
                        {
                            src: `${process.env.NEXT_PUBLIC_CLOUDINARY_URL}/video/upload/${result.video_path}`,
                            type: "video/mp4",
                        }
                    ],
                });
            }
        }
    });

    return (
        <>
            {movie &&
                <>
                    <Box sx={{p: 1, m: 1}}>
                        <Typography
                            variant="h4"
                        >
                            {movie.type_code === "TV" ? "TV Series" : "Movie"} - {movie.title}
                        </Typography>
                    </Box>
                    <Box sx={{p: 2, m: 2}}>
                        {episode &&
                            <Typography
                                variant="h5"
                            >
                                {episode?.season?.name} - {episode?.name}
                            </Typography>
                        }
                    </Box>

                    <Box sx={{p: 1, m: 1}}>
                        <Typography>
                            {episode &&
                                <small><em>{format(new Date(episode?.air_date!), "MMMM do, yyyy")} | </em></small>
                            }
                            <small><em>{episode?.runtime} minutes | </em></small>
                            <small><em>Rated {movie.mpaa_rating}</em></small>
                        </Typography>
                    </Box>

                    <Stack direction="row" justifyContent="space-between">
                        <Stack direction="row" sx={{p: 1, m: 1}}>
                            {movie.genres && movie.genres.map((g, index) => (
                                <Box key={`${g.id}-${index}`} sx={{p: 1}}>
                                    <Chip key={`${g.id}-${index}`} label={g.name}/>
                                </Box>
                            ))}
                        </Stack>

                        <BuyCollection
                            movie={movie}
                            episode={episode}
                        />
                    </Stack>
                    <Divider/>

                    <Stack component="span"
                           sx={{display: "flex", justifyContent: "center", p: 1, m: 1}}>
                        <Grid container spacing={3}>
                            <Grid item xs={12} md={2}>
                                <Stack>
                                    {movie.image_url !== "" &&
                                        <Box sx={{display: "flex", justifyContent: "center", width: 1, height: 1}}>
                                            <CardMedia
                                                sx={{borderRadius: "16px"}}
                                                component="img"
                                                src={movie.image_url}
                                                alt="poster"/>
                                        </Box>
                                    }
                                </Stack>
                            </Grid>

                            <Grid item xs={12} md={10}>
                                <Stack spacing={2}>
                                    <Box sx={{m: 1}}>
                                        <Views
                                            wasMutateView={mutateView}
                                            setWasMuateView={setMutateView}
                                            movieId={movie.id!}
                                        />
                                    </Box>
                                    <Typography sx={{m: 2}} variant="body1">{movie.description}</Typography>

                                    {!episodeId && movie.type_code === "TV" &&
                                        <SeasonInformation
                                            movieId={movie.id!}
                                        />
                                    }
                                </Stack>
                            </Grid>

                        </Grid>
                    </Stack>
                </>
            }

            {videoJsOptions.sources && videoJsOptions.sources.length > 0 &&
                <VideoPlayer
                    options={videoJsOptions}
                    movieId={movieId}
                    author={author}
                    title={episode?.name!}
                    duration={episode?.runtime!}
                    setMutateView={setMutateView}
                />
            }
        </>
    );

}