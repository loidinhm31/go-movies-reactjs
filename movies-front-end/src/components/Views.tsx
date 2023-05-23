import useSWR from "swr";
import React, {useEffect} from "react";
import {Chip} from "@mui/material";
import MovieIcon from "@mui/icons-material/Movie";
import {get} from "src/libs/api";
interface View {
    message: string;
    views: number;
}

interface ViewProps {
    movieId: number;
    wasMutateView: boolean;
    setWasMuateView: (flag: boolean) => void;
}

export function Views({movieId, wasMutateView, setWasMuateView}: ViewProps) {
    const {data, mutate, isValidating} = useSWR<View>(`/api/v1/views/${movieId}`, get);

    useEffect(() => {
        if (wasMutateView) {
            mutate().then(() => {
                setWasMuateView(false);
            });
        }
    }, [wasMutateView]);

    return (
        <>
            <Chip
                icon={<MovieIcon/>}
                label={`${data?.views ? data.views : 0} views`}
                color="info"
            />
        </>
    );
}