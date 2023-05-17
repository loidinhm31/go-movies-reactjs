import {
    Box,
    Divider,
    FormControl,
    FormControlLabel,
    FormLabel,
    LinearProgress,
    Radio,
    RadioGroup,
    Stack,
    TextField,
    Typography
} from "@mui/material";
import Button from "@mui/material/Button";
import {useEffect, useState} from "react";
import useSWRMutation from "swr/mutation";
import {ReferencesTable} from "../../../components/Tables/ReferencesTable";
import {post} from "../../../libs/api";
import {MovieType} from "../../../types/movies";
import NotifySnackbar, {NotifyState} from "src/components/shared/snackbar";


const types = ["MOVIE", "TV"] as const;

export default function References() {
    const [searchKey, setSearchKey] = useState<string>("");
    const [progress, setProgress] = useState(0);
    const [isClickSearch, setIsClickSearch] = useState(false);

    const [notifyState, setNotifyState] = useState<NotifyState>({open: false, vertical: "top", horizontal: "right"});

    const [data, setData] = useState<MovieType[] | null>(null);

    const [selectedType, setSelectedType] = useState<string>(types[0]);

    // Get Tables
    const {trigger} =
        useSWRMutation(`../api/v1/admin/movies/references`, post);


    useEffect(() => {
        if (isClickSearch) {
            if (progress < 100) {
                setProgress((oldProgress) => {
                    const diff = Math.random() * 2.5;
                    return Math.min(oldProgress + diff, 100);
                })
            } else if (progress === 100) {
                trigger({
                    type: selectedType,
                    title: searchKey,
                })
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


                <FormControl>
                    <FormLabel>Type</FormLabel>
                    <RadioGroup
                        row
                        value={selectedType}
                        onChange={(event) => setSelectedType(event.target.value)}
                    >
                        {types.map((t) => {
                            let label;
                            if (t === "MOVIE") {
                                label = "Movie";
                            } else if (t === "TV") {
                                label = "TV Series";
                            }
                            return (
                                <FormControlLabel value={t} control={<Radio/>} label={label}/>
                            );
                        })}
                    </RadioGroup>
                </FormControl>

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

                <Box>
                    <Button
                        variant="contained"
                        onClick={handleSearchClick}
                    >
                        Search
                    </Button>
                </Box>
                {data &&
                    <Box component="span"
                         sx={{display: "flex", justifyContent: "center", p: 1, m: 1}}>
                        <ReferencesTable
                            data={data}
                        />
                    </Box>
                }
            </Stack>

        </>
    )
}