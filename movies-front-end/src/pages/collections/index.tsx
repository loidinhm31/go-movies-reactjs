import NotifySnackbar, {NotifyState} from "src/components/shared/snackbar";
import {Box, Divider, Grid, Stack, TextField, Typography} from "@mui/material";
import MovieTypeSelect, {movieTypes} from "src/components/MovieTypeSelect";
import {useEffect, useState} from "react";
import {CollectionTable} from "src/components/Tables/CollectionTable";
import {get} from "src/libs/api";
import useSWR from "swr";

export default function Collection() {
    const [searchKey, setSearchKey] = useState<string>("");

    const [notifyState, setNotifyState] = useState<NotifyState>({open: false, vertical: "top", horizontal: "right"});

    const [selectedType, setSelectedType] = useState<string>(movieTypes[0]);

    const [pageIndex, setPageIndex] = useState(0);
    const [pageSize, setPageSize] = useState(9);


    const {data: page} =
        useSWR(`/api/v1/collections?type=${selectedType}&q=${searchKey}&pageIndex=${pageIndex}&pageSize=${pageSize}`, get);

    useEffect(() => {
        setPageIndex(0);
    }, [pageSize, searchKey]);

    return (
        <>
            <NotifySnackbar state={notifyState} setState={setNotifyState}/>
            <Stack spacing={2}>
                <Box sx={{p: 1, m: 1}}>
                    <Typography variant="h4">Your Collections</Typography>
                </Box>
                <Divider/>

                <Grid container spacing={2}>
                    <Grid item xs="auto">
                        <Box sx={{m: 1}}>
                            <MovieTypeSelect
                                selectedType={selectedType}
                                setSelectedType={setSelectedType}
                            />
                        </Box>
                    </Grid>
                    <Grid item xs={12}>
                        <Box sx={{m: 1}}>
                            <TextField
                                fullWidth
                                label="Keyword"
                                variant="outlined"
                                value={searchKey}
                                onChange={e => setSearchKey(e.target.value)}
                            />

                        </Box>

                    </Grid>


                    {page &&
                        <Box component="span"
                             sx={{display: "flex", justifyContent: "center", p: 1, m: 1}}>
                            <CollectionTable
                                page={page}
                                pageIndex={pageIndex}
                                pageSize={pageSize}
                                setPageIndex={setPageIndex}
                                setPageSize={setPageSize}
                            />
                        </Box>
                    }
                </Grid>
            </Stack>

        </>
    );
}