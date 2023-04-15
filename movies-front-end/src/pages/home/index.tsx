import Ticket from "../../assets/images/movie_tickets.jpg";
import Link from "next/link";
import Image from "next/image";
import {Box, Stack} from "@mui/material";

function Home() {
    return (
        <Stack>
            <Box sx={{display: "flex", p: 1, m: 1}}>
                <h2>Find a movie to watch tonight!</h2>
            </Box>
            <hr/>
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