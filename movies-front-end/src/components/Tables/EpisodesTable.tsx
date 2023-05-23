import {Button, Grid, Paper, Stack, Typography} from "@mui/material";
import Link from "next/link";
import format from "date-fns/format";
import {EpisodeType, SeasonType} from "src/types/seasons";
import {get} from "src/libs/api";
import useSWR from "swr";
import VisibilityIcon from "@mui/icons-material/Visibility";
import React from "react";


interface GridTableProps {
    season: SeasonType;
}

export function EpisodesTable({season}: GridTableProps) {
    const {data: episodes} = useSWR<EpisodeType[]>(`/api/v1/episodes?seasonId=${season.id}`, get)

    return (
        <Grid container spacing={2}>
            {episodes && episodes.map((e) => (
                <Grid key={e.id} item xs={4}>
                    <Paper
                        elevation={5}
                        sx={{
                            m: 2,
                            p: 2,
                            flexGrow: 1,
                        }}
                    >
                        <Grid container spacing={2}>
                            <Grid item xs={6}>
                                <Typography gutterBottom variant="subtitle1">
                                    <b>{e.name}</b>
                                </Typography>
                            </Grid>
                            <Grid item xs={6}>
                                <Typography variant="subtitle1">
                                    {format(new Date(e.air_date!), "MMMM do, yyyy")}
                                </Typography>
                            </Grid>
                            <Grid item xs={12}>
                                <Stack direction="row" spacing={2} sx={{display: "flex", alignItems: "center"}}>
                                    <Typography variant="subtitle2">Runtime</Typography>
                                    <Typography variant="subtitle2">
                                        {e.runtime} minutes
                                    </Typography>
                                </Stack>
                            </Grid>
                            <Grid item xs={12}>
                                <Stack direction="row" spacing={2} sx={{display: "flex", alignItems: "center"}}>
                                    <Typography variant="subtitle2">Video</Typography>

                                    {e.video_path ? (
                                        <Button
                                            variant="contained"
                                            href={`/movies/episodes/${e.id}?movieId=${season.movie_id}`}>
                                            Watch Movie <VisibilityIcon/>
                                        </Button>
                                    ) : (
                                        <Typography color="error" variant="subtitle2">
                                            Unavailable
                                        </Typography>
                                    )
                                    }
                                </Stack>
                            </Grid>
                        </Grid>
                    </Paper>
                </Grid>
            ))}
        </Grid>
    );
}