import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import {Accordion, AccordionDetails, AccordionSummary, MenuItem, Stack, TextField, Typography} from "@mui/material";

interface SearchDateProps {
    label: string;
    field: string;
    defType: string;
    handleDateField: (label: string, value: string, forField: string, defType: string, dateType: string) => void;
}

export function SearchDate({label, field, defType, handleDateField}: SearchDateProps) {
    return (
        <Accordion>
            <AccordionSummary
                expandIcon={<ExpandMoreIcon/>}
            >
                <Typography>Release Date</Typography>
            </AccordionSummary>
            <AccordionDetails>
                <Stack spacing={2} direction="row">
                    <TextField
                        select
                        variant="filled"
                        sx={{minWidth: 100}}
                        id={`${field}-1`}
                        label="Operator"
                        onChange={(event) =>
                            handleDateField(field, event.target.value, "operator", defType, "")}
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

                    <TextField
                        fullWidth
                        variant="outlined"
                        label="From"
                        type="date"
                        name="release-3"
                        onChange={e => handleDateField(field, e.target.value, "def", defType, "from")}
                    />

                    <TextField
                        fullWidth
                        variant="outlined"
                        label="To"
                        type="date"
                        name="release-4"
                        onChange={e => handleDateField(field, e.target.value, "def", defType, "to")}
                    />

                </Stack>
            </AccordionDetails>
        </Accordion>
    );
}
