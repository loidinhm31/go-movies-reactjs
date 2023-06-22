import Wallpaper from "@/assets/images/wallpaper.png";
import Link from "next/link";
import Image from "next/image";
import { Box, Paper, Stack, Typography } from "@mui/material";
import Divider from "@mui/material/Divider";

function Home() {
  return (
    <Stack spacing={2}>
      <Box>
        <Typography variant="h4">Find a movie to watch tonight!</Typography>
      </Box>
      <Divider />
      <Paper elevation={3}>
        <Box component="span" sx={{ display: "flex", justifyContent: "center", p: 1, m: 1 }}>
          <Link href="/movies">
            <Image style={{ borderRadius: "50px" }} src={Wallpaper} alt="movie tickets"></Image>
          </Link>
        </Box>
      </Paper>
    </Stack>
  );
}

export default Home;
