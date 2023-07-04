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
import { Fragment, useEffect, useState } from "react";
import { get } from "@/libs/api";
import { GenreType } from "@/types/movies";
import useSWRMutation from "swr/mutation";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import { useMovieType } from "@/hooks/useMovieType";

interface SearchGenreProps {
  movieType: string;
  handleStringField: (label: string, values: string | string[], forField: string, defType: string) => void;
}

export function SearchGenre({ movieType, handleStringField }: SearchGenreProps) {
  const [operator, setOperator] = useState("");
  const [open, setOpen] = useState(false);
  const [isLoading, setIsLoading] = useState(true);
  const [options, setOptions] = useState<readonly GenreType[]>();

  const selectedType = useMovieType(movieType);
  const [selectedGenres, setSelectedGenres] = useState<GenreType[]>([]);

  const { trigger: getGenres } = useSWRMutation<GenreType[]>(`/api/v1/genres?type=${selectedType}`, get);

  useEffect(() => {
    if (operator !== "") {
      handleStringField("genres", operator, "operator", "string");
    }
  }, [operator]);

  useEffect(() => {
    handleStringField(
      "genres",
      [],
      "def",
      "string"
    );

    setSelectedGenres([]);
  }, [selectedType]);

  useEffect(() => {
    if (!open) {
      return;
    }

    setIsLoading(true);
    getGenres()
      .then((data: GenreType[]) => {
        setOptions(data);
      })
      .catch((error) => console.log(error))
      .finally(() => setIsLoading(false));
  }, [open, selectedType]);

  const handleChangeGenre = (value) => {
    setSelectedGenres(value);

    handleStringField(
      "genres",
      value.map((v) => `${v.name}-${v.type_code}`),
      "def",
      "string"
    );
  };

  return (
    <Accordion
      expanded={open}
      onChange={() => setOpen(!open)}
    >
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
            value={operator}
            onChange={(event) => setOperator(event.target.value)}
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
              readOnly: true
            }}
          />
          {options &&
            <Autocomplete
              fullWidth
              isOptionEqualToValue={(option, value) => option.id === value.id}
              options={options}
              getOptionLabel={(option) => `${option.name} - ${option.type_code}`}
              loading={isLoading}
              value={selectedGenres}
              onChange={(_, value) => handleChangeGenre(value)}
              multiple
              id="genre-3"
              renderInput={(params) => (
                <TextField
                  {...params}
                  placeholder="Genres"
                  label="Genres"
                  InputProps={{
                    ...params.InputProps,
                    endAdornment: (
                      <Fragment>
                        {isLoading ? <CircularProgress color="inherit" size={20} /> : null}
                        {params.InputProps.endAdornment}
                      </Fragment>
                    )
                  }}
                />
              )}
            />
          }
        </Stack>
      </AccordionDetails>
    </Accordion>
  );
}
