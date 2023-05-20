import Link from "next/link";
import useSWR from "swr";
import {get} from "../../libs/api";
import {Box, Chip, ListItem, Paper, Stack, Typography} from "@mui/material";
import Divider from "@mui/material/Divider";
import {GenreType} from "../../types/movies";
import {useState} from "react";

const GenrePaper = ({title, genres}) => {
    return (
        <Box>
            <Typography variant="overline">{title}</Typography>
            <Paper
                sx={{
                    display: "flex",
                    justifyContent: "center",
                    flexWrap: "wrap",
                    listStyle: "none",
                    p: 0.5,
                    m: 0,
                }}
                component="ul"
            >
                {genres && genres.map((g) => (
                    <Link
                        key={g.id}
                        href={`/genres/${g.id}?genreName=${g.name}`}
                        style={{textDecoration: "none"}}
                    >
                        <ListItem>
                            <Chip
                                color="info"
                                variant="filled"
                                clickable
                                label={g.name}/>
                        </ListItem>

                    </Link>
                ))}
            </Paper>
        </Box>
    );
}

function Genres() {
    const [movieGenres, setMovieGenres] = useState<GenreType[]>([]);
    const [tvGenres, setTvGenres] = useState<GenreType[]>([]);


    const {} = useSWR<GenreType[]>(`../api/v1/genres`, get, {
        onSuccess: (data) => {
            const movies: GenreType[] = [];
            const tvs: GenreType[] = [];

            data.forEach((g) => {
                if (g.type_code === "MOVIE") {
                    movies.push(g);
                } else if (g.type_code === "TV") {
                    tvs.push(g);
                }
            });

            setMovieGenres(movies);
            setTvGenres(tvs);
        }
    });

    return (
        <Stack spacing={2}>
            <Box sx={{p: 1, m: 1}}>
                <Typography variant="h4">Genres</Typography>
            </Box>
            <Divider/>
            <Box component="span"
                 sx={{display: "flex", justifyContent: "center", p: 1, m: 1}}>
                <Stack spacing={3}>
                    <GenrePaper
                        title={"Movie"}
                        genres={movieGenres}
                    />

                    <GenrePaper
                        title={"TV Series"}
                        genres={tvGenres}
                    />
                </Stack>
            </Box>
        </Stack>
    );
}

export default Genres;