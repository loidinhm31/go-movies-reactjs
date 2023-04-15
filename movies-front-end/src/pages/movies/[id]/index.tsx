import {useState} from "react";
import {useRouter} from "next/router";
import {MovieType} from "../../../types/movies";
import useSWR from "swr";
import {get} from "../../../libs/api";
import {Box, Stack} from "@mui/material";

const Movie = () => {
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
                    <Box>
                        {movie.genres && movie.genres.map((g) => (
                            <span key={g.genre} className="badge bg-secondary me-2">{g.genre}</span>
                        ))}
                    </Box>
                    <hr/>
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