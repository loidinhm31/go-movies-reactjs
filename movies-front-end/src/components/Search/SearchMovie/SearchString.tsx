import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import {
  Accordion,
  AccordionDetails,
  AccordionSummary,
  Autocomplete,
  Chip,
  MenuItem,
  Stack,
  TextField,
  Typography,
} from "@mui/material";

interface SearchStringProps {
  label: string;
  field: string;
  defType: string;
  handleStringField: (label: string, values: string | string[], forField: string, defType: string) => void;
}

export function SearchString({ label, field, defType, handleStringField }: SearchStringProps) {
  return (
    <Accordion>
      <AccordionSummary expandIcon={<ExpandMoreIcon />}>
        <Typography>{label}</Typography>
      </AccordionSummary>
      <AccordionDetails>
        <Stack spacing={2} direction="row">
          <TextField
            select
            variant="filled"
            sx={{ minWidth: 100 }}
            id={`${field}-1`}
            label={"Operator"}
            onChange={(event) => handleStringField(field, event.target.value, "operator", defType)}
          >
            <MenuItem value={"and"}>AND</MenuItem>
            <MenuItem value={"or"}>OR</MenuItem>
          </TextField>
          <TextField
            fullWidth
            variant="filled"
            id={`${field}-2`}
            label="Field"
            defaultValue={label}
            InputProps={{
              readOnly: true,
            }}
          />
          <Autocomplete
            fullWidth
            onChange={(_, value) => handleStringField(field, value, "def", defType)}
            multiple
            id={`${field}-3`}
            options={[]}
            freeSolo
            renderTags={(vals: string[], getTagProps) =>
              vals.map((option: string, index: number) => <Chip label={option} {...getTagProps({ index })} />)
            }
            renderInput={(params) => (
              <TextField
                {...params}
                {...(defType === "number" ? { type: "number" } : {})}
                label="Value"
                placeholder="Value"
              />
            )}
          />
        </Stack>
      </AccordionDetails>
    </Accordion>
  );
}
