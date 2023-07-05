import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogContent from "@mui/material/DialogContent";
import { Accordion, AccordionDetails, AccordionSummary, Grid, IconButton, Paper, Typography } from "@mui/material";
import DialogTitle from "@mui/material/DialogTitle";
import React, { useEffect, useState } from "react";
import { MovieType } from "@/types/movies";
import { get } from "@/libs/api";
import { SeasonType } from "@/types/seasons";
import format from "date-fns/format";
import AddIcon from "@mui/icons-material/Add";
import useSWRMutation from "swr/mutation";
import { NotifyState } from "@/components/shared/snackbar";
import EditIcon from "@mui/icons-material/Edit";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import { EpisodesTable } from "@/components/Tables/EpisodesTable";

interface SeasonDialogProps {
  setNotifyState: (state: NotifyState) => void;
  selectedMovie: MovieType | null;
  setSelectedMovie: (obj: MovieType | null) => void;
  open: boolean;
  setOpen: (flag: boolean) => void;
}

export default function SeasonDialog({
  setNotifyState,
  selectedMovie,
  setSelectedMovie,
  open,
  setOpen,
}: SeasonDialogProps) {
  const [seasons, setSeasons] = useState<SeasonType[]>([]);

  const { trigger: fetchSeasons } = useSWRMutation<SeasonType[]>(`/api/v1/seasons?movieId=${selectedMovie?.id}`, get);

  useEffect(() => {
    if (open) {
      if (selectedMovie?.type_code !== "TV") {
        setOpen(false);
        return;
      }

      fetchSeasons()
        .then((result) => {
          setSeasons(result!);
        })
        .catch((error) => {
          setNotifyState({
            open: true,
            message: error.message.message,
            vertical: "top",
            horizontal: "right",
            severity: "error",
          });
        });
    }
  }, [open]);

  const handleClose = () => {
    setSelectedMovie(null);
    setOpen(false);
  };

  return (
    <>
      <Dialog fullWidth={true} maxWidth={"lg"} open={open} onClose={handleClose}>
        <DialogTitle>
          <Typography >
            <b>{`TV Series - ${selectedMovie?.title}`}</b>
          </Typography>
        </DialogTitle>
        <DialogContent>
          <Paper elevation={5} sx={{ p: 1 }}>
            <Button
              sx={{ m: 2, p: 2 }}
              variant="contained"
              href={`/admin/manage-catalogue/movies/seasons?movieId=${selectedMovie?.id}`}
            >
              Add Season <AddIcon />
            </Button>

            {seasons &&
              seasons.map((s, index) => (
                <Paper
                  key={`${s.id}-${index}`}
                  elevation={3}
                  sx={{
                    m: 2,
                    p: 2,
                    flexGrow: 1,
                  }}
                >
                  <Grid container spacing={2}>
                    <Grid item xs={2}>
                      <IconButton
                        color="inherit"
                        href={`/admin/manage-catalogue/movies/seasons?id=${s.id}&movieId=${s.movie_id}`}
                      >
                        <EditIcon />
                      </IconButton>
                    </Grid>
                    <Grid item xs={6}>
                      <Typography variant="subtitle1">
                        <b>{s.name}</b>
                      </Typography>
                    </Grid>
                    <Grid item xs={4}>
                      <Typography variant="subtitle2">{format(new Date(s.air_date!), "MMMM do, yyyy")}</Typography>
                    </Grid>
                    <Grid item xs={12}>
                      <Typography variant="body1">{s.description}</Typography>
                    </Grid>
                    <Grid item xs={12}>
                      <Accordion TransitionProps={{ unmountOnExit: true }}>
                        <AccordionSummary expandIcon={<ExpandMoreIcon />}>
                          <Typography variant="caption">Episodes</Typography>
                        </AccordionSummary>
                        <AccordionDetails>
                          <EpisodesTable season={s} />
                        </AccordionDetails>
                      </Accordion>
                    </Grid>
                  </Grid>
                </Paper>
              ))}
          </Paper>
        </DialogContent>
      </Dialog>
    </>
  );
}
