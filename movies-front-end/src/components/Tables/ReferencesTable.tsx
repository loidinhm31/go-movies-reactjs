import {Box, CardMedia, Chip, Grid, Paper, Stack, Typography} from "@mui/material";
import {styled} from "@mui/material/styles";
import format from "date-fns/format";
import Link from "next/link";
import {MovieType} from "src/types/movies";

const Img = styled("img")({
    margin: "auto",
    display: "block",
    maxWidth: "100%",
    maxHeight: "100%",
});

interface ReferencesTableProps {
    movieType: string;
    data: MovieType[];
}

export function ReferencesTable({movieType, data}: ReferencesTableProps) {

    return (
        <Grid container spacing={2}>
            {data && data.map((movie) => (
                <Grid key={movie.id} item xs={6}>
                    <Link href={`/admin/references/${movie.id}?type=${movieType}`} style={{textDecoration: "none"}}>
                        <Paper
                            sx={{
                                m: 2,
                                p: 2,
                                flexGrow: 1,
                            }}
                        >
                            <Grid container spacing={2}>
                                <Stack direction="column" spacing={3}>
                                    <Grid item xs>
                                        <Box sx={{p: 2}}>
                                            <CardMedia
                                                sx={{borderRadius: "16px"}}
                                                component="img"
                                                src={`https://image.tmdb.org/t/p/w200/${movie.image_path}`}
                                            />
                                        </Box>
                                    </Grid>
                                    <Grid item xs>
                                        <Stack direction="row" spacing={2}
                                               sx={{display: "flex", alignItems: "center", p: 2}}>
                                            <Typography>
                                                <b>Rate</b>
                                            </Typography>
                                            <Chip label={movie.vote_average} color="warning"/>
                                            <Typography gutterBottom variant="inherit" component="div">
                                                {`/ ${movie.vote_count} voters`}
                                            </Typography>
                                        </Stack>
                                    </Grid>
                                </Stack>
                                <Grid item xs container direction="column" spacing={2}>
                                    <Grid item xs>
                                        <Typography gutterBottom variant="subtitle1" component="div">
                                            <b>{movie.title}</b>
                                        </Typography>
                                        <Typography variant="body2" gutterBottom>
                                            {format(new Date(movie.release_date!), "MMMM do, yyyy")}
                                        </Typography>
                                        <Typography gutterBottom variant="inherit" component="div">
                                            {movie.description}
                                        </Typography>

                                    </Grid>
                                </Grid>
                            </Grid>
                        </Paper>
                    </Link>

                </Grid>
            ))}
        </Grid>
    );
}