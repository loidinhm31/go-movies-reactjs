import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import {
    Accordion,
    AccordionDetails,
    AccordionSummary,
    Autocomplete,
    Box,
    Button,
    MenuItem,
    Stack,
    TextField,
    Typography
} from "@mui/material";
import {FieldData} from "src/types/search";
import {SearchDate} from "./SearchDate";
import {SearchGenre} from "./SearchGenre";
import {SearchString} from "./SearchString";
import useSWR from "swr";
import {RatingType} from "../../../types/movies";
import {get} from "../../../libs/api";

interface SearchFieldProps {
    trigger: (fieldData: FieldData[]) => void;
    fieldDataMap: Map<string, FieldData>
    setFieldDataMap: (value: Map<string, FieldData>) => void;
}

export function SearchField({trigger, fieldDataMap: fieldDataMap, setFieldDataMap: setFieldData}: SearchFieldProps) {
    const {data: mpaaOptions} = useSWR<RatingType[]>("../api/v1/ratings", get);

    const handleStringField = (label: string, values: string | string[], forField: string, defType: string) => {
        let data = fieldDataMap.get(label) as FieldData;
        if (!data) {
            data = {field: label};
        }

        if (forField === "operator") {
            data.operator = values as string;
        } else if (forField === "def") {
            data.def = {
                type: defType,
                values: values as string[]
            };
        }

        fieldDataMap.set(label, data);
        setFieldData(new Map(fieldDataMap));
    }

    const handleDateField = (label: string, value: string, forField: string, defType: string, dateType: string) => {
        let data = fieldDataMap.get(label) as FieldData;
        if (!data) {
            data = {field: label};
        }

        if (forField === "operator") {
            data.operator = value as string;
        } else if (forField === "def") {
            if (!data.def?.values) {
                data.def = {
                    type: defType,
                    values: [],
                };
            }

            if (dateType === "from") {
                data.def.values[0] = value;
            } else if (dateType === "to") {
                data.def.values[1] = value;
            }
        }

        fieldDataMap.set(label, data);
        setFieldData(new Map(fieldDataMap));
    }

    const handleSearch = () => {
        let filters: FieldData[] = [];
        fieldDataMap.forEach((value: FieldData, key) => {
            if (value.field &&
                value.operator &&
                value.def && value.def.type && value.def.values.length > 0) {
                filters.push(value);
            }
        });
        trigger(filters);
    }

    return (
        <Stack sx={{width: 1}} spacing={2}>
            <SearchString
                label="Title"
                field="title"
                defType="string"
                handleStringField={handleStringField}
            />

            <SearchString
                label="Description"
                field="description"
                defType="string"
                handleStringField={handleStringField}
            />

            <SearchString
                label="Runtime"
                field="runtime"
                defType="number"
                handleStringField={handleStringField}
            />

            <SearchDate
                label="Release Date"
                field="release_date"
                defType="date"
                handleDateField={handleDateField}
            />

            <Accordion>
                <AccordionSummary
                    expandIcon={<ExpandMoreIcon/>}
                >
                    <Typography>MPAA Rating</Typography>
                </AccordionSummary>
                <AccordionDetails>
                    <Stack spacing={2} direction="row">
                        <TextField
                            select
                            variant="filled"
                            sx={{minWidth: 100}}
                            id="rating-1"
                            label="Operator"
                            onChange={(event) =>
                                handleStringField("mpaa_rating", event.target.value, "operator", "string")}
                        >

                            <MenuItem value={"and"}>AND</MenuItem>
                            <MenuItem value={"or"}>OR</MenuItem>
                        </TextField>

                        <TextField
                            fullWidth
                            variant="filled"
                            id="rating-2"
                            label="Field"
                            defaultValue="MPAA Rating"
                            InputProps={{
                                readOnly: true,
                            }}
                        />
                        {mpaaOptions &&
                            <Autocomplete
                                fullWidth
                                onChange={(_, value) => handleStringField("mpaa_rating", value.map(v => v.code), "def", "string")}
                                multiple
                                id="rating-3"
                                options={mpaaOptions}
                                getOptionLabel={(option) => option.code}
                                renderInput={(params) => (
                                    <TextField {...params} placeholder="MPAA Rating"/>
                                )}
                            />
                        }

                    </Stack>
                </AccordionDetails>
            </Accordion>

            <SearchGenre
                handleStringField={handleStringField}
            />

            <Box sx={{display: "flex", justifyContent: "flex-end"}}>
                <Button variant="contained" sx={{width: 0.1}}
                        onClick={handleSearch}>
                    Search
                </Button>
            </Box>

        </Stack>
    );
}