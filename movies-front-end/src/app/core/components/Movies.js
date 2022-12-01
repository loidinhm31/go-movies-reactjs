import {useEffect, useState} from "react";
import {Link} from "react-router-dom";
import {environment} from "../../../environments/environment";

const Movies = () => {
    const [movies, setMovies] = useState([]);

    useEffect(() => {
        const headers = new Headers();
        headers.append("Content-Type", "application/json");

        const requestOptions = {
            method: "GET",
            headers: headers,
        }

        fetch(`${environment.apiBaseUrl}/movies`, requestOptions)
            .then((response) => response.json())
            .then((data) => {
                setMovies(data);
            })
            .catch(err => {
                console.log(err);
            })

        setMovies(movies);
    }, []);

    return (
        <div>
            <h2>Movies</h2>
            <hr/>
            <table className="table table-striped table-hover">
                <thead>
                <tr>
                    <th>Movie</th>
                    <th>Release Date</th>
                    <th>Rating</th>
                </tr>
                </thead>
                <tbody>
                {movies.map((m) => (
                    <tr key={m.id}>
                        <td>
                            <Link to={`/movies/${m.id}`}>
                                {m.title}
                            </Link>
                        </td>
                        <td>{m.release_date}</td>
                        <td>{m.mpaa_rating}</td>
                    </tr>
                ))}
                </tbody>
            </table>
        </div>
    )
}

export default Movies;