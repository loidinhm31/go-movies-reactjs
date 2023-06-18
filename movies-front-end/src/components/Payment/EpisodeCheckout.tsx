import { useEffect, useState } from "react";
import { MovieType, PaymentType } from "src/types/movies";
import { EpisodeType } from "src/types/seasons";
import useSWRMutation from "swr/mutation";
import { get, post } from "src/libs/api";
import { useHasUsername } from "src/hooks/auth/useHasUsername";
import { Stripe } from "@stripe/stripe-js/types/stripe-js";
import { Box, CardMedia, Grid, Paper, Stack, Typography } from "@mui/material";
import CheckoutBox from "src/components/Payment/CheckoutBox";
import { loadStripe } from "@stripe/stripe-js";
import { MovieCheckoutProps } from "src/components/Payment/MovieCheckout";
import format from "date-fns/format";

export default function EpisodeCheckout({ refId, type }: MovieCheckoutProps) {
    const username = useHasUsername();

    const [stripePromise, setStripePromise] = useState<Promise<Stripe | null>>();
    const [clientSecret, setClientSecret] = useState("");

    const [wasPaid, setWasPaid] = useState(false);
    const [episode, setEpisode] = useState<EpisodeType>();
    const [movie, setMovie] = useState<MovieType>();

    const { trigger: getRootMovie } = useSWRMutation(`/api/v1/movies/checkout?episodeId=${refId}`, get);
    const { trigger: getEpisode } = useSWRMutation(`/api/v1/episodes/${refId}`, get);
    const { trigger: checkBuy } = useSWRMutation(`/api/v1/payments/check?type=${type}&refId=${refId}`, get);
    const { trigger: getPaymentIntent } = useSWRMutation("/api/v1/payments/intents", post);
    const [paymentId, setPaymentId] = useState();

    useEffect(() => {
        setStripePromise(loadStripe(`${process.env.NEXT_PUBLIC_STRIPE_PUBLISHABLE_KEY}`));
    }, []);

    useEffect(() => {
        if (refId) {
            getEpisode().then((result: EpisodeType) => {
                setEpisode(result);
            });

            getRootMovie().then((result: MovieType) => {
                setMovie(result);
            });
        }
    }, [refId, type]);

    useEffect(() => {
        if (episode && episode.price) {
            checkBuy().then((result: PaymentType) => {
                if (result && result.ref_id === episode.id! && result.status === "succeeded") {
                    setWasPaid(true);
                }
            });
        }
    }, [episode]);

    useEffect(() => {
        if (episode && !wasPaid) {
            getPaymentIntent({
                currency: "USD",
                amount: episode?.price,
                automatic_payment_methods: { enabled: true },
            }).then((result) => {
                if (result) {
                    setClientSecret(result.clientSecret);
                    setPaymentId(result.paymentId);
                }
            });
        }
    }, [episode, wasPaid]);

    return (
        <Grid container spacing={2} sx={{ m: 2, p: 1 }}>
            {movie && (
                <Grid item xs={6}>
                    <Paper sx={{ m: 2, p: 2 }}>
                        <Stack spacing={3}>
                            <Grid item xs sx={{ m: 1 }}>
                                <Typography variant="subtitle1">
                                    <b>{movie.title}</b>
                                </Typography>
                            </Grid>
                            <Grid item container spacing={1}>
                                <Grid item xs={6}>
                                    <CardMedia sx={{ borderRadius: "16px" }} component="img" src={movie.image_url} />
                                </Grid>
                                <Grid item xs={6} sx={{ display: "flex", justifyContent: "center" }}>
                                    <Stack spacing={3}>
                                        <Stack spacing={2} direction="row">
                                            <Box sx={{ display: "flex", justifyItems: "center" }}>
                                                <Typography variant="subtitle1">
                                                    <b>
                                                        {episode?.season?.name} - {episode?.name}
                                                    </b>
                                                </Typography>
                                            </Box>
                                            <Box sx={{ display: "flex", justifyItems: "center" }}>
                                                {episode && (
                                                    <Typography variant="subtitle1">
                                                        {format(new Date(episode?.air_date!), "MMMM do, yyyy")}
                                                    </Typography>
                                                )}
                                            </Box>
                                        </Stack>
                                        <Typography color="error" variant="h5">
                                            Price: {episode?.price} USD
                                        </Typography>
                                    </Stack>
                                </Grid>
                            </Grid>
                        </Stack>
                    </Paper>
                </Grid>
            )}
            <Grid item xs={6}>
                <CheckoutBox
                    paymentId={paymentId!}
                    type={type}
                    movieId={movie?.id!}
                    episodeId={Number(refId)}
                    price={episode?.price}
                    wasPaid={wasPaid}
                    clientSecret={clientSecret}
                    stripePromise={stripePromise!}
                />
            </Grid>
        </Grid>
    );
}
