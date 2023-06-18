import { Elements } from "@stripe/react-stripe-js";
import CheckoutForm from "./CheckoutForm";
import { Box, Button } from "@mui/material";
import { Stripe } from "@stripe/stripe-js/types/stripe-js";

interface CheckoutBoxProps {
    paymentId: string;
    type: string;
    movieId?: number;
    episodeId?: number;
    wasPaid: boolean;
    price?: number;
    clientSecret?: string;
    stripePromise?: Promise<Stripe | null>;
}
export default function CheckoutBox({
    paymentId,
    type,
    movieId,
    episodeId,
    wasPaid,
    price,
    clientSecret,
    stripePromise,
}: CheckoutBoxProps) {
    return (
        <>
            {!wasPaid && price && clientSecret && stripePromise && (
                <Elements stripe={stripePromise} options={{ clientSecret }}>
                    <CheckoutForm paymentId={paymentId} type={type} movieId={movieId} episodeId={episodeId} />
                </Elements>
            )}

            {!wasPaid && !price && (
                <Box sx={{ display: "flex", justifyContent: "center", p: 3 }}>
                    <Button variant="contained" color="warning" sx={{ p: 5 }}>
                        Cannot Pay for this movie
                    </Button>
                </Box>
            )}

            {wasPaid && (
                <Box sx={{ display: "flex", justifyContent: "center", p: 3 }}>
                    <Button variant="contained" color="warning" sx={{ p: 5 }}>
                        You bought this movie
                    </Button>
                </Box>
            )}
        </>
    );
}
