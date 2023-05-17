import useSWRMutation from "swr/mutation";
import {useEffect, useState} from "react";
import {post} from "src/libs/api";
import UserTable, {UserData} from "../../../components/Tables/UserTable";
import {Direction, PageType} from "../../../types/page";
import NotifySnackbar, {NotifyState} from "../../../components/shared/snackbar";
import {UserType} from "../../../types/users";
import {Box, Checkbox, Divider, Grid, Stack, TextField, Typography} from "@mui/material";
import Button from "@mui/material/Button";

export default function Users() {

    const [isNew, setIsNew] = useState(false);
    const [searchKey, setSearchKey] = useState("");

    const [page, setPage] = useState<PageType<UserType> | null>(null);


    const [pageIndex, setPageIndex] = useState(0);
    const [pageSize, setPageSize] = useState(5)
    const [order, setOrder] = useState<Direction>(Direction.ASC);
    const [orderBy, setOrderBy] = useState<keyof UserData>("created_at");

    const [notifyState, setNotifyState] = useState<NotifyState>({open: false, vertical: "top", horizontal: "right"});


    const {trigger: requestPage} = useSWRMutation(`../../api/v1/admin/users?pageSize=${pageSize}&pageIndex=${pageIndex}&isNew=${isNew}&query=${searchKey}`, post);

    useEffect(() => {
        handeRequestPage();
    }, [pageIndex, pageSize, order, orderBy, isNew])

    // Ensure the page index has been reset when the page size changes
    useEffect(() => {
        setPageIndex(0);
    }, [pageSize])

    const handeRequestPage = () => {
        requestPage({
            sort: {
                orders: [
                    {
                        property: orderBy,
                        direction: order
                    }
                ]
            }
        }).then((data) => {
            setPage(data);
        }).catch((error) => {
            setNotifyState({
                open: true,
                message: error.message.message,
                vertical: "top",
                horizontal: "right",
                severity: "error"
            });
        });
    }

    const handleSearchClick = () => {
        if (searchKey !== "") {
            handeRequestPage();
        }
    }

    const handleKeyPressSearch = (event) => {
        if (event.key === "Enter") {
            handleSearchClick();
        }
    }

    return (
        <>
            <NotifySnackbar state={notifyState} setState={setNotifyState}/>

            <Stack spacing={2}>
                <Box sx={{p: 1, m: 1}}>
                    <Typography variant="h4">Manage User</Typography>
                </Box>
                <Divider/>


                <Grid container spacing={2} sx={{display: "flex", alignItems: "center"}}>
                    <Grid item xs={10}>
                        <TextField
                            fullWidth
                            label="Keyword"
                            variant="outlined"
                            value={searchKey}
                            onChange={e => setSearchKey(e.target.value)}
                            onKeyDown={handleKeyPressSearch}
                        />
                    </Grid>
                    <Grid item xs={2}>
                        <Stack sx={{display: "flex", alignItems: "center"}}>
                            <Typography>Is New?</Typography>
                            <Checkbox value={isNew} onChange={(event) => setIsNew(event.target.checked)}/>
                        </Stack>
                    </Grid>
                    <Grid item>
                        <Button
                            variant="contained"
                            onClick={handleSearchClick}
                        >
                            Search
                        </Button>
                    </Grid>
                </Grid>

                {page &&
                    <UserTable
                        page={page}
                        pageIndex={pageIndex}
                        setPageIndex={setPageIndex}
                        rowsPerPage={pageSize}
                        setRowsPerPage={setPageSize}
                        order={order}
                        setOrder={setOrder}
                        orderBy={orderBy}
                        setOrderBy={setOrderBy}
                        setNotifyState={setNotifyState}
                    />
                }
            </Stack>

        </>
    );
}