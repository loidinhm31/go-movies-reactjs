import {Box, CardMedia, Chip, Divider, Grid, Stack, Typography} from "@mui/material";
import {useSession} from "next-auth/react";
import {useRouter} from "next/router";
import {useEffect, useState} from "react";
import {get, post} from "src/libs/api";
import useSWR from "swr";
import useSWRMutation from "swr/mutation";
import {MovieType} from "../../../types/movies";
import VideoPlayer, {VideoJsOption} from "src/components/Player/VideoPlayer";
import {format} from "date-fns"
import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogContent from "@mui/material/DialogContent";

function Movie() {
    const router = useRouter();
    const session = useSession();
    let {id} = router.query

    const [isRecognize, setIsRecognize] = useState(false);

    const [open, setOpen] = useState(false);

    const [videoJsOptions, setVideoJsOptions] = useState<VideoJsOption>({
        autoplay: true,
        controls: true,
    });

    const {data: movie, error} = useSWR<MovieType>(`../api/v1/movies/${id}`, get, {
        onSuccess: (result) => {
            if (result.video_path) {
                setVideoJsOptions({
                    ...videoJsOptions,
                    sources: [
                        {
                            src: `${process.env.NEXT_PUBLIC_URL}/video/upload/${result.video_path}`,
                            type: "video/mp4",
                        }
                    ],
                });
            }
        }
    });

    const {trigger} = useSWRMutation(`../api/v1/users`, post);


    useEffect(() => {
        setIsRecognize(false);
    }, [id]);

    const handleClickOpen = () => {
        setOpen(true);

        if (!isRecognize) {
            let author: string = "anonymous";
            if (session && session.data?.user) {
                author = session.data.user.id;
            }

            const request: any = {
                movie_id: id,
                viewer: author
            };
            trigger(request)
                .catch((error) => {

                })
                .finally(() => setIsRecognize(true));
        }
    };

    const handleClose = () => {
        setOpen(false);
    };

    return (
        <Stack spacing={2}>
            {movie &&
                <>
                    <Box sx={{p: 1, m: 1}}>
                        <Typography variant="h4">Movie: {movie.title}</Typography>
                    </Box>
                    <Box sx={{p: 1, m: 1}}>
                        <Typography>
                            <small><em>{format(new Date(movie.release_date!), "MMMM do, yyyy")} | </em></small>
                            <small><em>{movie.runtime} minutes | </em></small>
                            <small><em>Rated {movie.mpaa_rating}</em></small>
                        </Typography>
                    </Box>
                    <Stack direction="row" sx={{p: 1, m: 1}}>
                        {movie.genres && movie.genres.map((g, index) => (
                            <Box key={`${g.id}-${index}`} sx={{p: 1}}>
                                <Chip key={`${g.id}-${index}`} label={g.genre}/>
                            </Box>
                        ))}
                    </Stack>
                    <Divider/>

                    <Stack component="span"
                           sx={{display: "flex", justifyContent: "center", p: 1, m: 1}}>
                        <Grid container spacing={3}>
                            <Grid item xs={12} md={2}>
                                <Stack>
                                    {movie.image_path !== "" &&
                                        <Box sx={{display: "flex", justifyContent: "center", width: 1, height: 1}}>
                                            <CardMedia
                                                sx={{borderRadius: "16px"}}
                                                component="img"
                                                src={`https://image.tmdb.org/t/p/w200/${movie.image_path}`}
                                                alt="poster"/>
                                        </Box>
                                    }

                                    {movie.video_path && movie.video_path !== "" &&
                                        <Box sx={{m: 1, p: 1}}>
                                            <Button variant="contained" color="secondary" onClick={handleClickOpen}>
                                                Watch this movie
                                            </Button>
                                            <Dialog
                                                fullWidth={true}
                                                maxWidth={"lg"}
                                                open={open}
                                                onClose={handleClose}
                                            >
                                                <DialogContent>
                                                    <VideoPlayer
                                                        options={videoJsOptions}
                                                        title={movie.title}
                                                        duration={movie.runtime}/>

                                                </DialogContent>
                                            </Dialog>
                                        </Box>
                                    }
                                </Stack>
                            </Grid>
                            <Grid item xs={12} md={10}>
                                <Typography sx={{m: 2}} variant="body1">{movie.description}</Typography>
                            </Grid>
                        </Grid>

                    </Stack>
                </>
            }
        </Stack>
    )
}

export default Movie;