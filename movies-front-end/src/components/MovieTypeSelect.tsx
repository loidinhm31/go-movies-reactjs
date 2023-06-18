import { FormControl, FormControlLabel, FormLabel, Paper, Radio, RadioGroup } from "@mui/material";

export const movieTypes = ["MOVIE", "TV"] as const;

interface MovieTypeSelectProps {
    optionalType?: string[];
    selectedType: string;
    setSelectedType: (type: string) => void;
}

const MovieTypeSelect = ({ optionalType, selectedType, setSelectedType }: MovieTypeSelectProps) => {
    return (
        <Paper elevation={12} sx={{ p: 2 }}>
            <FormControl>
                <FormLabel>Type</FormLabel>
                <RadioGroup row value={selectedType} onChange={(event) => setSelectedType(event.target.value)}>
                    <>
                        {optionalType &&
                            optionalType!.map((t, index) => {
                                return (
                                    <FormControlLabel key={`${t}-${index}`} value={t} control={<Radio />} label={t} />
                                );
                            })}
                        {movieTypes.map((t, index) => {
                            let label;
                            if (t === "MOVIE") {
                                label = "Movie";
                            } else if (t === "TV") {
                                label = "TV Series";
                            }
                            return <FormControlLabel key={index} value={t} control={<Radio />} label={label} />;
                        })}
                    </>
                </RadioGroup>
            </FormControl>
        </Paper>
    );
};

export default MovieTypeSelect;
