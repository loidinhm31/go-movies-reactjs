import React, { useEffect, useState } from "react";
import Link from "next/link";
import {useRouter} from "next/router";
import {MovieType} from "../../types/movies";
import useSWR from "swr";
import {get} from "../../libs/api";

const ManageCatalogue = () => {
    const router = useRouter();
    const [movies, setMovies] = useState<MovieType[]>([]);
    // const { jwtToken } = useJwtToken();

    const { data } = useSWR(`../api/admin/movies`, get, {
        onSuccess: (data) => {
            setMovies(data);
        }
    });

    useEffect( () => {
        // if (jwtToken === "") {
        //     router.push("/login");
        //     return
        // }

    }, [router]);

    return(
        <div>
            <h2>Manage Catalogue</h2>
            <hr />
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
        </div>
    )
}

export default ManageCatalogue;