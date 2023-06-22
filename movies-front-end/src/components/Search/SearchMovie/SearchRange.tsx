import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import {
  Accordion,
  AccordionDetails,
  AccordionSummary,
  MenuItem,
  Slider,
  Stack,
  TextField,
  Typography,
} from "@mui/material";
import { useEffect, useState } from "react";

interface SearchRangeProps {
  label: string;
  field: string;
  defType: string;
  min: number;
  max: number;
  step: number;
  handleRangeField: (label: string, values: string | number[], forField: string, defType: string) => void;
}

function valueLabelFormat(value: number) {
  let scaledValue = value;

  if (scaledValue > 500) {
    return "500+";
  }
  return value;
}

export function SearchRange({ label, field, defType, min, max, step, handleRangeField }: SearchRangeProps) {
  const [values, setValues] = useState<number[]>([0, 0]);

  useEffect(() => {
    handleRangeField(field, values, "def", defType);
  }, [values]);

  const handleChange = (event: Event, newValue: number[]) => {
    setValues(newValue);
  };

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
            label="Operator"
            onChange={(event) => handleRangeField(field, event.target.value, "operator", defType)}
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

          <Slider
            value={values}
            onChange={handleChange}
            valueLabelDisplay="auto"
            min={min}
            step={step}
            max={max}
            valueLabelFormat={valueLabelFormat}
          />
        </Stack>
      </AccordionDetails>
    </Accordion>
  );
}
