import {useRouter} from "next/router";
import {MovieType} from "../../../types/movies";
import useSWR from "swr";
import {get} from "../../../libs/api";
import {Box, Chip, Container, Divider, Grid, Stack, Typography} from "@mui/material";
import moment from "moment";

function Movie() {
    const router = useRouter();
    let {id} = router.query;

    const {data: movie} = useSWR<MovieType | Record<string, never>>(`../api/v1/movies/${id}`, get);

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
                        {movie.genres && movie.genres.map((g) => (
                            <Box sx={{p: 1}}>
                                <Chip key={g.id} label={g.genre}/>
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