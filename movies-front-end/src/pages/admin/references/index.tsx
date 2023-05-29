import {Box, Divider, Grid, LinearProgress, Stack, TextField, Typography} from "@mui/material";
import Button from "@mui/material/Button";
import {useEffect, useState} from "react";
import useSWRMutation from "swr/mutation";
import {ReferencesTable} from "src/components/Tables/ReferencesTable";
import {post} from "src/libs/api";
import {MovieType} from "src/types/movies";
import NotifySnackbar, {NotifyState} from "src/components/shared/snackbar";
import MovieTypeSelect, {movieTypes} from "src/components/MovieTypeSelect";
import {useCheckTokenAndRole} from "src/hooks/auth/useCheckTokenAndRole";
import {signIn} from "next-auth/react";

export default function References() {
    const isInvalid = useCheckTokenAndRole(["admin", "moderator"]);

    const [searchKey, setSearchKey] = useState<string>("");
    const [progress, setProgress] = useState(0);
    const [isClickSearch, setIsClickSearch] = useState(false);

    const [notifyState, setNotifyState] = useState<NotifyState>({open: false, vertical: "top", horizontal: "right"});

    const [data, setData] = useState<MovieType[] | null>(null);

    const [selectedType, setSelectedType] = useState<string>(movieTypes[0]);

    // Get Tables
    const {trigger} =
        useSWRMutation(`/api/v1/admin/movies/references`, post);

    useEffect(() => {
        if (isInvalid) {
            signIn();
            return;
        }
    }, [isInvalid]);

    useEffect(() => {
        if (isClickSearch) {
            if (progress < 100) {
                setProgress((oldProgress) => {
                    const diff = Math.random() * 2.5;
                    return Math.min(oldProgress + diff, 100);
                })
            } else if (progress === 100) {
                trigger({
                    type_code: selectedType,
                    title: searchKey,
                } as MovieType)
                    .then((data) => setData(data))
                    .finally(() => {
                        setIsClickSearch(false);
                    })
                    .catch((error) => {
                        setNotifyState({
                            open: true,
                            message: error.message.message,
                            vertical: "top",
                            horizontal: "right",
                            severity: "error"
                        });
                    });
            }
        }
    }, [progress, isClickSearch]);
    const handleSearchClick = () => {
        if (searchKey !== "") {
            setProgress(0);
            setIsClickSearch(true);
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
                    <Typography variant="h4">The Movie Database Reference</Typography>
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
                                onKeyDown={handleKeyPressSearch}
                            />
                            {progress > 0 &&
                                <LinearProgress color="success" variant="determinate" value={progress}/>
                            }
                        </Box>

                        <Box sx={{m: 1}}>
                            <Button
                                variant="contained"
                                onClick={handleSearchClick}
                            >
                                Search
                            </Button>
                        </Box>
                    </Grid>


                    {data &&
                        <Box component="span"
                             sx={{display: "flex", justifyContent: "center", p: 1, m: 1}}>
                            <ReferencesTable
                                movieType={selectedType}
                                data={data}
                            />
                        </Box>
                    }
                </Grid>
            </Stack>

        </>
    )
}