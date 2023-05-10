import {ArcElement, Chart as ChartJS, Legend, Tooltip} from "chart.js";
import {Doughnut} from "react-chartjs-2";
import useSWR from "swr";
import {Result} from "../../types/dashboard";
import {get} from "../../libs/api";
import Skeleton from "@mui/material/Skeleton";
import {useState} from "react";
import NotifySnackbar, {NotifyState} from "../shared/snackbar";
import {CircularProgress} from "@mui/material";


export default function DoughnutChart() {
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
    ]

    const {isLoading, data: result, error} = useSWR<Result>("../../api/v1/admin/dashboard/movies/genres", get, {
        onSuccess: (result) => {
            const labels = result.data.map((d) => {
                return d.genre!;
            });

            const numberData = result.data.map((d) => {
                return d.count;
            });

            setDataChart({
                labels: labels,
                datasets: [
                    {
                        label: "# of Movies By Genre",
                        data: numberData,
                        backgroundColor: backgroundColors,
                        borderColor: backgroundColors,
                        borderWidth: 1,
                    },
                ],
            });
        },
        onError: (error) => {
            setNotifyState({
                open: true,
                message: error.message,
                vertical: "top",
                horizontal: "right",
                severity: "error"
            });
        }
    });

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