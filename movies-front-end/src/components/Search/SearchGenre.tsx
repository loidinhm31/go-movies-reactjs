import {
    Accordion,
    AccordionDetails,
    AccordionSummary,
    Autocomplete,
    CircularProgress,
    MenuItem,
    Stack,
    TextField,
    Typography
} from "@mui/material";
import {Fragment, useEffect, useState} from "react";
import {post} from "src/libs/api";
import {GenreType} from "src/types/movies";
import useSWRMutation from "swr/mutation";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";

interface SearchGenreProps {
    handleStringField: (label: string, values: string | string[], forField: string, defType: string) => void
}

export function SearchGenre({handleStringField}: SearchGenreProps) {
    const [open, setOpen] = useState(false);
    const [options, setOptions] = useState<readonly GenreType[]>([]);
    const loading = open && options.length === 0;

    const {trigger} = useSWRMutation<GenreType[]>(`../../api/v1/genres`, post);

    useEffect(() => {
        if (!loading) {
            return undefined;
        }
        console.log(options);

        if (options.length == 0) {
            trigger()
                .then((data: GenreType[]) => {
                    setOptions(data);
                })
                .catch((error) => console.log(error))
        }
    }, [loading]);

    return (
        <Accordion>
            <AccordionSummary
                expandIcon={<ExpandMoreIcon/>}
            >
                <Typography>Genres</Typography>
            </AccordionSummary>
            <AccordionDetails>
                <Stack spacing={2} direction="row">
                    <TextField
                        select
                        variant="filled"
                        sx={{minWidth: 100}}
                        id="genre-1"
                        label="Operator"
                        onChange={(event) =>
                            handleStringField("genres", event.target.value, "operator", "string")}
                    >

                        <MenuItem value={"and"}>AND</MenuItem>
                        <MenuItem value={"or"}>OR</MenuItem>
                    </TextField>

                    <TextField
                        fullWidth
                        variant="filled"
                        id="genre-2"
                        label="Field"
                        defaultValue="Genres"
                        InputProps={{
                            readOnly: true,
                        }}
                    />
                    <Autocomplete
                        fullWidth
                        open={open}
                        onOpen={() => {
                            setOpen(true);
                        }}
                        onClose={() => {
                            setOpen(false);
                        }}
                        isOptionEqualToValue={(option, value) => option.genre === value.genre}
                        options={options}
                        getOptionLabel={(option) => option.genre}
                        loading={loading}
                        onChange={(_, value) => handleStringField("genres", value.map(v => v.genre), "def", "string")}
                        multiple
                        id="genre-3"
                        renderInput={(params) => (
                            <TextField
                                {...params}
                                label="Genres"
                                InputProps={{
                                    ...params.InputProps,
                                    endAdornment: (
                                        <Fragment>
                                            {loading ? <CircularProgress color="inherit" size={20}/> : null}
                                            {params.InputProps.endAdornment}
                                        </Fragment>
                                    ),
                                }}
                            />
                        )}
                    />
                </Stack>
            </AccordionDetails>
        </Accordion>
    );
}
