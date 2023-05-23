import React, {useEffect, useState} from "react";
import {useRouter} from "next/router";
import {Box, Button, Divider, Paper, Stack, Typography} from "@mui/material";
import {useSession} from "next-auth/react";
import ManageMoviesTable from "src/components/Tables/ManageMoviesTable";
import NotifySnackbar, {NotifyState} from "src/components/shared/snackbar";
import SeasonDialog from "src/components/Dialog/SeasonDialog";
import {MovieType} from "src/types/movies";
import AddIcon from "@mui/icons-material/Add";

const ManageCatalogue = () => {
    const {data: session, status} = useSession();
    const router = useRouter();
    const [notifyState, setNotifyState] = useState<NotifyState>({open: false, vertical: "top", horizontal: "right"});

    const [selectedMovie, setSelectedMovie] = useState<MovieType | null>(null);
    const [openSeasonDialog, setOpenSeasonDialog] = useState(false);


    useEffect(() => {
        if (status === "loading") {
            return;
        }
        const role = session?.user.role;

        if (role === "admin" || role === "moderator") {
            return;
        }
        router.push("/");
    }, [router, session, status])


    return (
        <>
            <NotifySnackbar state={notifyState} setState={setNotifyState}/>
            <Stack spacing={2}>
                <Box sx={{display: "flex", p: 1, m: 1}}>
                    <Typography variant="h4">Manage Catalogue</Typography>
                </Box>
                <Divider/>
                <Paper
                    elevation={6}
                    sx={{p: 2}}
                >
                    <Box>
                        <Button
                            sx={{m: 1, p: 2}}
                            variant="contained"
                            href="/admin/movies"
                        >
                            Add Movie <AddIcon/>
                        </Button>
                    </Box>

                    <Box sx={{m: 1}}>
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
    )
}

export default ManageCatalogue;