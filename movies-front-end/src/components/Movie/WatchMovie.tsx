import { Box, CardMedia, Chip, Divider, Grid, Stack, Typography } from "@mui/material";
import { format } from "date-fns";
import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogContent from "@mui/material/DialogContent";
import VideoPlayer, { VideoJsOption } from "src/components/Player/VideoPlayer";
import { Views } from "src/components/Views";
import SeasonInformation from "src/components/Movie/SeasonInformation";
import React, { useState } from "react";
import { get } from "src/libs/api";
import useSWR from "swr";
import { MovieType } from "src/types/movies";
import { BuyCollection } from "src/components/Payment/BuyCollection";

interface WatchMovieProps {
    author: string;
    movieId: number;
    episodeId?: number;
}

export default function WatchMovie({ author, movieId, episodeId }: WatchMovieProps) {
    const [mutateView, setMutateView] = useState(false);

    const [open, setOpen] = useState(false);

    const [videoJsOptions, setVideoJsOptions] = useState<VideoJsOption>({
        autoplay: true,
        controls: true,
    });

    const { data: movie, error } = useSWR<MovieType>(`/api/v1/movies/${movieId}`, get, {
        onSuccess: (result) => {
            if (result.video_path) {
                setVideoJsOptions({
                    ...videoJsOptions,
                    sources: [
                        {
                            src: `${process.env.NEXT_PUBLIC_CLOUDINARY_URL}/video/upload/${result.video_path}`,
                            type: "video/mp4",
                        },
                    ],
                });
            }
        },
    });

    const handleClickOpen = () => {
        setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);

        setMutateView(false);
    };

    return (
        <>
            {movie && (
                <>
                    <Box sx={{ p: 1, m: 1 }}>
                        <Typography variant="h4">
                            {movie.type_code === "TV" ? "TV Series" : "Movie"} - {movie.title}
                        </Typography>
                    </Box>

                    <Box sx={{ p: 1, m: 1 }}>
                        <Typography>
                            <small>
                                <em>{format(new Date(movie.release_date!), "MMMM do, yyyy")} | </em>
                            </small>
                            <small>
                                <em>{movie.runtime} minutes | </em>
                            </small>
                            <small>
                                <em>Rated {movie.mpaa_rating}</em>
                            </small>
                        </Typography>
                    </Box>

                    <Stack direction="row" justifyContent="space-between">
                        <Stack direction="row" sx={{ p: 1, m: 1 }}>
                            {movie.genres &&
                                movie.genres.map((g, index) => (
                                    <Box key={`${g.id}-${index}`} sx={{ p: 1 }}>
                                        <Chip key={`${g.id}-${index}`} label={g.name} />
                                    </Box>
                                ))}
                        </Stack>

                        {movie.type_code === "MOVIE" && <BuyCollection movie={movie} />}
                    </Stack>
                    <Divider />

                    <Stack component="span" sx={{ display: "flex", justifyContent: "center", p: 1, m: 1 }}>
                        <Grid container spacing={3}>
                            <Grid item xs={12} md={2}>
                                <Stack>
                                    {movie.image_url !== "" && (
                                        <Box sx={{ display: "flex", justifyContent: "center", width: 1, height: 1 }}>
                                            <CardMedia
                                                sx={{ borderRadius: "16px" }}
                                                component="img"
                                                src={movie.image_url}
                                                alt="poster"
                                            />
                                        </Box>
                                    )}

                                    {movie.video_path && movie.video_path !== "" && (
                                        <Box sx={{ m: 1, p: 1 }}>
                                            <Button variant="contained" color="secondary" onClick={handleClickOpen}>
                                                Watch this movie
                                            </Button>
                                            <Dialog fullWidth={true} maxWidth={"lg"} open={open} onClose={handleClose}>
                                                <DialogContent>
                                                    <VideoPlayer
                                                        options={videoJsOptions}
                                                        movieId={movie.id!}
                                                        author={author}
                                                        setMutateView={setMutateView}
                                                        title={movie.title}
                                                        duration={movie.runtime}
                                                    />
                                                </DialogContent>
                                            </Dialog>
                                        </Box>
                                    )}
                                </Stack>
                            </Grid>

                            <Grid item xs={12} md={10}>
                                <Stack spacing={2}>
                                    <Box sx={{ m: 1 }}>
                                        <Views
                                            wasMutateView={mutateView}
                                            setWasMuateView={setMutateView}
                                            movieId={movie.id!}
                                        />
                                    </Box>
                                    <Typography sx={{ m: 2 }} variant="body1">
                                        {movie.description}
                                    </Typography>

                                    {!episodeId && movie.type_code === "TV" && (
                                        <SeasonInformation movieId={movie.id!} />
                                    )}
                                </Stack>
                            </Grid>
                        </Grid>
                    </Stack>
                </>
            )}
        </>
    );
}
