import {useEffect, useState} from "react";
import {useRouter} from "next/router";
import {useSession} from "next-auth/react";
import useSWR from "swr";
import {GenreType, MovieType} from "../../../types/movies";
import {get, post} from "../../../libs/api";
import useSWRMutation from "swr/mutation";
import {Checkbox, FormControl, Input, MenuItem, Select, TextField} from "@mui/material";

const EditMovie = () => {
    const router = useRouter();

    const {data: session} = useSession();

    const [movie, setMovie] = useState<MovieType>({
        title: "",
        description: "",
        release_date: null,
        runtime: 0,
        mpaa_rating: "",
        genres: [],
        genres_array: [Array(13).fill(false)],  // 13 values genres can have
    });

    // Get id from the URL
    let {id} = router.query;

    const {data: genres} = useSWR<GenreType[]>(`../api/genres`, get);
    const {trigger: fetchMovie} = useSWRMutation<MovieType>(`../api/movies/${id}`, get);
    const {trigger: triggerMovie} = useSWRMutation(`../api/admin/movies/save`, post);
    const {trigger: deleteMovie} = useSWRMutation(`../api/admin/movies/delete/${id}`, get);

    const mpaaOptions = [
        {id: "G", value: "G"},
        {id: "PG", value: "PG"},
        {id: "PG13", value: "PG13"},
        {id: "R", value: "R"},
        {id: "NC17", value: "NC17"},
        {id: "18A", value: "18A"},
    ];

    useEffect(() => {
        // Check user
        if (!session) {
            router.push("/auth/signin")
        }

        if (id === undefined) {
            // Adding a movies
            setMovie({
                title: "",
                description: "",
                release_date: null,
                runtime: 0,
                mpaa_rating: "",
                genres: [],
                genres_array: [Array(13).fill(false)],
            });

            const checks = [];
            genres.forEach((g) => {
                checks.push({id: g.id, checked: false, genre: g.genre});
            });

            setMovie((m) => ({
                ...m,
                genres: checks,
                genres_array: [],
            }));
        } else {
            fetchMovie()
                .then((movie) => {
                    const checks = [];

                    movie.genres.forEach((g) => {
                        if (movie.genres_array.indexOf(g.id) !== -1) {
                            checks.push({id: g.id, checked: true, genre: g.genre});
                        } else {
                            checks.push({id: g.id, checked: false, genre: g.genre});
                        }
                    });

                    // Set state
                    setMovie({
                        ...movie,
                        genres: checks,
                    });
                })
                .catch((err) => {
                    console.log(err);
                });
        }
    }, [id, router]);

    const handleSubmit = (event) => {
        event.preventDefault();

        let errors = [];
        let required = [
            {field: movie.title, name: "title"},
            {field: movie.release_date, name: "release_date"},
            {field: movie.runtime, name: "runtime"},
            {field: movie.description, name: "description"},
            {field: movie.mpaa_rating, name: "mpaa_rating"},
        ];

        required.forEach(function (obj) {
            if (obj.field === "") {
                errors.push(obj.name);
            }
        });

        // if (movie.genres_array.length === 0) {
        //     Swal.fire({
        //         title: "Error!",
        //         text: "You must choose at least one genre!",
        //         icon: "error",
        //         confirmButtonText: "OK",
        //     });
        //     errors.push("genres");
        // }

        if (errors.length > 0) {
            return false;
        }

        triggerMovie(movie).then((data) => {
            if (data.error) {
                console.log(data.error);
            } else {
                router.push("/manage-catalogue");
            }
        }).catch((err) => {
            console.log(err);
        });
    };

    const handleChange = (name: string) => (event) => {
        let value = event.target.value;
        setMovie({
            ...movie,
            [name]: value,
        });
    };

    const handleCheck = (event, position) => {
        let tmpArr = movie.genres;
        tmpArr[position].checked = !tmpArr[position].checked;

        let tmpIDs = movie.genres_array;
        if (!event.target.checked) {
            tmpIDs.splice(tmpIDs.indexOf(event.target.value));
        } else {
            tmpIDs.push(parseInt(event.target.value, 10));
        }

        setMovie({
            ...movie,
            genres_array: tmpIDs,
        });
    };

    const confirmDelete = () => {
        // Swal.fire({
        //     title: "Delete movies?",
        //     text: "You cannot undo this action!",
        //     icon: "warning",
        //     showCancelButton: true,
        //     confirmButtonColor: "#3085d6",
        //     cancelButtonColor: "#d33",
        //     confirmButtonText: "Yes"
        // }).then((result) => {
        //     if (result.isConfirmed) {
        //
        //         deleteMovie()
        //             .then((data) => {
        //                 if (data.error) {
        //                     console.log(data.error);
        //                 } else {
        //                     router.push("/manage-catalogue");
        //                 }
        //             })
        //             .catch(err => {
        //                 console.log(err)
        //             });
        //     }
        // })
    }

    return (
        <div>
            <h2>Add/Edit Movie</h2>
            <hr/>

            <FormControl onSubmit={handleSubmit} className="mb-3">
                <input type="hidden" name="id" value={movie.id} id="id"></input>

                <Input
                    title={"Title"}
                    className={"form-control"}
                    type={"text"}
                    name={"title"}
                    value={movie.title}
                    onChange={handleChange("title")}
                />

                <Input
                    title={"Release Date"}
                    className={"form-control"}
                    type={"date"}
                    name={"release_date"}
                    value={movie.release_date}
                    onChange={handleChange("release_date")}
                />

                <Input
                    title={"Runtime"}
                    className={"form-control"}
                    type={"text"}
                    name={"runtime"}
                    value={movie.runtime}
                    onChange={handleChange("runtime")}
                />

                <Select
                    value={movie.mpaa_rating}
                    onChange={handleChange("mpaa_rating")}
                >
                    {mpaaOptions.map((o) =>
                        <MenuItem key={o.id} value={o.value}>{o.value}</MenuItem>
                    )}
                </Select>
                <TextField
                    id="outlined-multiline-flexible"
                    label="Description"
                    multiline
                    maxRows={4}
                    value={movie.description}
                />

                <hr/>

                <h3>Genres</h3>

                {movie.genres && movie.genres.length > 1 && (
                    <>
                        {Array.from(movie.genres).map((g, index) => (
                            <Checkbox
                                title={g.genre}
                                name={"genre"}
                                key={index}
                                id={"genre-" + index}
                                onChange={(event) => handleCheck(event, index)}
                                value={g.id}
                                checked={movie.genres[index].checked}
                            />
                        ))}
                    </>
                )}

                <hr/>

                <button className="btn btn-primary">Save</button>

                {movie.id > 0 && (
                    <a href="src/app/core/components#!" className="btn btn-danger ms-2" onClick={confirmDelete}>
                        Delete Movie
                    </a>
                )}
            </FormControl>
        </div>
    );

};

export default EditMovie;
