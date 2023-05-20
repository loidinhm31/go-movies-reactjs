import {useState} from "react";
import {Box, Divider, Stack, Tab, Tabs, Typography} from "@mui/material";
import TheatersIcon from "@mui/icons-material/Theaters";
import TvIcon from "@mui/icons-material/Tv";
import {TabMovie} from "../../components/Tab/TabMovie";
import {TabTvSeries} from "../../components/Tab/TabTvSeries";

interface TabPanelProps {
    children?: React.ReactNode;
    index: number;
    value: number;
}

function TabPanel(props: TabPanelProps) {
    const {children, value, index, ...other} = props;

    return (
        <div
            role="tabpanel"
            hidden={value !== index}
            id={`simple-tabpanel-${index}`}
            aria-labelledby={`simple-tab-${index}`}
            {...other}
        >
            {value === index && (
                <Box sx={{p: 3}}>
                    {children}
                </Box>
            )}
        </div>
    );
}


function Movies() {
    const [tabValue, setTabValue] = useState(0);

    const handleChangeTab = (event: React.SyntheticEvent, newValue: number) => {
        setTabValue(newValue);
    };

    return (
        <Stack spacing={2}>
            <Box sx={{p: 1, m: 1}}>
                <Typography variant="h4">Movies</Typography>
            </Box>
            <Divider/>

            <Tabs value={tabValue} onChange={handleChangeTab} aria-label="icon label tabs example">
                <Tab icon={<TheatersIcon/>} label="Movies"/>
                <Tab icon={<TvIcon/>} label="TV Series"/>
            </Tabs>

            <TabPanel value={tabValue} index={0}>
                <TabMovie/>
            </TabPanel>
            <TabPanel value={tabValue} index={1}>
                <TabTvSeries/>
            </TabPanel>

        </Stack>
    )
}

export default Movies;