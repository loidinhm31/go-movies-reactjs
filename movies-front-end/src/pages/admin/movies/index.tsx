import {useEffect, useState} from "react";
import {useRouter} from "next/router";
import {useSession} from "next-auth/react";
import useSWR from "swr";
import {GenreType, MovieType} from "../../../types/movies";
import {get, post} from "../../../libs/api";
import useSWRMutation from "swr/mutation";
import {
    Box,
    Button,
    Checkbox,
    Divider,
    FormControlLabel,
    FormGroup,
    Grid,
    InputAdornment,
    MenuItem,
    Stack,
    TextField,
    Typography
} from "@mui/material";
import AlertDialog from "../../../components/shared/alert";
import moment from "moment";

const EditMovie = () => {
    const router = useRouter();

    const [isOpenAlertDialog, setIsOpenAlertDialog] = useState<boolean>(false);
    const [isOpenDeleteDialog, setIsOpenDeleteDialog] = useState<boolean>(false);
    const [isConfirmDelete, setIsConfirmDelete] = useState<boolean>(false);

    const {data: session} = useSession();

    const [movie, setMovie] = useState<MovieType>({
        title: "",
        description: "",
        release_date: null,
        runtime: 0,
        mpaa_rating: "",
        genres: [],
    });

    // Get id from the URL
    let {id} = router.query;

    const {data: genres, isLoading} = useSWR<GenreType[]>(`../api/v1/genres`, get);
    const {trigger: fetchMovie} = useSWRMutation<MovieType>(`../api/v1/movies/${id}`, get);
    const {trigger: triggerMovie} = useSWRMutation(`../api/v1/admin/movies/save`, post);
    const {trigger: deleteMovie} = useSWRMutation(`../api/v1/admin/movies/delete/${id}`, get);

    const mpaaOptions = [
        {id: "G", value: "G"},
        {id: "PG", value: "PG"},
        {id: "PG13", value: "PG-13"},
        {id: "R", value: "R"},
        {id: "NC17", value: "NC-17"},
        {id: "18A", value: "18A"},
    ];

    useEffect(() => {
        // Check user
        // if (!session) {
        //     router.push("/auth/signin")
        // }
        if (!isLoading) {
            if (id === undefined) {

                // Adding a movies
                setMovie({
                    title: "",
                    description: "",
                    release_date: null,
                    runtime: 0,
                    mpaa_rating: "",
                    genres: [],
                });

                const checks = [];
                genres.forEach((g) => {
                    checks.push({id: g.id, checked: false, genre: g.genre});
                });

                setMovie((m) => ({
                    ...m,
                    genres: checks,
                }));

            } else {
                fetchMovie()
                    .then((movie) => {
                        const checks = [];

                        genres.forEach((g) => {
                            if (movie.genres.some(mg => mg.id === g.id)) {
                                checks.push({id: g.id, genre: g.genre, checked: true});
                            } else {
                                checks.push({id: g.id, genre: g.genre, checked: false});
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
        }

    }, [id, router, genres]);

    useEffect(() => {
        if (isConfirmDelete) {
            deleteMovie()
                .then((data) => {
                    if (data.error) {
                        console.log(data.error);
                    } else {
                        router.push("/manage-catalogue");
                    }
                })
                .catch(err => {
                    console.log(err)
                });
        }
    }, [isConfirmDelete])

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

        // Check genres
        if (!movie.genres.some(g => g.checked)) {
            setIsOpenAlertDialog(true);
            errors.push("genres");
        }

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

    const handleChange = (event, name: string) => {
        let value: string | number = event.target.value;
        if (name === "runtime") {
            value = Number(value);
        }
        setMovie({
            ...movie,
            [name]: value,
        });
    };

    const handleCheck = (event, position: number) => {
        let tmpArr = movie.genres;

        tmpArr[position].checked = event.target.checked;

        setMovie({
            ...movie,
            genres: tmpArr,
        });
    };

    const confirmDelete = (event) => {
        event.preventDefault();
        setIsOpenDeleteDialog(true);
    }

    return (
        <>
            {isOpenAlertDialog &&
                <AlertDialog
                    open={isOpenAlertDialog}
                    setOpen={setIsOpenAlertDialog}
                    title={"Error!"}
                    description={"You must choose at least one genre!"}
                    confirmText={"Agree"}/>
            }
            {isOpenDeleteDialog &&
                <AlertDialog
                    open={isOpenDeleteDialog}
                    setOpen={setIsOpenDeleteDialog}
                    title={"Delete Item"}
                    description={"You cannot undo this action!"}
                    confirmText={"Yes"}
                    showCancelButton={true}
                    setConfirmDelete={setIsConfirmDelete}
                />
            }
            <Stack spacing={2}>

                <Box sx={{p: 1, m: 1}}>
                    <Typography variant="h4">Add/Edit Movie</Typography>
                </Box>
                <Divider/>
                <Box sx={{display: "flex", justifyContent: "center", p: 1, m: 1, width: 1}}>
                    <form onSubmit={handleSubmit}>
                        <Grid container spacing={2}>
                            <input type="hidden" name="id" defaultValue={movie.id} id="id" readOnly={true}></input>
                            <Grid item xs={12}>
                                <TextField
                                    sx={{width: 1}}
                                    label="Title"
                                    variant="outlined"
                                    value={movie.title}
                                    onChange={e => handleChange(e, "title")}
                                />
                            </Grid>

                            <Grid item xs={4}>
                                <TextField
                                    fullWidth
                                    focused
                                    variant="outlined"
                                    label="Release Date"
                                    type="date"
                                    name="release_date"
                                    value={moment(movie.release_date, "yyyy-MM-DD").format("yyyy-MM-DD")}
                                    onChange={e => handleChange(e, "release_date")}
                                />
                            </Grid>

                            <Grid item xs={4}>
                                <TextField
                                    fullWidth
                                    label="Runtime"
                                    variant="outlined"
                                    type="number"
                                    name="runtime"
                                    InputProps={{
                                        endAdornment: <InputAdornment position="end">minutes</InputAdornment>,
                                    }}
                                    value={movie.runtime}
                                    onChange={e => handleChange(e, "runtime")}
                                />
                            </Grid>

                            <Grid item xs={4}>
                                <TextField
                                    fullWidth
                                    select
                                    label="MPAA Rating"
                                    variant="outlined"
                                    value={movie.mpaa_rating}
                                    onChange={e => handleChange(e, "mpaa_rating")}
                                >
                                    {mpaaOptions.map((o) =>
                                        <MenuItem key={o.id} value={o.value}>{o.value}</MenuItem>
                                    )}
                                </TextField>
                            </Grid>

                            <Grid item xs={12}>
                                <TextField
                                    fullWidth
                                    label="Image Path"
                                    variant="outlined"
                                    value={movie.image}
                                    onChange={e => handleChange(e, "image")}
                                />
                            </Grid>

                            <Grid item xs={12}>
                                <TextField
                                    fullWidth
                                    label="Description"
                                    variant="outlined"
                                    multiline
                                    rows={4}
                                    value={movie.description}
                                    onChange={e => handleChange(e, "description")}
                                />
                            </Grid>

                            <Divider component="div" variant="middle"/>
                        </Grid>

                        <Typography sx={{p: 2}} variant="h6">Genres</Typography>
                        <Grid item xs={12}>
                            <FormGroup>
                                <Grid container spacing={1}>
                                    {movie.genres && movie.genres.length > 1 && (
                                        <>
                                            {Array.from(movie.genres).map((g, index) => (
                                                <Grid key={g.id} item xs={2} sx={{m: 1}}>
                                                    <FormControlLabel
                                                        label={g.genre}
                                                        name="genre"
                                                        key={index}
                                                        id={"genre-" + index}
                                                        onChange={(event) => handleCheck(event, index)}
                                                        value={g.id}
                                                        control={<Checkbox checked={movie.genres[index].checked}/>}
                                                    />
                                                </Grid>
                                            ))}
                                        </>
                                    )}
                                </Grid>

                            </FormGroup>
                        </Grid>

                        <Divider component="div" variant="middle"/>

                        <Box sx={{display: "flex", justifyContent: "center", m: 2}}>
                            <Stack direction="row" spacing={2}>
                                <Button variant="contained" type="submit">Save</Button>
                                {movie.id > 0 && (
                                    <Button variant="contained" color="error" href="src/app/core/components#!"
                                            onClick={confirmDelete}>
                                        Delete Movie
                                    </Button>
                                )}
                            </Stack>
                        </Box>
                    </form>
                </Box>
            </Stack>
        </>
    );
};

export default EditMovie;
