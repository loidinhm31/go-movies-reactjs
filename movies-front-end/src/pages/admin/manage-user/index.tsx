import { useEffect, useState } from "react";
import NotifySnackbar, { NotifyState } from "@/components/shared/snackbar";
import { Box, Divider, Paper, Stack, Typography } from "@mui/material";
import SearchUsers from "@/components/Search/SearchUser/SearchUsers";
import SearchUsersOIDC from "@/components/Search/SearchUser/SearchUsersOIDC";
import { useCheckTokenAndRole } from "@/hooks/auth/useCheckTokenAndRole";
import { signIn } from "next-auth/react";

export default function Users() {
  const isInvalid = useCheckTokenAndRole(["admin", "moderator"]);

  const [wasUpdated, setWasUpdated] = useState(false);

  const [notifyState, setNotifyState] = useState<NotifyState>({ open: false, vertical: "top", horizontal: "right" });

  useEffect(() => {
    if (isInvalid) {
      signIn();
      return;
    }
  }, [isInvalid]);

  return (
    <>
      <NotifySnackbar state={notifyState} setState={setNotifyState} />

      <Stack spacing={2}>
        <Box sx={{ p: 1, m: 1 }}>
          <Typography variant="h4">Manage User</Typography>
        </Box>
        <Divider />

        <Paper elevation={6} sx={{ p: 2 }}>
          <SearchUsersOIDC setNotifyState={setNotifyState} wasUpdated={wasUpdated} setWasUpdated={setWasUpdated} />
        </Paper>

        <Divider />

        <Paper elevation={6} sx={{ p: 2 }}>
          <SearchUsers setNotifyState={setNotifyState} wasUpdated={wasUpdated} setWasUpdated={setWasUpdated} />
        </Paper>
      </Stack>
    </>
  );
}
