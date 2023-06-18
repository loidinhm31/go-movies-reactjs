import { useEffect, useState } from "react";

export const useMovieType = (selectedType: string) => {
    const [movieType, setMovieType] = useState(selectedType.toLowerCase() !== "both" ? selectedType : "");

    useEffect(() => {
        if (selectedType.toLowerCase() === "both") {
            setMovieType("");
        } else {
            setMovieType(selectedType);
        }
    }, [selectedType]);
    return movieType;
};
