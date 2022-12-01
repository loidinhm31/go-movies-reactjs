import {useEffect, useState} from "react";
import {useParams} from "react-router-dom";

const Movie = () => {
    const [movie, setMovie] = useState({});
    let {id} = useParams();

    useEffect(() => {
        let movie = {
            id: 1,
            title: "test 1",
            release_date: "1986-03-07",
            runtime: 111,
            mpaa_rating: "R",
            description: "test desc 1",
        };
        setMovie(movie);
    }, [id]);

    return (
        <div>
            <h2>Movie: {movie.title}</h2>
            <small><em>{movie.release_date}, {movie.runtime} minutes, Rated {movie.mpaa_rating}</em></small>
            <hr/>
            <p>{movie.description}</p>
        </div>
    )
}

export default Movie;