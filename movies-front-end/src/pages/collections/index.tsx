import NotifySnackbar, { NotifyState } from "@/components/shared/snackbar";
import { Box, Divider, Paper, Stack, Tab, Tabs, Typography } from "@mui/material";
import { useEffect, useState } from "react";
import { CollectionMovieTab } from "@/components/Tab/CollectionMovieTab";
import { TabPanel } from "@/components/Tab/TabPanel";
import TheatersIcon from "@mui/icons-material/Theaters";
import TvIcon from "@mui/icons-material/Tv";
import { CollectionEpisodeTab } from "@/components/Tab/CollectionEpisodeTab";
import { useCheckTokenAndRole } from "@/hooks/auth/useCheckTokenAndRole";
import { signIn } from "next-auth/react";

export default function Collection() {
  const isInvalid = useCheckTokenAndRole(["admin", "moderator"]);

  const [notifyState, setNotifyState] = useState<NotifyState>({ open: false, vertical: "top", horizontal: "right" });

  const [tabValue, setTabValue] = useState(0);

  useEffect(() => {
    if (isInvalid) {
      signIn();
      return;
    }
  }, [isInvalid]);

  const handleChangeTab = (event: React.SyntheticEvent, newValue: number) => {
    setTabValue(newValue);
  };

  return (
    <>
      <NotifySnackbar state={notifyState} setState={setNotifyState} />
      <Stack spacing={2}>
        <Box sx={{ p: 1, m: 1 }}>
          <Typography variant="h4">Your Collections</Typography>
        </Box>
        <Divider />

        <Paper elevation={3} sx={{ p: 2 }}>
          <Tabs value={tabValue} onChange={handleChangeTab} aria-label="icon label tabs example">
            <Tab icon={<TheatersIcon />} label="Movies" />
            <Tab icon={<TvIcon />} label="TV Series" />
          </Tabs>
        </Paper>

        <Box component="span" sx={{ display: "flex", justifyContent: "center", p: 1, m: 1 }}>
          <TabPanel value={tabValue} index={0}>
            <CollectionMovieTab />
          </TabPanel>
          <TabPanel value={tabValue} index={1}>
            <CollectionEpisodeTab />
          </TabPanel>
        </Box>
      </Stack>
    </>
  );
}
