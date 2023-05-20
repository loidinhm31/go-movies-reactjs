import useSWRMutation from "swr/mutation";
import {useEffect, useState} from "react";
import {post} from "src/libs/api";
import UserTable, {UserData} from "../../../components/Tables/UserTable";
import {Direction, PageType} from "../../../types/page";
import NotifySnackbar, {NotifyState} from "../../../components/shared/snackbar";
import {UserType} from "../../../types/users";
import {Box, Checkbox, Divider, Grid, Paper, Stack, TextField, Typography} from "@mui/material";
import Button from "@mui/material/Button";
import SearchUsers from "../../../components/Search/SearchUser/SearchUsers";
import SearchUsersOIDC from "../../../components/Search/SearchUser/SearchUsersOIDC";

export default function Users() {

    const [notifyState, setNotifyState] = useState<NotifyState>({open: false, vertical: "top", horizontal: "right"});



    return (
        <>
            <NotifySnackbar state={notifyState} setState={setNotifyState}/>

            <Stack spacing={2}>
                <Box sx={{p: 1, m: 1}}>
                    <Typography variant="h4">Manage User</Typography>
                </Box>
                <Divider/>


                <Paper
                    elevation={6}
                    sx={{p: 2}}
                >
                    <SearchUsersOIDC setNotifyState={setNotifyState} />
                </Paper>

                <Divider/>

                <Paper
                    elevation={6}
                    sx={{p: 2}}
                >
                    <SearchUsers setNotifyState={setNotifyState} />

                </Paper>
            </Stack>

        </>
    );
}