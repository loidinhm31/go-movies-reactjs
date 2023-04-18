import {useRouter} from "next/router";
import {MovieType} from "../../../types/movies";
import useSWR from "swr";
import {get} from "../../../libs/api";
import {Box, Chip, Divider, Stack} from "@mui/material";

function Movie() {
    const router = useRouter();
    let {id} = router.query;

    const {data: movie, isLoading} = useSWR<MovieType | Record<string, never>>(`../api/movies/${id}`, get);

    return (
        <Stack>
            {!isLoading &&
                <>
                    <Box sx={{display: "flex", p: 1, m: 1}}>
                        <h2>Movie: {movie.title}</h2>
                    </Box>
                    <Box>
                        <small><em>{new Date(movie.release_date).toDateString()}, </em></small>
                        <small><em>{movie.runtime} minutes, </em></small>
                        <small><em>Rated {movie.mpaa_rating}</em></small>
                    </Box>
                    <Stack direction="row">
                        {movie.genres && movie.genres.map((g) => (
                            <Box sx={{p: 1}}>
                                <Chip key={g.id} label={g.genre}/>
                            </Box>
                        ))}
                    </Stack>
                    <Divider/>
                    <Box component="span"
                         sx={{display: "flex", justifyContent: "center", p: 1, m: 1}}>
                        {movie.image !== "" &&
                            <img src={`https://image.tmdb.org/t/p/w200/${movie.image}`} alt="poster"/>
                        }
                    </Box>
                    <Box>
                        <p>{movie.description}</p>

                    </Box>
                </>
            }
        </Stack>

    )
}

export default Movie;