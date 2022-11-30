import {useEffect, useState} from "react";
import {Link} from "react-router-dom";

const Movies = () => {
    const [movies, setMovies] = useState([]);

    useEffect(() => {
        let movies = [
            {
                id: 1,
                title: "test 1",
                release_date: "1986-03-07",
                runtime: 111,
                mpaa_rating: "R",
                description: "test desc 1",
            },
            {
                id: 2,
                title: "test 2",
                release_date: "1981-03-07",
                runtime: 112,
                mpaa_rating: "PG-13",
                description: "test desc 2",
            },
        ];

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