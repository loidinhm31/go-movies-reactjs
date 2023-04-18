import {Box, Stack, Typography} from "@mui/material";
import Divider from '@mui/material/Divider';

function Dashboard() {
    return (
        <Stack spacing={2}>
            <Box sx={{p: 1, m: 1}}>
                <Typography variant="h4">Dashboard</Typography>
            </Box>
            <Divider />
        </Stack>
    );
}

export default Dashboard;