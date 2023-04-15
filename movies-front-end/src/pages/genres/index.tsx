import Link from "next/link";
import useSWR from "swr";
import {get} from "../../libs/api";
import {Box, List, ListItem, Stack} from "@mui/material";
import Divider from '@mui/material/Divider';
import {GenreType} from "../../types/movies";

const Genres = () => {
    const {data: genres} = useSWR<GenreType[]>(`../api/genres`, get);

    return (
        <Stack>
            <Box sx={{display: "flex", p: 1, m: 1}}>
                <h2>Genres</h2>
            </Box>
            <hr/>
            <Box component="span"
                 sx={{p: 1, m: 1}}>
                <List>
                    {genres && genres.map((g) => (
                        <>
                            <ListItem alignItems="flex-start" key={g.id}>
                                <Link
                                    key={g.id}
                                    className="list-group-item list-group-item-action"
                                    href={`/genres/${g.id}?genreName=${g.genre}`}
                                >{g.genre}</Link>

                            </ListItem>
                            <Divider variant="middle" component="li" />
                        </>
                    ))}
                </List>

            </Box>
        </Stack>
    );

};

export default Genres;
