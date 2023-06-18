import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import { Accordion, AccordionDetails, AccordionSummary, MenuItem, Stack, TextField, Typography } from "@mui/material";
import { useEffect, useState } from "react";

interface SearchDateProps {
    label: string;
    field: string;
    defType: string;
    handleDateField: (label: string, value: string, forField: string, defType: string, dateType: string) => void;
}

export function SearchDate({ label, field, defType, handleDateField }: SearchDateProps) {
    const [startDate, setStartDate] = useState("");
    const [endDate, setEnDate] = useState("");

    useEffect(() => {
        if (startDate === "") {
            setEnDate("");
        }
        handleDateField(field, startDate, "def", defType, "from");
    }, [startDate]);

    useEffect(() => {
        if (endDate === "") {
            setStartDate("");
        }
        handleDateField(field, endDate, "def", defType, "to");
    }, [endDate]);

    return (
        <Accordion>
            <AccordionSummary expandIcon={<ExpandMoreIcon />}>
                <Typography>Release Date</Typography>
            </AccordionSummary>
            <AccordionDetails>
                <Stack spacing={2} direction="row">
                    <TextField
                        select
                        variant="filled"
                        sx={{ minWidth: 100 }}
                        id={`${field}-1`}
                        label="Operator"
                        onChange={(event) => handleDateField(field, event.target.value, "operator", defType, "")}
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
                        value={startDate}
                        onChange={(e) => setStartDate(e.target.value)}
                    />

                    <TextField
                        fullWidth
                        variant="outlined"
                        label="To"
                        type="date"
                        name="release-4"
                        value={endDate}
                        onChange={(e) => setEnDate(e.target.value)}
                    />
                </Stack>
            </AccordionDetails>
        </Accordion>
    );
}
