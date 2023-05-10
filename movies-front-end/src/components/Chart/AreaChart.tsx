import {
    CategoryScale,
    Chart as ChartJS,
    Filler,
    Legend,
    LinearScale,
    LineElement,
    PointElement,
    Title,
    Tooltip,
} from "chart.js";
import {Line} from "react-chartjs-2";
import useSWRMutation from "swr/mutation";
import {post} from "src/libs/api";
import {useEffect, useState} from "react";
import {Analysis, AnalysisRequest, Result} from "../../types/dashboard";
import NotifySnackbar, {NotifyState} from "../shared/snackbar";
import {CircularProgress} from "@mui/material";
import Skeleton from "@mui/material/Skeleton";
import {format} from "date-fns";

export default function AreaChart() {
    ChartJS.register(
        CategoryScale,
        LinearScale,
        PointElement,
        LineElement,
        Title,
        Tooltip,
        Filler,
        Legend
    );

    const [dataChart, setDataChart] = useState<any>(null);

    const [notifyState, setNotifyState] = useState<NotifyState>({open: false, vertical: "top", horizontal: "right"});

    const {trigger, error} = useSWRMutation("../../api/v1/admin/dashboard/views", post);

    const [isLoading, setIsLoading] = useState<boolean>(false);

    const options = {
        responsive: true,
        plugins: {
            legend: {
                position: "top" as const,
            },
            title: {
                display: true,
                text: "# of viewers in the last 12 months",
            },
        },
    };

    useEffect(() => {
        setIsLoading(true);

        const labels: string[] = [];
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
        console.log(timeArr)

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
        }

        trigger(request)
            .then((result: Result) => {
                const numbers: number[] = [];

                timeArr.forEach((t) => {
                    const d = result.data.find(a => t === (a.year + "-" + a.month));
                    if (d) {
                        numbers.push(d.count);
                    } else {
                        numbers.push(0);
                    }
                })

                setDataChart({
                    labels: labels,
                    datasets: [
                        {
                            fill: true,
                            label: "Viewers",
                            data: numbers,
                            borderColor: "rgb(53, 162, 235)",
                            backgroundColor: "rgba(53, 162, 235, 0.5)",
                        },
                    ],
                })
            })
            .catch((error) => {
                setNotifyState({
                    open: true,
                    message: error.message,
                    vertical: "top",
                    horizontal: "right",
                    severity: "error"
                });
            })
            .finally(() => {
                setIsLoading(false);
            })
    }, [])

    return (
        <>
            <NotifySnackbar state={notifyState} setState={setNotifyState}/>

            {isLoading &&
                <CircularProgress />
            }

            {error &&
                <Skeleton variant="rectangular" width={100} height={50}/>
            }

            {dataChart !== null &&
                <Line options={options} data={dataChart}/>
            }
        </>
    );
}