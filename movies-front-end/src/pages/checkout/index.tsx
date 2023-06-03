import {Elements} from "@stripe/react-stripe-js";
import {loadStripe} from "@stripe/stripe-js";
import {useEffect, useState} from "react";
import {get, post} from "src/libs/api";
import {Stripe} from "@stripe/stripe-js/types/stripe-js";
import CheckoutForm from "src/components/Payment/CheckoutForm";
import useSWRMutation from "swr/mutation";
import {useRouter} from "next/router";
import {MovieType} from "src/types/movies";
import {Box, Button, CardMedia, Grid, Paper, Stack, Typography} from "@mui/material";
import format from "date-fns/format";
import {useHasUsername} from "src/hooks/auth/useHasUsername";

export default function Payment() {
    const router = useRouter();
    const username = useHasUsername();
    const [stripePromise, setStripePromise] = useState<Promise<Stripe | null>>();
    const [clientSecret, setClientSecret] = useState("");
    const [paymentId, setPaymentId] = useState();
    const {movieId} = router.query;

    const [wasPaid, setWasPaid] = useState(false);
    const [movie, setMovie] = useState<MovieType>();

    const {trigger: getPaymentIntent} = useSWRMutation("/api/v1/payments", post);
    const {trigger: getMovie} = useSWRMutation(`/api/v1/movies/${movieId}`, get);
    const {trigger: checkBuy} = useSWRMutation(`/api/v1/collections/check?username=${username}&movieId=${movie?.id!}`, get);

    useEffect(() => {
        setStripePromise(loadStripe(`${process.env.NEXT_PUBLIC_STRIPE_PUBLISHABLE_KEY}`));
    }, []);

    useEffect(() => {
        if (movieId) {
            getMovie().then((result) => {
                setMovie(result);
            });
        }
    }, [movieId]);


    useEffect(() => {
        if (movie && movie.price) {
            checkBuy().then((result) => {
                if (result.movie_id === movie.id! && result.username === username) {
                    setWasPaid(true);
                }
            });
        }
    }, [movie]);

    useEffect(() => {
        if (movie && !wasPaid) {
            getPaymentIntent({
                currency: "USD",
                amount: movie?.price,
                automatic_payment_methods: {enabled: true},
            }).then((result) => {
                setClientSecret(result.clientSecret);
                setPaymentId(result.paymentId);
            })
        }
    }, [movie, wasPaid]);


    return (
        <Grid container spacing={2} sx={{m: 2, p: 1}}>
            {movie &&
                <>
                    <Grid item xs={6}>
                        <Paper
                            sx={{m: 2, p: 2,}}
                        >
                            <Stack spacing={3}>
                                <Grid item xs>
                                    <Typography gutterBottom variant="subtitle1" component="div" sx={{m: 1}}>
                                        <b>{movie.title}</b>
                                    </Typography>
                                </Grid>
                                <Grid item container spacing={1}>
                                    <Grid item xs={6}>
                                        <CardMedia
                                            sx={{borderRadius: "16px"}}
                                            component="img"
                                            src={movie.image_url}
                                        />
                                    </Grid>
                                    <Grid item xs={6}>
                                        <Typography variant="body2" gutterBottom>
                                            {format(new Date(movie.release_date!), "MMMM do, yyyy")}
                                        </Typography>
                                        <Typography gutterBottom variant="inherit" component="div">
                                            {movie.description}
                                        </Typography>

                                    </Grid>
                                </Grid>
                            </Stack>
                        </Paper>
                    </Grid>
                    <Grid item xs={6}>

                        {!wasPaid && movie.price && clientSecret && stripePromise && (
                            <Elements stripe={stripePromise} options={{clientSecret}}>
                                <CheckoutForm
                                    paymentId={paymentId!}
                                    movieId={movie.id!}
                                />
                            </Elements>
                        )}

                        {!wasPaid && !movie.price &&
                            <Box sx={{display: "flex", justifyContent: "center", p: 3}}>
                                <Button variant="contained" color="warning" sx={{p: 5}}>
                                    Cannot Pay for this movie
                                </Button>
                            </Box>
                        }

                        {wasPaid &&
                            <Box sx={{display: "flex", justifyContent: "center", p: 3}}>
                                <Button variant="contained" color="warning" sx={{p: 5}}>
                                    You bought this movie
                                </Button>
                            </Box>
                        }
                    </Grid>

                </>
            }


        </Grid>
    );
}

