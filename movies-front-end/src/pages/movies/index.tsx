import React, { useState } from "react";
import { Box, Divider, Paper, Stack, Tab, Tabs, Typography } from "@mui/material";
import TheatersIcon from "@mui/icons-material/Theaters";
import TvIcon from "@mui/icons-material/Tv";
import { TabMovie } from "@/components/Tab/TabMovie";
import { TabTvSeries } from "@/components/Tab/TabTvSeries";
import { TabPanel } from "@/components/Tab/TabPanel";

function Movies() {
  const [tabValue, setTabValue] = useState(0);

  const handleChangeTab = (event: React.SyntheticEvent, newValue: number) => {
    setTabValue(newValue);
  };

  return (
    <Stack spacing={2}>
      <Box sx={{ p: 1, m: 1 }}>
        <Typography variant="h4">All Movies</Typography>
      </Box>
      <Divider />

      <Paper elevation={3} sx={{ p: 2 }}>
        <Tabs value={tabValue} onChange={handleChangeTab} aria-label="icon label tabs example">
          <Tab icon={<TheatersIcon />} label="Movies" />
          <Tab icon={<TvIcon />} label="TV Series" />
        </Tabs>
      </Paper>

      <TabPanel value={tabValue} index={0}>
        <TabMovie />
      </TabPanel>
      <TabPanel value={tabValue} index={1}>
        <TabTvSeries />
      </TabPanel>
    </Stack>
  );
}

export default Movies;
