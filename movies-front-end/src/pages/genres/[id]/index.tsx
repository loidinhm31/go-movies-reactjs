import {useRouter} from "next/router";
import useSWR from "swr";
import {get} from "../../../libs/api";
import EnhancedTable from "../../../components/movies/table/MoviesTable";


function OneGenre() {
    const router = useRouter();

    // Get the id from the url
    let {id} = router.query;

    // Need to get the "prop" passed to this component
    const {genreName} = router.query;

    // Get list of movies
    const { data: movies } = useSWR(`../api/v1/movies/genres/${id}`, get);

    return (
        <>
            <h2>Genre: {genreName}</h2>

            <hr/>

            {movies ? (
                <EnhancedTable rows={movies} />

            ) : (
                <p>No movies in this genre (yet)!</p>
            )}
        </>
    )
}

export default OneGenre;