import {
    Accordion,
    AccordionDetails,
    AccordionSummary,
    Autocomplete,
    CircularProgress,
    MenuItem,
    Stack,
    TextField,
    Typography,
} from "@mui/material";
import { Fragment, useEffect, useState } from "react";
import { post } from "src/libs/api";
import { GenreType } from "src/types/movies";
import useSWRMutation from "swr/mutation";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import { useMovieType } from "src/hooks/useMovieType";

interface SearchGenreProps {
    movieType: string;
    handleStringField: (label: string, values: string | string[], forField: string, defType: string) => void;
}

export function SearchGenre({ movieType, handleStringField }: SearchGenreProps) {
    const [open, setOpen] = useState(false);
    const [options, setOptions] = useState<readonly GenreType[]>([]);

    const selectedType = useMovieType(movieType);

    const { trigger } = useSWRMutation<GenreType[]>(`/api/v1/genres?type=${selectedType}`, post);

    useEffect(() => {
        if (!open) {
            return;
        }

        trigger()
            .then((data: GenreType[]) => {
                setOptions(data);
            })
            .catch((error) => console.log(error));
    }, [open, movieType]);

    return (
        <Accordion>
            <AccordionSummary expandIcon={<ExpandMoreIcon />}>
                <Typography>Genres</Typography>
            </AccordionSummary>
            <AccordionDetails>
                <Stack spacing={2} direction="row">
                    <TextField
                        select
                        variant="filled"
                        sx={{ minWidth: 100 }}
                        id="genre-1"
                        label="Operator"
                        onChange={(event) => handleStringField("genres", event.target.value, "operator", "string")}
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
                        isOptionEqualToValue={(option, value) => option.name === value.name}
                        options={options}
                        getOptionLabel={(option) => `${option.name} - ${option.type_code}`}
                        loading={open}
                        onChange={(_, value) =>
                            handleStringField(
                                "genres",
                                value.map((v) => `${v.name}-${v.type_code}`),
                                "def",
                                "string"
                            )
                        }
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
                                            {open ? <CircularProgress color="inherit" size={20} /> : null}
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
