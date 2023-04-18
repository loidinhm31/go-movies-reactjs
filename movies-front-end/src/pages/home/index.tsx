import Ticket from "../../assets/images/movie_tickets.jpg";
import Link from "next/link";
import Image from "next/image";
import {Box, Stack, Typography} from "@mui/material";
import Divider from "@mui/material/Divider";

function Home() {
    return (
       <Stack spacing={2}>
           <Box>
               <Typography variant="h4">Find a movie to watch tonight!</Typography>
           </Box>
           <Divider />
           <Box component="span"
                sx={{display: "flex", justifyContent: "center", p: 1, m: 1 }}>
               <Link href="/movies">
                   <Image src={Ticket} alt="movie tickets"></Image>
               </Link>
           </Box>
       </Stack>
    )
}

export default Home;