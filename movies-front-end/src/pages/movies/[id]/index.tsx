import {useRouter} from "next/router";
import {MovieType} from "../../../types/movies";
import useSWR from "swr";
import {get, post} from "src/libs/api";
import {Box, Chip, Divider, Grid, Stack, Typography} from "@mui/material";
import moment from "moment";
import useSWRMutation from "swr/mutation";
import {useEffect, useState} from "react";
import {useSession} from "next-auth/react";

function Movie() {
    const router = useRouter();
    const session = useSession();
    let {id} = router.query

    const [isRecognize, setIsRecognize] = useState(false);

    const {data: movie, error} = useSWR<MovieType>(`../api/v1/movies/${id}`, get, {
        onSuccess: (result) => {
            if (!isRecognize) {
                let author;
                if (session && session.data?.user) {
                    author = session.data.user.id;
                }
                if (!author) {
                    author = "anonymous";
                }

                const request: any = {
                    movie_id: id,
                    viewer: author
                };
                trigger(request)
                    .catch((error) => console.log(error))
                    .finally(() => setIsRecognize(true));
            }
        }
    });

    const {trigger} = useSWRMutation(`../api/v1/users`, post);

    useEffect(() => {
        setIsRecognize(false);
    }, [id]);

    return (
        <Stack spacing={2}>
            {movie &&
                <>
                    <Box sx={{p: 1, m: 1}}>
                        <Typography variant="h4">Movie: {movie.title}</Typography>
                    </Box>
                    <Box sx={{p: 1, m: 1}}>
                        <Typography>
                            <small><em>{moment(movie.release_date).format("MMMM Do, yyyy")} | </em></small>
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

                    <Box component="span"
                         sx={{display: "flex", justifyContent: "center", p: 1, m: 1}}>
                        <Grid container spacing={1}>
                            <Grid item xs={12} md={2}>
                                {movie.image !== "" &&
                                    <Box sx={{display: "flex", justifyContent: "center", width: 1, height: 1}}>
                                        <img src={`https://image.tmdb.org/t/p/w200/${movie.image}`} alt="poster"/>
                                    </Box>
                                }
                            </Grid>
                            <Grid item xs={12} md={10}>
                                <Typography variant="body1">{movie.description}</Typography>
                            </Grid>
                        </Grid>
                    </Box>
                </>
            }
        </Stack>

    )
}

export default Movie;