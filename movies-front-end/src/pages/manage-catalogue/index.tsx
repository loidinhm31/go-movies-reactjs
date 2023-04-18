import React, {useEffect, useState} from "react";
import Link from "next/link";
import {useRouter} from "next/router";
import {MovieType} from "../../types/movies";
import useSWR from "swr";
import {get} from "../../libs/api";
import {Box, Button, Divider, Stack, Typography} from "@mui/material";
import {useSession} from "next-auth/react";

const ManageCatalogue = () => {
    const session = useSession();
    const router = useRouter();
    const [movies, setMovies] = useState<MovieType[]>([]);

    const { data, isLoading } = useSWR(`../api/admin/movies`, get, {
        onSuccess: (data) => {
            setMovies(data);
        }
    });

    useEffect( () => {
        // if (jwtToken === "") {
        //     router.push("/auth");
        //     return
        // }

    }, [router]);

    return(
        <Stack spacing={2}>
            <Box sx={{display: "flex", p: 1, m: 1}}>
                <Typography variant="h4">Manage Catalogue</Typography>
            </Box>
            <Divider/>
            <Box>
                <Button variant="contained" href="/admin/movies" >Add Movie</Button>
            </Box>

            {!isLoading &&
                <Box component="span"
                     sx={{display: "flex", justifyContent: "center", p: 1, m: 1}}>
                    <table className="table table-striped table-hover">
                        <thead>
                        <tr>
                            <th>Movie</th>
                            <th>Release Date</th>
                            <th>Rating</th>
                        </tr>
                        </thead>
                        <tbody>
                        {movies && movies.map((m) => (
                            <tr key={m.id}>
                                <td>
                                    <Link href={`/admin/movie/${m.id}`}>
                                        {m.title}
                                    </Link>
                                </td>
                                <td>{new Date(m.release_date).toDateString()}</td>
                                <td>{m.mpaa_rating}</td>
                            </tr>
                        ))}
                        </tbody>
                    </table>
                </Box>
            }
        </Stack>
    )
}

export default ManageCatalogue;