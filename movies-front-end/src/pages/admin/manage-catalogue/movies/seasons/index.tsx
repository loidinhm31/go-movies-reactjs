import React, { useEffect, useState } from "react";
import { useRouter } from "next/router";
import { signIn } from "next-auth/react";
import { del, get, post } from "@/libs/api";
import useSWRMutation from "swr/mutation";
import {
  Accordion,
  AccordionDetails,
  AccordionSummary,
  Box,
  Button,
  Divider,
  Grid,
  Stack,
  TextField,
  Typography
} from "@mui/material";
import AlertDialog from "@/components/shared/alert";
import NotifySnackbar, { NotifyState, sleep } from "@/components/shared/snackbar";
import { format } from "date-fns";
import { useCheckTokenAndRole } from "@/hooks/auth/useCheckTokenAndRole";
import { SeasonType } from "@/types/seasons";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import { ManageEpisodesTable } from "@/components/Tables/ManageEpisodesTable";
import { useDispatch } from "react-redux";
import { setData } from "@/redux/store";

const EditSeason = () => {
  const router = useRouter();
  const isInvalid = useCheckTokenAndRole(["admin", "moderator"]);
  const dispatch = useDispatch();

  const [isOpenDeleteDialog, setIsOpenDeleteDialog] = useState<boolean>(false);
  const [isConfirmDelete, setIsConfirmDelete] = useState<boolean>(false);

  const [notifyState, setNotifyState] = useState<NotifyState>({ open: false, vertical: "top", horizontal: "right" });

  // Get id from the URL
  let { id, movieId } = router.query;

  const [season, setSeason] = useState<SeasonType>({
    name: "",
    description: "",
    air_date: format(new Date(), "yyyy-MM-dd"),
    movie_id: Number(movieId)
  });

  const { trigger: fetchSeason } = useSWRMutation<SeasonType>(`/api/v1/seasons/${id}`, get);
  const { trigger: triggerSeason } = useSWRMutation(`/api/v1/admin/movies/seasons/save`, post);
  const { trigger: deleteSeason } = useSWRMutation(`/api/v1/admin/movies/seasons/delete/${id}`, del);

  useEffect(() => {
    if (isInvalid) {
      signIn();
      return;
    }
  }, [isInvalid]);

  useEffect(() => {
    if (id !== undefined) {
      fetchSeason()
        .then((result) => {
          setSeason(result!);
        })
        .catch((error) => {
          setNotifyState({
            open: true,
            message: error.message.message,
            vertical: "top",
            horizontal: "right",
            severity: "error"
          });
        });
    }
  }, [id, router]);

  useEffect(() => {
    if (isConfirmDelete) {
      deleteSeason()
        .then((data) => {
          if (data) {
            dispatch(setData({
              severity: "info",
              message: data.message
            }));

            router.push("/admin/manage-catalogue");
          }
        })
        .catch((error) => {
          setNotifyState({
            open: true,
            message: error.message.message,
            vertical: "top",
            horizontal: "right",
            severity: "error"
          });
        })
        .finally(() => {
          setIsConfirmDelete(false);
        });
    }
  }, [isConfirmDelete]);

  const handleSubmit = () => {
    let errors: any = [];
    let required = [
      { field: season.name, name: "name", label: "Name" },
      { field: season.air_date, name: "air_date", label: "Air Date" },
      { field: season.description, name: "description", label: "Description" }
    ];

    required.forEach(function({ field, label }: any) {
      if (field === "" || field === undefined) {
        errors.push(label);
      }
    });

    if (errors.length > 0) {
      setNotifyState({
        open: true,
        message: `Fill value for ${errors.join(", ")}`,
        vertical: "bottom",
        horizontal: "center",
        severity: "warning"
      });
      return false;
    }

    // Make sure movieId was set
    season.movie_id = Number(movieId);

    triggerSeason(season)
      .then((data) => {
        if (data) {
          if (!id) {
            dispatch(setData({
              severity: "success",
              message: "Season Saved"
            }));

            router.push("/admin/manage-catalogue");
          } else {
            setNotifyState({
              open: true,
              message: "Season Saved",
              vertical: "top",
              horizontal: "right",
              severity: "success"
            });
          }
        }
      })
      .catch((error) => {
        setNotifyState({
          open: true,
          message: error.message.message,
          vertical: "top",
          horizontal: "right",
          severity: "error"
        });
      });
  };

  const handleChange = (event, name: string) => {
    let value: string | number = event.target.value;
    if (name === "air_date") {
      if (Number.isNaN(new Date(event.target.value).getTime())) return;
    }

    setSeason({
      ...season,
      [name]: value
    });
  };

  const considerDelete = (event) => {
    event.preventDefault();
    setIsOpenDeleteDialog(true);
  };

  return (
    <>
      <NotifySnackbar state={notifyState} setState={setNotifyState} />

      {isOpenDeleteDialog && (
        <AlertDialog
          open={isOpenDeleteDialog}
          setOpen={setIsOpenDeleteDialog}
          title={"Delete Item"}
          description={"You cannot undo this action!"}
          confirmText={"Yes"}
          showCancelButton={true}
          setConfirmDelete={setIsConfirmDelete}
        />
      )}
      <Stack spacing={2}>
        <Box sx={{ p: 1, m: 1 }}>
          <Typography variant="h4">Add/Edit Season</Typography>
        </Box>
        <Divider />
        <Box sx={{ display: "flex", justifyContent: "center", p: 1, m: 1, width: 1 }}>
          <Grid container spacing={2}>
            <Grid item xs={8}>
              <TextField
                fullWidth
                label="Name"
                variant="outlined"
                value={season.name}
                onChange={(e) => handleChange(e, "name")}
              />
            </Grid>

            <Grid item xs={4}>
              <TextField
                fullWidth
                variant="outlined"
                label="Release Date"
                type="date"
                name="release_date"
                value={format(new Date(season.air_date!), "yyyy-MM-dd")}
                onChange={(e) => handleChange(e, "air_date")}
              />
            </Grid>

            <Grid item xs={12}>
              <TextField
                fullWidth
                label="Description"
                variant="outlined"
                multiline
                rows={4}
                value={season.description}
                onChange={(e) => handleChange(e, "description")}
              />
            </Grid>

            {id && (
              <Grid item xs={12}>
                <Accordion TransitionProps={{ unmountOnExit: true }}>
                  <AccordionSummary expandIcon={<ExpandMoreIcon />}>
                    <Typography>Episodes</Typography>
                  </AccordionSummary>
                  <AccordionDetails>
                    <ManageEpisodesTable setNotifyState={setNotifyState} season={season} />
                  </AccordionDetails>
                </Accordion>
              </Grid>
            )}

            <Divider component="div" variant="middle" />

            <Grid item xs={12}>
              <Box sx={{ display: "flex", justifyContent: "center", m: 2 }}>
                <Stack direction="row" spacing={2}>
                  <Button variant="contained" onClick={handleSubmit}>
                    Save
                  </Button>
                  {season.id! > 0 && (
                    <Button variant="contained" color="error" onClick={considerDelete}>
                      Delete Season
                    </Button>
                  )}
                </Stack>
              </Box>
            </Grid>
          </Grid>
        </Box>
      </Stack>
    </>
  );
};

export default EditSeason;
