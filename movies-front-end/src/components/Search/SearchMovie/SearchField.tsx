import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import {
    Accordion,
    AccordionDetails,
    AccordionSummary,
    Autocomplete,
    Box,
    Button,
    Container,
    MenuItem,
    Stack,
    TextField,
    Typography
} from "@mui/material";
import {useState} from "react";
import MovieTypeSelect from "src/components/MovieTypeSelect";
import {SearchDate} from "src/components/Search/SearchMovie/SearchDate";
import {SearchGenre} from "src/components/Search/SearchMovie/SearchGenre";
import {SearchRange} from "src/components/Search/SearchMovie/SearchRange";
import {SearchString} from "src/components/Search/SearchMovie/SearchString";
import {get} from "src/libs/api";
import {RatingType} from "src/types/movies";
import {FieldData, SearchRequest} from "src/types/search";
import useSWR from "swr";

interface SearchFieldProps {
    setIsClickSearch: (flag: boolean) => void;
    setSearchRequest: (searchRequest: SearchRequest) => void;
    fieldDataMap: Map<string, FieldData>
    setFieldDataMap: (value: Map<string, FieldData>) => void;
}

export function SearchField({
                                setIsClickSearch,
                                setSearchRequest,
                                fieldDataMap: fieldDataMap,
                                setFieldDataMap: setFieldData
                            }: SearchFieldProps) {
    const optionalType = ["Both"];

    const [selectedType, setSelectedType] = useState<string>(optionalType[0]);

    const {data: mpaaOptions} = useSWR<RatingType[]>("/api/v1/ratings", get);

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

        // Reset empty date
        if (value === "") {
            data.def = {
                type: defType,
                values: [],
            };
            return;
        }

        if (forField === "operator") {
            data.operator = value as string;
        } else if (forField === "def") {
            if (!data.def?.values) {
                data.def = {
                    type: defType,
                    values: [],
                };
            } else {
                if (dateType === "from") {
                    data.def.values[0] = value;
                } else if (dateType === "to") {
                    data.def.values[1] = value;
                }
            }
        }

        fieldDataMap.set(label, data);
        setFieldData(new Map(fieldDataMap));
    }

    const handleRangeField = (label: string, values: string | number[], forField: string, defType: string) => {
        let data = fieldDataMap.get(label) as FieldData;
        if (!data) {
            data = {field: label};
        }

        if (forField === "operator") {
            data.operator = values as string;
        } else if (forField === "def") {
            console.log(values)
            if ((!data.def?.values) ||
                (values[0] === 0 && values[1] === 0)) {
                data.def = {
                    type: defType,
                    values: [],
                };
            } else {
                data.def.values = (values as number[]).map((v) => v.toString());
            }
        }

        fieldDataMap.set(label, data);
        setFieldData(new Map(fieldDataMap));
    }

    const handleSearch = () => {
        let filters: FieldData[] = [];
        filters.push({
            field: "type_code",
            operator: "and",
            def: {
                type: "string",
                values: [selectedType.toLowerCase() !== "both" ? selectedType : ""]
            }
        });

        fieldDataMap.forEach((value: FieldData, key) => {
            if (value.field &&
                value.operator &&
                value.def && value.def.type && value.def.values.length > 0) {
                filters.push(value);
            }
        });

        setSearchRequest({
            filters: filters,
        } as SearchRequest);

        setIsClickSearch(true);
    }

    return (
        <Stack sx={{width: 1}} spacing={2}>
            <Container>
                <MovieTypeSelect
                    optionalType={optionalType}
                    selectedType={selectedType}
                    setSelectedType={setSelectedType}
                />
            </Container>

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

            <SearchRange
                label="Runtime"
                field="runtime"
                defType="number"
                min={0}
                max={500}
                step={1}
                handleRangeField={handleRangeField}
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
                movieType={selectedType}
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