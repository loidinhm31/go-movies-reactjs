import {
  Accordion,
  AccordionDetails,
  AccordionSummary,
  Autocomplete,
  MenuItem,
  Stack,
  TextField,
  Typography
} from "@mui/material";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import { RatingType } from "@/types/movies";
import { get } from "@/libs/api";
import useSWRMutation from "swr/mutation";
import { useEffect, useState } from "react";

interface SearchOptionsProps {
  label: string;
  field: string;
  defType: string;
  handleStringField: (label: string, values: string | string[], forField: string, defType: string) => void;
}

export default function SearchOptions({ label, field, defType, handleStringField }: SearchOptionsProps) {
  const [operator, setOperator] = useState("");
  const [open, setOpen] = useState(false);
  const [mpaaOptions, setMpaaOtions] = useState<RatingType[]>();

  const { trigger: getRatings } = useSWRMutation<RatingType[]>("/api/v1/ratings", get);

  useEffect(() => {
    if (operator !== "") {
      handleStringField(field, operator, "operator", defType);
    }
  }, [operator]);

  useEffect(() => {
    if (!open) {
      return;
    }

    getRatings()
      .then((data: RatingType[]) => {
        setMpaaOtions(data);
      })
      .catch((error) => console.log(error));
  }, [open]);

  return (
    <Accordion
      expanded={open}
      onChange={() => setOpen(!open)}
    >
      <AccordionSummary expandIcon={<ExpandMoreIcon />}>
        <Typography>{label}</Typography>
      </AccordionSummary>
      <AccordionDetails>
        <Stack spacing={2} direction="row">
          <TextField
            select
            variant="filled"
            sx={{ minWidth: 100 }}
            id="rating-1"
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
            id="rating-2"
            label="Field"
            defaultValue={label}
            InputProps={{
              readOnly: true
            }}
          />
          {mpaaOptions && (
            <Autocomplete
              fullWidth
              onChange={(_, value) =>
                handleStringField(
                  field,
                  value.map((v) => v.code),
                  "def",
                  defType
                )
              }
              multiple
              id="rating-3"
              options={mpaaOptions}
              getOptionLabel={(option) => option.name}
              renderInput={(params) => <TextField {...params} placeholder={label} />}
            />
          )}
        </Stack>
      </AccordionDetails>
    </Accordion>
  );
}