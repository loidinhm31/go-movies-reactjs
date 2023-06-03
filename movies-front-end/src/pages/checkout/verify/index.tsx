import useSWRMutation from "swr/mutation";
import {get} from "src/libs/api";
import {useRouter} from "next/router";
import {useEffect, useState} from "react";
import {Box, Button, Typography} from "@mui/material";

export default function CheckoutAuthorize() {
    const router = useRouter();
    const {providerPaymentId, movieId} = router.query;

    const [isProcessing, setIsProcessing] = useState(true);

    const [isError, setIsError] = useState(false);

    const {trigger: verifyPayment} = useSWRMutation(`/api/v1/payments/${providerPaymentId}/verification?&movieId=${movieId}`, get);

    useEffect(() => {
        if (movieId) {
            setIsProcessing(true);
            verifyPayment().then((result) => {
                if (result.message === "ok") {
                    return router.replace(`/checkout/completion?movieId=${movieId}`);
                }
            }).catch((error) => {
                setIsError(true);
            });
            setIsProcessing(false);
        }
    }, [movieId])

    return (
        <Box sx={{display: "flex", justifyContent: "center", p: 5, m: 5}}>
            {isProcessing &&
                <Button>
                    <Typography variant="button">Verifying Your Payment...</Typography>
                </Button>
            }

            {isError &&
                <Button color="error">
                    <Typography variant="button">Cannot verify your payment</Typography>
                </Button>
            }
        </Box>
    );
}