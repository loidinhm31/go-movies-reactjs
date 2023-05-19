import {Line} from "react-chartjs-2";
import {
    CategoryScale,
    Chart as ChartJS,
    Legend,
    LinearScale,
    LineElement,
    PointElement,
    Title,
    Tooltip,
} from "chart.js";
import useSWRMutation from "swr/mutation";
import {post} from "src/libs/api";
import {GenreType} from "src/types/movies";
import {Box, CircularProgress, MenuItem, TextField} from "@mui/material";
import {useEffect, useState} from "react";
import {Analysis, AnalysisRequest, Data, Result} from "src/types/dashboard";
import NotifySnackbar, {NotifyState} from "src/components/shared/snackbar";
import Skeleton from "@mui/material/Skeleton";
import {format} from "date-fns";

export default function LineChart() {
    ChartJS.register(
        CategoryScale,
        LinearScale,
        PointElement,
        LineElement,
        Title,
        Tooltip,
        Legend
    );

    const [genreOptions, setGenreOptions] = useState<readonly GenreType[]>([]);

    const [selectedGenre, setSelectedGenre] = useState<string>("");

    const [notifyState, setNotifyState] = useState<NotifyState>({open: false, vertical: "top", horizontal: "right"});

    const [dataChart, setDataChart] = useState<any>(null);

    const options = {
        responsive: true,
        plugins: {
            legend: {
                position: "top" as const,
            },
            title: {
                display: true,
                text: "# Views and Cumulative Movies By Genre",
            },
        },
    };

    const [isLoading, setIsLoading] = useState<boolean>(false);

    const {trigger: triggerGenres} = useSWRMutation<GenreType[]>(`../../api/v1/genres`, post);
    const {trigger: triggerViews, error: viewErr} = useSWRMutation(`../../api/v1/admin/dashboard/views/genres`, post);
    const {
        trigger: triggerMovies,
        error: movieErr
    } = useSWRMutation(`../../api/v1/admin/dashboard/movies/genres/release-date`, post);

    useEffect(() => {
        if (genreOptions.length == 0) {
            triggerGenres()
                .then((data: GenreType[]) => {
                    setGenreOptions(data);
                    if (data.length > 0) {
                        setSelectedGenre(data[0].name);
                    }
                })
                .catch((error) => console.log(error));
        } else {
            setUpChart(selectedGenre);
        }
    }, [selectedGenre]);

    const handleSelectedGenre = (value: string) => {
        setSelectedGenre(value);

        setIsLoading(true);

        setUpChart(value);

        setIsLoading(false);

    }

    const setUpChart = (genre: string) => {

        const labels: string[] = [];
        const moviesData: number[] = [];
        const viewersData: number[] = [];
        const cumulativeViewersData: number[] = [];

        const timeArr: string[] = [];
        const timeMap: Map<string, string[]> = new Map();

        let currMoment = new Date().getTime();
        let t1 = format(currMoment, "yyyy-M");
        let t2 = format(currMoment,"yyyy-MMMM");
        timeArr.push(t1);
        labels.push(t2);

        for (let i = 0; i < 11; i++) {
            currMoment = currMoment - 30*24*60*60*1000 // subtract 1 month
            t1 = format(currMoment, "yyyy-M");
            t2 = format(currMoment,"MMM-yyyy");
            timeArr.push(t1);
            labels.push(t2);
        }

        labels.reverse();

        timeArr.reverse();
        timeArr.forEach((t) => {
            const splitTime = t.split("-");
            if (!timeMap.get(splitTime[0])) {
                timeMap.set(splitTime[0], []);
            }
            timeMap.get(splitTime[0])?.push(splitTime[1])
        });

        const analysis: Analysis[] = [];
        timeMap.forEach((value, key) => {
            analysis.push({
                year: key,
                months: value,
            })
        });

        const request: AnalysisRequest = {
            analysis: analysis,
            genre: selectedGenre,
        };

        triggerMovies(request)
            .then((result: Result) => {
                timeArr.forEach((t, index) => {
                    const d = result.data?.find(a => t === (a.year + "-" + a.month));
                    if (d) {
                        moviesData.push(d.cumulative!);
                    } else {
                        if (index === 0) {
                            moviesData.push(0);
                        } else {
                            for (let i = index - 1; i >= 0; i--) {
                                if (moviesData[i] !== null) {
                                    moviesData.push(moviesData[i]);
                                    break;
                                }
                            }
                        }
                    }
                })
            })
            .catch((error) => {
                setNotifyState({
                    open: true,
                    message: error.message.message,
                    vertical: "top",
                    horizontal: "right",
                    severity: "error"
                });
            });

        request.isCumulative = true;
        triggerViews(request)
            .then((result: Result) => {
                if (result.data !== null) {
                    timeArr.forEach((t, index) => {
                        const d = result.data?.find(a => t === (a.year + "-" + a.month));
                        if (d) {
                            viewersData.push(d.count);
                            cumulativeViewersData.push(d.cumulative!);
                        } else {
                            viewersData.push(0);
                            if (index === 0) {
                                cumulativeViewersData.push(0);
                            } else {
                                for (let i = index - 1; i >= 0; i--) {
                                    if (cumulativeViewersData[i] !== null) {
                                        cumulativeViewersData.push(cumulativeViewersData[i]);
                                        break;
                                    }
                                }
                            }
                        }
                    })
                } else {
                    timeArr.forEach((t) => {
                        viewersData.push(0);
                        cumulativeViewersData.push(0);
                    })
                }
            })
            .catch((error) => {
                setNotifyState({
                    open: true,
                    message: error.message.message,
                    vertical: "top",
                    horizontal: "right",
                    severity: "error"
                });
            });

        setDataChart({
            labels: labels,
            datasets: [
                {
                    label: "Cumulative Movies in Release Date",
                    data: moviesData,
                    backgroundColor: 'rgba(53, 162, 235, 0.5)',
                },
                {
                    label: "Viewers",
                    data: viewersData,
                    backgroundColor: 'rgba(255, 99, 132, 0.5)',
                },
                {
                    label: "Cumulative Viewers",
                    data: cumulativeViewersData,
                    backgroundColor: 'rgba(105, 250, 132, 0.5)',
                },
            ],
        });
    }

    return (
        <>
            <NotifySnackbar state={notifyState} setState={setNotifyState}/>

            <Box sx={{m: 1, p: 1}}>
                <TextField
                    select
                    variant="outlined"
                    sx={{minWidth: 100}}
                    label="Select Genre"
                    value={selectedGenre}
                    onChange={(event) =>
                        handleSelectedGenre(event.target.value)}
                >
                    {genreOptions &&
                        genreOptions.map((g) => {
                            return (
                                <MenuItem key={g.id} value={g.name}>{g.name}</MenuItem>
                            );
                        })
                    }
                </TextField>

                {isLoading &&
                    <CircularProgress/>
                }

                {(viewErr || movieErr) &&
                    <Skeleton variant="rectangular" width={100} height={50}/>
                }

                {dataChart !== null && !isLoading &&
                    <Line options={options} data={dataChart}/>
                }
            </Box>
        </>

    );
}