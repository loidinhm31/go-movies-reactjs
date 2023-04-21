import {useState} from "react";
import useSWR from "swr";
import {get} from "../../libs/api";
import {Box, ButtonBase, Chip, Grid, Paper, Stack, Typography,} from "@mui/material";
import {MovieType} from "../../types/movies";
import Divider from "@mui/material/Divider";
import {styled} from "@mui/material/styles";
import Link from "next/link";

const Img = styled("img")({
    margin: "auto",
    display: "block",
    maxWidth: "100%",
    maxHeight: "100%",
});

function Movies() {
    const [page, setPage] = useState(0);

    // Get all movies
    const {data: movies} = useSWR<MovieType[]>(`../api/v1/movies`, get);

    return (

        <Stack spacing={2}>
            <Box sx={{p: 1, m: 1}}>
                <Typography variant="h4">Movies</Typography>
            </Box>
            <Divider/>
            {movies &&
                <Box component="span"
                     sx={{display: "flex", justifyContent: "center", p: 1, m: 1}}>
                    <Grid container spacing={2}>
                        {movies && movies.map((movie) => (
                            <Grid key={movie.id} item xs={4}>
                                <Link href={`/movies/${movie.id}`} style={{textDecoration: "none"}}>
                                    <Paper
                                        sx={{
                                            m: 2,
                                            p: 2,
                                            flexGrow: 1,
                                        }}
                                    >
                                        <Grid container spacing={2}>
                                            <Grid item>
                                                <ButtonBase sx={{width: 128, height: 128}}>
                                                    <Img src={`https://image.tmdb.org/t/p/w200/${movie.image}`}/>
                                                </ButtonBase>
                                            </Grid>
                                            <Grid item xs container direction="column" spacing={2}>
                                                <Grid item xs>
                                                    <Typography gutterBottom variant="subtitle1" component="div">
                                                        <b>{movie.title}</b>
                                                    </Typography>
                                                    <Typography variant="body2" gutterBottom>
                                                        {new Date(movie.release_date).toDateString()}
                                                    </Typography>
                                                    <Chip label={movie.mpaa_rating} color="error"/>
                                                </Grid>
                                            </Grid>
                                        </Grid>
                                    </Paper>
                                </Link>

                            </Grid>
                        ))}
                    </Grid>
                </Box>
            }
        </Stack>
    )
}

export default Movies;