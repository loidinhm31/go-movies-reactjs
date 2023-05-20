import {Box, Container, Grid, Stack, Typography} from "@mui/material";
import Divider from '@mui/material/Divider';
import LineChart from "../../../components/Chart/LineChart";
import DoughnutChart from "../../../components/Chart/DoughnutChart";
import AreaChart from "../../../components/Chart/AreaChart";
import {signIn, useSession} from "next-auth/react";
import {useEffect, useState} from "react";
import MovieTypeSelect from "../../../components/MovieTypeSelect";

function Dashboard() {
    const {data: session} = useSession();

    const optionalType = ["Both"];

    const [selectedType, setSelectedType] = useState<string>(optionalType[0]);

    useEffect(() => {
        if (session?.error === "RefreshAccessTokenError") {
            signIn(); // Force sign in to hopefully resolve error
        }
    }, [session]);

    return (
        <Stack spacing={2}>
            <Box sx={{p: 1, m: 1}}>
                <Typography variant="h4">Dashboard</Typography>
            </Box>
            <Divider/>

            <Grid container>
                <Grid item xs="auto">
                    <Box sx={{m: 1}}>
                        <MovieTypeSelect
                            optionalType={optionalType}
                            selectedType={selectedType}
                            setSelectedType={setSelectedType}
                        />
                    </Box>
                </Grid>
            </Grid>

            <Grid container spacing={2}>


                <Grid item xs={6} sx={{display: "flex", justifyContent: "center", alignItems: "center"}}>
                    <Container maxWidth="sm">
                        <DoughnutChart movieType={selectedType}/>
                    </Container>
                </Grid>
                <Grid item xs={6} sx={{display: "flex", justifyContent: "center", alignItems: "center"}}>
                    <Container maxWidth="sm" >
                        <AreaChart movieType={selectedType}/>
                    </Container>
                </Grid>

                <Grid item xs={12} sx={{display: "flex", justifyContent: "center", alignItems: "center"}}>
                    <Container maxWidth="sm">
                        <LineChart movieType={selectedType}/>
                    </Container>
                </Grid>
            </Grid>
        </Stack>
    );
}

export default Dashboard;