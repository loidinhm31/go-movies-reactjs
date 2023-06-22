import { Grid, Paper, Stack } from "@mui/material";
import { EpisodeType, SeasonType } from "@/types/seasons";
import { get } from "@/libs/api";
import useSWR from "swr";
import EditEpisode from "@/components/Episode/EditEpisode";
import { NotifyState } from "@/components/shared/snackbar";
import { useEffect, useState } from "react";
import ViewEpisode from "@/components/Episode/ViewEpisode";

interface ManageEpisodesTableProps {
  season: SeasonType;
  setNotifyState: (state: NotifyState) => void;
}

export function ManageEpisodesTable({ season, setNotifyState }: ManageEpisodesTableProps) {
  const [wasUpdated, setWasUpdated] = useState(false);
  const { data: episodes, mutate } = useSWR<EpisodeType[]>(`/api/v1/episodes?seasonId=${season.id}`, get);

  useEffect(() => {
    if (wasUpdated) {
      mutate().then(() => {
        setWasUpdated(false);
      });
    }
  }, [wasUpdated]);

  return (
    <Grid container spacing={2}>
      <Grid item xs={12}>
        <Paper sx={{ p: 2 }}>
          <EditEpisode seasonId={season.id!} setNotifyState={setNotifyState} setWasUpdated={setWasUpdated} />
        </Paper>
      </Grid>

      {episodes &&
        episodes.map((e, index) => (
          <Grid key={`view-${e.id}-${index}`} item xs={12}>
            <Paper
              elevation={5}
              sx={{
                m: 2,
                p: 2,
                flexGrow: 1,
              }}
            >
              <Stack direction="row" spacing={2}>
                <ViewEpisode episode={e} />

                <Grid container spacing={2}>
                  <EditEpisode
                    key={`edit-${e.id}-${index}`}
                    id={e.id!}
                    seasonId={e.season_id!}
                    setNotifyState={setNotifyState}
                    setWasUpdated={setWasUpdated}
                  />
                </Grid>
              </Stack>
            </Paper>
          </Grid>
        ))}
    </Grid>
  );
}
