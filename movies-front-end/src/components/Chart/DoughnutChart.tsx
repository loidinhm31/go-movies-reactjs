import {ArcElement, Chart as ChartJS, Legend, Tooltip} from "chart.js";
import {Doughnut} from "react-chartjs-2";
import useSWR from "swr";
import {Result} from "../../types/dashboard";
import {get} from "../../libs/api";
import Skeleton from "@mui/material/Skeleton";
import {useEffect, useState} from "react";
import NotifySnackbar, {NotifyState} from "../shared/snackbar";
import {CircularProgress} from "@mui/material";
import useSWRMutation from "swr/mutation";
import {useMovieType} from "../../hooks/useMovieType";


export default function DoughnutChart({movieType}) {
    ChartJS.register(ArcElement, Tooltip, Legend);

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
                text: "# of Movies By Genre",
            },
        },
    };

    const backgroundColors = [
        "rgba(255, 99, 132, 0.2)",
        "rgba(54, 162, 235, 0.2)",
        "rgba(255, 206, 86, 0.2)",
        "rgba(75, 192, 192, 0.2)",
        "rgba(153, 102, 255, 0.2)",
        "rgba(255, 159, 64, 0.2)",
        "rgba(250, 102, 64, 0.2)",
        "rgba(100, 159, 64, 0.2)",
        "rgba(25, 102, 150, 0.2)",
        "rgba(50, 25, 64, 0.2)",
        "rgba(50, 25, 200, 0.2)",
        "rgba(250, 25, 200, 0.2)",
        "rgba(50, 250, 200, 0.2)",
        "rgba(50, 206, 50, 0.2)",
        "rgba(50, 130, 111, 0.2)",
        "rgba(80, 10, 100, 0.2)",
        "rgba(255, 122, 180, 0.2)",
    ];

    const borderColors = [
        "rgba(255, 99, 132, 1)",
        "rgba(54, 162, 235, 1)",
        "rgba(255, 206, 86, 1)",
        "rgba(75, 192, 192, 1)",
        "rgba(153, 102, 255, 1)",
        "rgba(255, 159, 64, 1)",
        "rgba(250, 102, 64, 1)",
        "rgba(100, 159, 64, 1)",
        "rgba(25, 102, 150, 1)",
        "rgba(50, 25, 64, 1)",
        "rgba(50, 25, 200, 1)",
        "rgba(250, 25, 200, 1)",
        "rgba(50, 250, 200, 1)",
        "rgba(50, 206, 50, 1)",
        "rgba(50, 130, 111, 1)",
        "rgba(80, 10, 100, 1)",
        "rgba(255, 122, 180, 1)",
    ]

    const [isLoading, setIsLoading] = useState<boolean>(false);

    const selectedType = useMovieType(movieType);

    const {trigger: fetchData, error} = useSWRMutation<Result>(`../../api/v1/admin/dashboard/movies/genres?type=${selectedType}`, get);


    useEffect(() => {
        console.log("trigger")
        console.log(selectedType)
        if (selectedType !== undefined) {
            setIsLoading(true);
            fetchData().then((result) => {
                console.log(result)
                const labels = result!.data.map((d) => {
                    return `${d.name!} - ${d.type_code!}`;
                });

                const numberData = result!.data.map((d) => {
                    return d.count;
                });

                setDataChart({
                    labels: labels,
                    datasets: [
                        {
                            label: "# of Movies By Genre",
                            data: numberData,
                            backgroundColor: backgroundColors,
                            borderColor: borderColors,
                            borderWidth: 1,
                        },
                    ],
                });
            }).catch((error) => {
                setNotifyState({
                    open: true,
                    message: error.message.message,
                    vertical: "top",
                    horizontal: "right",
                    severity: "error"
                });
            }).finally(() => setIsLoading(false));
        }
    }, [selectedType]);

    return (
        <>
            <NotifySnackbar state={notifyState} setState={setNotifyState}/>

            {isLoading &&
                <CircularProgress />
            }

            {error &&
                <Skeleton variant="circular" width={80} height={80}/>

            }
            {dataChart !== null &&
                <Doughnut options={options} data={dataChart}/>
            }
        </>
    );
}