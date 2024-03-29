import { useRouter } from "next/router";
import { Box, Button, Stack, Typography } from "@mui/material";

export default function PaymentComplete() {
  const router = useRouter();
  const { type, movieId, episodeId } = router.query;

  function handleGoTo() {
    if (type === "MOVIE") {
      router.replace(`/movies/${movieId}`);
    } else if (type === "TV") {
      router.replace(`/movies/episodes/${episodeId}?movieId=${movieId}`);
    }
  }

  return (
    <Box sx={{ display: "flex", justifyContent: "center", m: 2, p: 3 }}>
      <Stack sx={{ display: "flex", justifyContent: "center" }}>
        <Box sx={{ m: 1, p: 1 }}>
          <Typography variant="subtitle1">Payment Completed, Enjoy your movie</Typography>
        </Box>

        <Box sx={{ display: "flex", justifyContent: "center" }}>
          <Button variant="contained" onClick={handleGoTo}>
            Go to movie
          </Button>
        </Box>
      </Stack>
    </Box>
  );
}
