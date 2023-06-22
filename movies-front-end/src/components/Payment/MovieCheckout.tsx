import { useEffect, useState } from "react";
import { MovieType, PaymentType } from "@/types/movies";
import useSWRMutation from "swr/mutation";
import { get, post } from "@/libs/api";
import { useHasUsername } from "@/hooks/auth/useHasUsername";
import { Stripe } from "@stripe/stripe-js/types/stripe-js";
import { CardMedia, Grid, Paper, Stack, Typography } from "@mui/material";
import format from "date-fns/format";
import CheckoutBox from "@/components/Payment/CheckoutBox";
import { loadStripe } from "@stripe/stripe-js";

export interface MovieCheckoutProps {
  refId: number;
  type: string;
}

export default function MovieCheckout({ refId, type }: MovieCheckoutProps) {
  const username = useHasUsername();

  const [stripePromise, setStripePromise] = useState<Promise<Stripe | null>>();
  const [clientSecret, setClientSecret] = useState("");

  const [wasPaid, setWasPaid] = useState(false);
  const [movie, setMovie] = useState<MovieType>();
  const [paymentId, setPaymentId] = useState();

  const { trigger: getMovie } = useSWRMutation(`/api/v1/movies/${refId}`, get);
  const { trigger: checkBuy } = useSWRMutation(`/api/v1/payments/check?type=${type}&refId=${refId}`, get);
  const { trigger: getPaymentIntent } = useSWRMutation("/api/v1/payments/intents", post);

  useEffect(() => {
    setStripePromise(loadStripe(`${process.env.NEXT_PUBLIC_STRIPE_PUBLISHABLE_KEY}`));
  }, []);

  useEffect(() => {
    if (refId) {
      if (type === "MOVIE") {
        getMovie().then((result: MovieType) => {
          setMovie(result);
        });
      }
    }
  }, [refId, type]);

  useEffect(() => {
    if (movie && movie.price) {
      checkBuy().then((result: PaymentType) => {
        if (result && result.ref_id === movie.id! && result.status === "succeeded") {
          setWasPaid(true);
        }
      });
    }
  }, [movie]);

  useEffect(() => {
    if (movie && !wasPaid) {
      getPaymentIntent({
        currency: "USD",
        amount: movie.price,
        automatic_payment_methods: { enabled: true },
      }).then((result) => {
        if (result) {
          setClientSecret(result.clientSecret);
          setPaymentId(result.paymentId);
        }
      });
    }
  }, [movie, wasPaid]);

  return (
    <Grid container spacing={2} sx={{ m: 2, p: 1 }}>
      {movie && type === "MOVIE" && (
        <>
          <Grid item xs={6}>
            <Paper sx={{ m: 2, p: 2 }}>
              <Stack spacing={3}>
                <Grid item xs sx={{ m: 1 }}>
                  <Typography variant="subtitle1">
                    <b>{movie.title}</b>
                  </Typography>
                  <Typography variant="body2">{format(new Date(movie.release_date!), "MMMM do, yyyy")}</Typography>
                </Grid>
                <Grid item container spacing={1}>
                  <Grid item xs={6}>
                    <CardMedia sx={{ borderRadius: "16px" }} component="img" src={movie.image_url} />
                  </Grid>
                  <Grid item xs={6} sx={{ display: "flex", justifyContent: "center" }}>
                    <Typography color="error" variant="h5">
                      Price: {movie.price} USD
                    </Typography>
                  </Grid>
                </Grid>
              </Stack>
            </Paper>
          </Grid>
          <Grid item xs={6}>
            <CheckoutBox
              paymentId={paymentId!}
              type={type}
              movieId={Number(refId)}
              price={movie.price}
              wasPaid={wasPaid}
              clientSecret={clientSecret}
              stripePromise={stripePromise!}
            />
          </Grid>
        </>
      )}
    </Grid>
  );
}
