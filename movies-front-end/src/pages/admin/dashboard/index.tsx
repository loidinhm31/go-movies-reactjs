import { Box, Container, Grid, Stack, Typography } from "@mui/material";
import Divider from "@mui/material/Divider";
import LineChart from "src/components/Chart/LineChart";
import DoughnutChart from "src/components/Chart/DoughnutChart";
import AreaChart from "src/components/Chart/AreaChart";
import { signIn } from "next-auth/react";
import { useEffect, useState } from "react";
import MovieTypeSelect from "src/components/MovieTypeSelect";
import { useCheckTokenAndRole } from "src/hooks/auth/useCheckTokenAndRole";
import PaymentBarChart from "src/components/Chart/PaymentBarChart";

function Dashboard() {
    const isInvalid = useCheckTokenAndRole(["admin", "moderator"]);

    const optionalType = ["Both"];

    const [selectedType, setSelectedType] = useState<string>(optionalType[0]);

    useEffect(() => {
        if (isInvalid) {
            signIn();
            return;
        }
    }, [isInvalid]);

    return (
        <Stack spacing={2}>
            <Box sx={{ p: 1, m: 1 }}>
                <Typography variant="h4">Dashboard</Typography>
            </Box>
            <Divider />

            <Grid container>
                <Grid item xs="auto">
                    <Box sx={{ m: 1 }}>
                        <MovieTypeSelect
                            optionalType={optionalType}
                            selectedType={selectedType}
                            setSelectedType={setSelectedType}
                        />
                    </Box>
                </Grid>
            </Grid>

            <Grid container spacing={2}>
                <Grid item xs={6} sx={{ display: "flex", justifyContent: "center", alignItems: "center" }}>
                    <Container maxWidth="sm">
                        <DoughnutChart movieType={selectedType} />
                    </Container>
                </Grid>
                <Grid item xs={6} sx={{ display: "flex", justifyContent: "center", alignItems: "center" }}>
                    <Container maxWidth="sm">
                        <AreaChart movieType={selectedType} />
                    </Container>
                </Grid>

                <Grid item xs={6} sx={{ display: "flex", justifyContent: "center", alignItems: "center" }}>
                    <Container maxWidth="sm">
                        <LineChart movieType={selectedType} />
                    </Container>
                </Grid>

                <Grid item xs={6} sx={{ display: "flex", justifyContent: "center", alignItems: "center" }}>
                    <Container maxWidth="sm" sx={{ p: 2 }}>
                        <PaymentBarChart movieType={selectedType} />
                    </Container>
                </Grid>
            </Grid>
        </Stack>
    );
}

export default Dashboard;
