import {Grid, Stack, Typography} from "@mui/material";
import format from "date-fns/format";
import {EpisodeType} from "src/types/seasons";

interface ViewEpisodeProps {
    episode: EpisodeType;
}

export default function ViewEpisode({episode}: ViewEpisodeProps) {
    return (
        <Grid container spacing={2}>
            <Grid item xs={6}>
                <Typography gutterBottom variant="subtitle1">
                    <b>{episode.name}</b>
                </Typography>
            </Grid>
            <Grid item xs={6}>
                <Typography variant="subtitle1">
                    {format(new Date(episode.air_date!), "MMMM do, yyyy")}
                </Typography>
            </Grid>
            <Grid item xs={12}>
                <Stack direction="row" spacing={2} sx={{display: "flex", alignItems: "center"}}>
                    <Typography variant="subtitle2">Runtime</Typography>
                    <Typography variant="subtitle2">
                        {episode.runtime} minutes
                    </Typography>
                </Stack>
            </Grid>
            <Grid item xs={12}>
                <Stack direction="row" spacing={2} sx={{display: "flex", alignItems: "center"}}>
                    <Typography variant="subtitle2">Video</Typography>

                    <Typography variant="subtitle2">
                        {`${episode.video_path !== "" ? episode.video_path : "Unavailable"}`}
                    </Typography>
                </Stack>
            </Grid>
        </Grid>
    );
}