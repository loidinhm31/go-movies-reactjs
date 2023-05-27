import {useEffect, useRef, useState} from "react";
import {useRouter} from "next/router";
import {useSession} from "next-auth/react";
import {GenreType, MovieType, RatingType} from "src/types/movies";
import {get, post, postForm} from "src/libs/api";
import useSWRMutation from "swr/mutation";
import {
    Box,
    Button,
    Checkbox,
    Chip,
    Divider,
    FormControl,
    FormControlLabel,
    FormGroup,
    FormLabel,
    Grid,
    IconButton,
    InputAdornment,
    Link,
    MenuItem,
    Radio,
    RadioGroup,
    Stack,
    TextField,
    Typography
} from "@mui/material";
import AlertDialog from "src/components/shared/alert";
import NotifySnackbar, {NotifyState, sleep} from "src/components/shared/snackbar";
import {format} from "date-fns";
import {RemoveCircle} from "@mui/icons-material";
import {ClientError} from "../../../../libs/api_client";
import {movieTypes} from "../../../../components/MovieTypeSelect";
import useSWR from "swr";

const ReferenceMovie = () => {
    const router = useRouter();

    const [isOpenAlertDialog, setIsOpenAlertDialog] = useState<boolean>(false);
    const [isOpenDeleteDialog, setIsOpenDeleteDialog] = useState<boolean>(false);

    const [notifyState, setNotifyState] = useState<NotifyState>({open: false, vertical: "top", horizontal: "right"});

    const {data: session, status} = useSession();

    // Get id from the URL
    let {id, type} = router.query;

    const [movie, setMovie] = useState<MovieType>({
        title: "",
        type_code: "",
        description: "",
        release_date: null,
        runtime: 0,
        mpaa_rating: "",
        genres: [],
        image_path: "",
    });

    const videoFileRef = useRef<any>(null);
    const [videoFile, setVideoFile] = useState<HTMLInputElement | null>(null);
    const [videoPath, setVideoPath] = useState<string>("");

    const [newGenres, setNewGenres] = useState<string[]>([]);

    const {trigger: fetchGenres} = useSWRMutation<GenreType[]>(`../../../api/v1/genres?type=${type}`, get);
    const {trigger: fetchMovie} = useSWRMutation<MovieType>(`../../api/v1/admin/movies/references/${id}?type=${type}`, get);
    const {trigger: triggerMovie} = useSWRMutation(`../../api/v1/admin/movies/save`, post);
    const {trigger: uploadVideo} = useSWRMutation(`../../api/v1/admin/movies/video/upload`, postForm);
    const {trigger: removeVideo} = useSWRMutation(`../../api/v1/admin/movies/video/remove`, post);
    const {trigger: addNewGenres} = useSWRMutation(`../../api/v1/admin/genres`, post);

    const {data: mpaaOptions} = useSWR<RatingType[]>("../api/v1/ratings", get);

    useEffect(() => {
        if (status === "loading") {
            return;
        }
        const role = session?.user.role;

        if (role === "admin" || role === "moderator") {
            return;
        }
        router.push("/");
    }, [router, session, status])

    useEffect(() => {
        if (id) {
            fetchMovie().then((movie) => {
                setMovie(movie!);

                // Set file video
                if (movie?.video_path) {
                    setVideoPath(movie?.video_path!);
                }
            }).catch((error) => {
                setNotifyState({
                    open: true,
                    message: error.message.message,
                    vertical: "top",
                    horizontal: "right",
                    severity: "error"
                });
            });
        }

    }, [id, router]);

    useEffect(() => {
        if (movie.type_code !== undefined && movie.type_code !== "") {
            const checks: GenreType[] = [];

            // Check genre from TMDB, use genre name to compare instead of id
            fetchGenres().then((results) => {
                results?.forEach((g) => {
                    if (movie?.genres.some(mg => mg.name === g.name)) {
                        checks.push({id: g.id, name: g.name, type_code: g.type_code, checked: true});
                    } else {
                        checks.push({id: g.id, name: g.name, type_code: g.type_code, checked: false});
                    }
                });

                setMovie({
                    ...movie,
                    genres: checks,
                } as MovieType);

                // Check new genres
                const nGenres: string[] = [];
                const checked = checks?.filter(g => g.checked);

                movie?.genres.forEach((mg) => {
                    if (!checked!.some(g => g.name === mg.name)) {
                        nGenres.push(mg.name);
                    }
                });
                setNewGenres(nGenres);
            });
        }
    }, [id, movie.type_code]);

    const handleSubmit = (event) => {
        event.preventDefault();

        let errors: any = [];
        let required = [
            {field: movie.title, name: "title", label: "Title"},
            {field: movie.type_code, name: "type_code", label: "Type Movie"},
            {field: movie.release_date, name: "release_date", label: "Release Date"},
            {field: movie.runtime, name: "runtime", label: "Runtime"},
            {field: movie.description, name: "description", label: "Description"},
            {field: movie.mpaa_rating, name: "mpaa_rating", label: "MPAA Rating"},
        ];
        required.forEach(function ({field, label}: any) {
            if (field === "" || field === undefined) {
                errors.push(label);
            }
        });

        // Check genres
        if (!movie.genres.some(g => g.checked)) {
            setIsOpenAlertDialog(true);
            errors.push("Genres");
        }

        if (errors.length > 0) {
            setNotifyState({
                open: true,
                message: `Fill value for ${errors.join(", ")}`,
                vertical: "bottom",
                horizontal: "center",
                severity: "warning"
            });
            return false;
        }

        // Set id to undefined value for adding new movie
        movie.id = undefined;
        triggerMovie(movie).then((data) => {
            if (data.message === "ok") {
                setNotifyState({
                    open: true,
                    message: "Movie Saved",
                    vertical: "top",
                    horizontal: "right",
                    severity: "success"
                });

                (async () => {
                    await sleep(1500);
                    await router.push("/admin/manage-catalogue");
                })();
            }
        }).catch((error: ClientError) => {
            setNotifyState({
                open: true,
                message: error.message.message,
                vertical: "top",
                horizontal: "right",
                severity: "error"
            });
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

    const handleChooseVideoFileClick = () => {
        videoFileRef.current.click();
    };

    const handleVideoFileChange = event => {
        const fileObj = event.target.files && event.target.files[0];
        if (!fileObj) {
            return;
        }

        // Reset file input
        event.target.value = null;

        setVideoFile(fileObj);

        const formData = new FormData();
        formData.append("video", fileObj);

        uploadVideo(formData).then((result) => {
            if (result.fileName) {
                setVideoPath(result.fileName);

                setMovie({
                    ...movie,
                    video_path: result.fileName.split(".")[0],
                });

                setNotifyState({
                    open: true,
                    message: "Video Uploaded",
                    vertical: "top",
                    horizontal: "right",
                    severity: "info"
                });
            }
        }).catch((error) => {
            setNotifyState({
                open: true,
                message: error.message.message,
                vertical: "top",
                horizontal: "right",
                severity: "error"
            });
        });
    };

    const handleRemoveVideoFileClick = () => {
        if (videoPath !== "") {
            const paths = videoPath.split("/");
            const fileName = paths[paths.length - 1];
            removeVideo({
                fileName: fileName,
            }).then((result) => {
                if (result.message === "ok") {
                    setVideoFile(null);
                    setVideoPath("");

                    setMovie({
                        ...movie,
                        video_path: "",
                    });

                    setNotifyState({
                        open: true,
                        message: "Video Removed",
                        vertical: "top",
                        horizontal: "right",
                        severity: "info"
                    });
                } else {
                    setNotifyState({
                        open: true,
                        message: `Cannot remove video, ${result.result}`,
                        vertical: "top",
                        horizontal: "right",
                        severity: "info"
                    });
                }
            }).catch((error) => {
                setNotifyState({
                    open: true,
                    message: error.message.message,
                    vertical: "top",
                    horizontal: "right",
                    severity: "error"
                });
            })
        }
    };

    const handleAddNewGenresClick = () => {
        addNewGenres({
            genres: newGenres,
        }).then((result) => {
            if (result.message === "ok") {
                setNewGenres([]);

                setNotifyState({
                    open: true,
                    message: "New Genres Added",
                    vertical: "top",
                    horizontal: "right",
                    severity: "info"
                });

                fetchGenres().then((result) => {

                });
            } else {
                setNotifyState({
                    open: true,
                    message: `Cannot add new genres, ${result.message}`,
                    vertical: "top",
                    horizontal: "right",
                    severity: "info"
                });
            }
        }).catch((error: ClientError) => {
            setNotifyState({
                open: true,
                message: error.message.message,
                vertical: "top",
                horizontal: "right",
                severity: "error"
            });
        })
    }

    return (
        <>
            <NotifySnackbar state={notifyState} setState={setNotifyState}/>

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
                />
            }
            <Stack spacing={2}>

                <Box sx={{p: 1, m: 1}}>
                    <Typography variant="h4">Add Reference Movie</Typography>
                </Box>
                <Divider/>
                <Box sx={{display: "flex", justifyContent: "center", p: 1, m: 1, width: 1}}>
                    <form onSubmit={handleSubmit}>
                        <Grid container spacing={2}>
                            <input type="hidden" name="id" defaultValue={movie.id} id="id" readOnly={true}></input>
                            <Grid item xs={8}>
                                <TextField
                                    fullWidth
                                    label="Title"
                                    variant="outlined"
                                    value={movie.title}
                                    onChange={e => handleChange(e, "title")}
                                />
                            </Grid>

                            <Grid item xs={4}>
                                <FormControl>
                                    <FormLabel>Movie Type</FormLabel>
                                    <RadioGroup
                                        row
                                        value={movie.type_code}
                                        onChange={(e) => handleChange(e, "type_code")}
                                    >
                                        {movieTypes.map((t, index) => {
                                            let label;
                                            if (t === "MOVIE") {
                                                label = "Movie";
                                            } else if (t === "TV") {
                                                label = "TV Series";
                                            }
                                            return (
                                                <FormControlLabel
                                                    disabled
                                                    key={`${t}-${index}`}
                                                    value={t}
                                                    control={<Radio/>}
                                                    label={label}
                                                />
                                            );
                                        })}
                                    </RadioGroup>
                                </FormControl>

                            </Grid>


                            <Grid item xs={4}>
                                <TextField
                                    fullWidth
                                    variant="outlined"
                                    label="Release Date"
                                    type="date"
                                    name="release_date"
                                    value={format(new Date(movie.release_date!), "yyyy-MM-dd")}
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
                                    {mpaaOptions && mpaaOptions.map((o) =>
                                        <MenuItem key={o.id} value={o.code}>{o.name}</MenuItem>
                                    )}
                                </TextField>
                            </Grid>

                            <Grid item xs={12}>
                                <TextField
                                    fullWidth
                                    label="Image Path"
                                    variant="outlined"
                                    value={movie.image_path}
                                    onChange={e => handleChange(e, "image_path")}
                                />
                            </Grid>

                            <Grid item xs={12}>
                                <input
                                    ref={videoFileRef}
                                    hidden={true}
                                    type="file"
                                    name="video"
                                    onChange={handleVideoFileChange}
                                />
                                <Stack spacing={2} direction="row">
                                    <Box sx={{display: "flex", alignItems: "center"}}>
                                        <Typography variant="subtitle1">Upload Video</Typography>
                                    </Box>

                                    <Button variant="contained" onClick={handleChooseVideoFileClick}>
                                        Choose File
                                    </Button>

                                    <Box sx={{display: "flex", alignItems: "center"}}>
                                        <Typography>{videoFile?.name}</Typography>
                                    </Box>

                                    {videoPath !== "" &&
                                        <>
                                            <Box sx={{display: "flex", alignItems: "center"}}>
                                                <Link
                                                    href={`${process.env.NEXT_PUBLIC_URL}/video/upload/${videoPath}`}
                                                    target="_blank"
                                                >
                                                    {
                                                        videoPath.split("/").reverse()[0]
                                                    }
                                                </Link>
                                            </Box>
                                            <IconButton aria-label="delete" color="error"
                                                        onClick={handleRemoveVideoFileClick}>
                                                <RemoveCircle/>
                                            </IconButton>
                                        </>
                                    }
                                </Stack>

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
                        {newGenres.length > 0 &&
                            <Stack direction="row" spacing={3} sx={{display: "flex", alignItems: "center"}}>

                                <Typography>New Genres for Reference Movie</Typography>
                                <Box sx={{p: 1, border: 1, borderRadius: 1, borderWidth: 1}}>
                                    {newGenres.map((g) => {
                                        return (
                                            <Chip label={g} color="warning" sx={{p: 1, m: 1}}/>
                                        );
                                    })}
                                </Box>
                                <Button variant="contained"
                                        onClick={handleAddNewGenresClick}>
                                    Add New Genres
                                </Button>
                            </Stack>
                        }

                        <Grid item xs={12}>
                            <FormGroup>
                                <Grid container spacing={1}>
                                    {movie.genres && movie.genres.length > 1 && (
                                        <>
                                            {Array.from(movie.genres).map((g, index) => (
                                                <Grid key={g.id} item xs={2} sx={{m: 1}}>
                                                    <FormControlLabel
                                                        label={g.name}
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
                            </Stack>
                        </Box>
                    </form>
                </Box>
            </Stack>
        </>
    );
};

export default ReferenceMovie;
