import useSWRMutation from "swr/mutation";
import { get } from "@/libs/api";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";
import { Box, Button, Typography } from "@mui/material";

export default function CheckoutAuthorize() {
  const router = useRouter();
  const { providerPaymentId, type, movieId, episodeId } = router.query;

  const [isProcessing, setIsProcessing] = useState(true);

  const [isError, setIsError] = useState(false);

  const [refId, setRefId] = useState<number>();

  const { trigger: verifyPayment } = useSWRMutation(
    `/api/v1/payments/${providerPaymentId}/verification?type=${type}&refId=${refId}`,
    get
  );

  useEffect(() => {
    if (type === "MOVIE") {
      setRefId(Number(movieId));
    } else if (type === "TV") {
      setRefId(Number(episodeId));
    }
  }, [type, movieId, episodeId]);

  useEffect(() => {
    if (refId) {
      setIsProcessing(true);
      verifyPayment()
        .then((result) => {
          if (result.message === "ok") {
            if (type === "MOVIE") {
              return router.replace(`/checkout/completion?type=${type}&movieId=${movieId}`);
            } else if (type === "TV") {
              return router.replace(`/checkout/completion?type=${type}&movieId=${movieId}&episodeId=${episodeId}`);
            }
          }
        })
        .catch((error) => {
          setIsError(true);
        });
      setIsProcessing(false);
    }
  }, [refId]);

  return (
    <Box sx={{ display: "flex", justifyContent: "center", p: 5, m: 5 }}>
      {isProcessing && (
        <Button>
          <Typography variant="button">Verifying Your Payment...</Typography>
        </Button>
      )}

      {isError && (
        <Button color="error">
          <Typography variant="button">Cannot verify your payment</Typography>
        </Button>
      )}
    </Box>
  );
}
