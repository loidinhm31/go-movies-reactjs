import {useState} from "react";
import {PaymentElement, useElements, useStripe} from "@stripe/react-stripe-js";
import {Box, Button, Container, Typography} from "@mui/material";

interface CheckoutFormProps {
    paymentId: string;
    movieId: number;
}

export default function CheckoutForm({paymentId, movieId}: CheckoutFormProps) {
    const stripe = useStripe();
    const elements = useElements();

    const [message, setMessage] = useState<string | null>(null);
    const [isProcessing, setIsProcessing] = useState(false);


    const handleSubmit = async (e) => {
        e.preventDefault();

        if (!stripe || !elements) {
            // Stripe.js has not yet loaded.
            // Make sure to disable form submission until Stripe.js has loaded.
            return;
        }

        setIsProcessing(true);

        const {error} = await stripe.confirmPayment({
            elements,
            confirmParams: {
                // Make sure to change this to your payment completion page
                return_url: `${window.location.origin}/checkout/verify?providerPaymentId=${paymentId}&movieId=${movieId}`,

            },
        });

        if (error.type === "card_error" || error.type === "validation_error") {
            setMessage(error.message!);
        } else {

        }

        setIsProcessing(false);
    };

    return (
        <Container>
            <form id="payment-form" onSubmit={handleSubmit}>
                <PaymentElement id="payment-element"/>
                <Box sx={{display: "flex", justifyContent: "center", m: 2}}>
                    <Button
                        variant="contained"
                        disabled={isProcessing || !stripe || !elements}
                        type="submit"
                    >
                        {isProcessing ? "Processing ... " : "Pay now"}
                    </Button>
                </Box>
                {/* Show any error or success messages */}
                {message && <Typography color="error" id="payment-message">{message}</Typography>}
            </form>
        </Container>
    );
}
