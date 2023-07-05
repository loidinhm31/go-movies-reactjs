import React, { useEffect, useState } from "react";
import { Box, Button, Divider, Paper, Stack, Typography } from "@mui/material";
import { signIn } from "next-auth/react";
import ManageMoviesTable from "@/components/Tables/ManageMoviesTable";
import NotifySnackbar, { NotifyState } from "@/components/shared/snackbar";
import SeasonDialog from "@/components/Dialog/SeasonDialog";
import { MovieType } from "@/types/movies";
import AddIcon from "@mui/icons-material/Add";
import { useCheckTokenAndRole } from "@/hooks/auth/useCheckTokenAndRole";
import { setData } from "@/redux/store";
import { useDispatch, useSelector } from "react-redux";

const ManageCatalogue = () => {
  const data = useSelector((state: any) => state.data);
  const dispatch = useDispatch();

  const isInvalid = useCheckTokenAndRole(["admin", "moderator"]);

  const [notifyState, setNotifyState] = useState<NotifyState>({ open: false, vertical: "top", horizontal: "right" });

  const [selectedMovie, setSelectedMovie] = useState<MovieType | null>(null);
  const [openSeasonDialog, setOpenSeasonDialog] = useState(false);

  useEffect(() => {
    if (isInvalid) {
      signIn();
      return;
    }
  }, [isInvalid]);

  useEffect(() => {
    if (data.severity !== undefined && data.message !== undefined) {
      setNotifyState({
        open: true,
        message: data.message,
        vertical: "top",
        horizontal: "right",
        severity:  data.severity,
      });

      dispatch(setData({}));
    }
  }, [data]);

  return (
    <>
      <NotifySnackbar state={notifyState} setState={setNotifyState} />
      <Stack spacing={2}>
        <Box sx={{ display: "flex", p: 1, m: 1 }}>
          <Typography variant="h4">Manage Catalogue</Typography>
        </Box>
        <Divider />
        <Paper elevation={6} sx={{ p: 2 }}>
          <Box>
            <Button sx={{ m: 1, p: 2 }} variant="contained" href="/admin/manage-catalogue/movies">
              Add Movie <AddIcon />
            </Button>
          </Box>

          <Box sx={{ m: 1 }}>
            <ManageMoviesTable
              selectedMovie={selectedMovie}
              setSelectedMovie={setSelectedMovie}
              setOpenSeasonDialog={setOpenSeasonDialog}
              setNotifyState={setNotifyState}
            />
          </Box>
        </Paper>
      </Stack>

      <SeasonDialog
        setNotifyState={setNotifyState}
        selectedMovie={selectedMovie}
        setSelectedMovie={setSelectedMovie}
        open={openSeasonDialog}
        setOpen={setOpenSeasonDialog}
      />
    </>
  );
};

export default ManageCatalogue;
