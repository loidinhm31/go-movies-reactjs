import { Chart as ChartJS, BarElement, CategoryScale, Legend, LinearScale, Title, Tooltip } from "chart.js";
import { Bar } from "react-chartjs-2";
import { useMovieType } from "src/hooks/useMovieType";
import useSWRMutation from "swr/mutation";
import { Result } from "src/types/dashboard";
import { get } from "src/libs/api";
import { useEffect, useState } from "react";
import NotifySnackbar, { NotifyState } from "../shared/snackbar";
import { CircularProgress } from "@mui/material";
import Skeleton from "@mui/material/Skeleton";

export default function PaymentBarChart({ movieType }) {
    ChartJS.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend);

    const options = {
        indexAxis: "y" as const,
        elements: {
            bar: {
                borderWidth: 2,
            },
        },
        responsive: true,
        plugins: {
            legend: {
                position: "top" as const,
            },
            title: {
                display: true,
                text: "Payment  ",
            },
        },
    };

    const [notifyState, setNotifyState] = useState<NotifyState>({ open: false, vertical: "top", horizontal: "right" });

    const [dataChart, setDataChart] = useState<any>(null);

    const [isLoading, setIsLoading] = useState<boolean>(false);

    const selectedType = useMovieType(movieType);

    const { trigger: fetchData, error } = useSWRMutation<Result>(
        `/api/v1/admin/dashboard/payments?type=${selectedType}`,
        get
    );

    useEffect(() => {
        if (selectedType !== undefined) {
            setIsLoading(true);
            fetchData()
                .then((result: any) => {
                    setDataChart({
                        labels: ["Total Payment (USD) " + selectedType],
                        datasets: [
                            {
                                label: "Total Amount",
                                data: [result.total_amount],
                                borderColor: "rgb(255, 99, 132)",
                                backgroundColor: "rgba(255, 99, 132, 0.5)",
                            },
                            {
                                label: "Total Received",
                                data: [result.total_received],
                                borderColor: "rgb(53, 162, 235)",
                                backgroundColor: "rgba(53, 162, 235, 0.5)",
                            },
                        ],
                    });
                })
                .catch((error) => {
                    setNotifyState({
                        open: true,
                        message: error.message.message,
                        vertical: "top",
                        horizontal: "right",
                        severity: "error",
                    });
                })
                .finally(() => setIsLoading(false));
        }
    }, [selectedType]);

    return (
        <>
            <NotifySnackbar state={notifyState} setState={setNotifyState} />

            {isLoading && <CircularProgress />}

            {error && <Skeleton variant="circular" width={80} height={80} />}

            {dataChart !== null && <Bar options={options} data={dataChart} />}
        </>
    );
}
