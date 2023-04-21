import Link from "next/link";
import useSWR from "swr";
import {get} from "../../libs/api";
import {Box, Chip, List, ListItem, Paper, Stack, Typography} from "@mui/material";
import Divider from "@mui/material/Divider";
import {GenreType} from "../../types/movies";

function Genres() {
    const {data: genres} = useSWR<GenreType[]>(`../api/v1/genres`, get);

    return (
        <Stack spacing={2}>
            <Box sx={{p: 1, m: 1}}>
                <Typography variant="h4">Genres</Typography>
            </Box>
            <Divider/>
            <Box component="span"
                 sx={{display: "flex", justifyContent: "center", p: 1, m: 1}}>
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
                            href={`/genres/${g.id}?genreName=${g.genre}`}
                            style={{textDecoration: "none"}}
                        >
                            <ListItem>
                                <Chip
                                    color="info"
                                    variant="filled"
                                    clickable
                                    label={g.genre}/>
                            </ListItem>

                        </Link>
                    ))}
                </Paper>

            </Box>
        </Stack>
    );
}

export default Genres;