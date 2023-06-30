import { RemoveCircle } from "@mui/icons-material";
import {
  Box,
  Button,
  CardMedia,
  Checkbox,
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
import { format } from "date-fns";
import { signIn } from "next-auth/react";
import { useRouter } from "next/router";
import { useEffect, useRef, useState } from "react";
import { movieTypes } from "@/components/MovieTypeSelect";
import AlertDialog from "@/components/shared/alert";
import NotifySnackbar, { NotifyState } from "@/components/shared/snackbar";
import { useCheckTokenAndRole } from "@/hooks/auth/useCheckTokenAndRole";
import { del, get, put, post, postForm } from "@/libs/api";
import { FileType, GenreType, MovieType, RatingType } from "@/types/movies";
import useSWRMutation from "swr/mutation";
import PriceChangeIcon from "@mui/icons-material/PriceChange";
import { useDispatch } from "react-redux";
import { setData } from "@/redux/store";

const EditMovie = () => {
  const router = useRouter();
  const isInvalid = useCheckTokenAndRole(["admin", "moderator"]);
  const dispatch = useDispatch();

  const [isOpenAlertDialog, setIsOpenAlertDialog] = useState<boolean>(false);
  const [isOpenDeleteDialog, setIsOpenDeleteDialog] = useState<boolean>(false);
  const [isConfirmDelete, setIsConfirmDelete] = useState<boolean>(false);

  const [notifyState, setNotifyState] = useState<NotifyState>({ open: false, vertical: "top", horizontal: "right" });

  const [movie, setMovie] = useState<MovieType>({
    title: "",
    type_code: "",
    description: "",
    release_date: format(new Date(), "yyyy-MM-dd"),
    runtime: 0,
    price: 0,
    genres: [],
    image_url: ""
  });

  // Get id from the URL
  let { id } = router.query;

  const [ratings, setRatings] = useState<RatingType[]>();
  const [rating, setRating] = useState("");
  const videoFileRef = useRef<any>(null);
  const [videoFile, setVideoFile] = useState<HTMLInputElement | null>(null);
  const [videoPath, setVideoPath] = useState<string>("");

  const imageFileRef = useRef<any>(null);
  const [imageFile, setImageFile] = useState<HTMLInputElement | null>(null);
  const [imageUrl, setImageUrl] = useState<string>("");

  const { trigger: fetchGenres } = useSWRMutation<GenreType[]>(`/api/v1/genres?type=${movie.type_code}`, get);
  const { trigger: fetchMovie } = useSWRMutation<MovieType>(`/api/v1/movies/${id}`, get);
  const { trigger: triggerMovie } = useSWRMutation(`/api/v1/admin/movies/save`, post);
  const { trigger: deleteMovie } = useSWRMutation(`/api/v1/admin/movies/delete/${id}`, del);

  const { trigger: uploadFile } = useSWRMutation(`/api/v1/admin/movies/files/upload`, postForm);
  const { trigger: removeFile } = useSWRMutation(`/api/v1/admin/movies/files/remove`, post);

  const { trigger: updateMoviePrice } = useSWRMutation(`/api/v1/admin/movies/price`, put);

  const { trigger: getRatings } = useSWRMutation<RatingType[]>("/api/v1/ratings", get);

  useEffect(() => {
    if (isInvalid) {
      signIn();
      return;
    }
  }, [isInvalid]);

  useEffect(() => {
    getRatings().then((result) => {
      setRatings(result);
    });

    if (id === undefined) {
      if (movie.type_code !== "") {
        handleFetchGenres();
      }
    } else {
      fetchMovie()
        .then((movie) => {
          setMovie(movie!);

          // Set file video
          if (movie?.video_path) {
            setVideoPath(movie?.video_path!);
          }
        })
        .catch((error) => {
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
      handleFetchGenres();
    }
  }, [id, movie.type_code]);

  useEffect(() => {
    if (rating !== "") {
      setMovie({
        ...movie,
        mpaa_rating: rating
      });
    }
  }, [rating]);

  useEffect(() => {
    if (imageUrl !== "") {
      setMovie({
        ...movie,
        image_url: imageUrl
      });
    }
  }, [imageUrl]);

  useEffect(() => {
    if (isConfirmDelete) {
      deleteMovie()
        .then((data) => {
          if (data) {
            dispatch(setData({
              severity: "info",
              message: `${movie.title} was deleted`
            }));

            router.push("/admin/manage-catalogue");
          }
        })
        .catch((error) => {
          setIsConfirmDelete(false);
          setNotifyState({
            open: true,
            message: error.message.message,
            vertical: "top",
            horizontal: "right",
            severity: "error"
          });
        });
    }
  }, [isConfirmDelete]);

  const handleFetchGenres = () => {
    const checks: GenreType[] = [];
    fetchGenres().then((result) => {
      result?.forEach((g) => {
        if (movie?.genres.some((mg) => mg.name === g.name && mg.type_code == g.type_code)) {
          checks.push({ id: g.id, name: g.name, type_code: g.type_code, checked: true });
        } else {
          checks.push({ id: g.id, name: g.name, type_code: g.type_code, checked: false });
        }
      });
      setMovie({
        ...movie,
        genres: checks
      } as MovieType);
    });
  };

  const handleSubmit = (event) => {
    event.preventDefault();

    let errors: any = [];
    let required = [
      { field: movie.title, name: "title", label: "Title" },
      { field: movie.type_code, name: "type_code", label: "Type Movie" },
      { field: movie.release_date, name: "release_date", label: "Release Date" },
      { field: movie.runtime, name: "runtime", label: "Runtime" },
      { field: movie.description, name: "description", label: "Description" },
      { field: movie.mpaa_rating, name: "mpaa_rating", label: "MPAA Rating" }
    ];

    required.forEach(function({ field, label }: any) {
      if (field === "" || field === undefined) {
        errors.push(label);
      }
    });

    // Check genres
    if (!movie.genres.some((g) => g.checked)) {
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

    triggerMovie(movie)
      .then((data) => {
        if (data) {
          dispatch(setData({
            severity: "success",
            message: "Movie Saved"
          }));
          router.push("/admin/manage-catalogue");
        }
      })
      .catch((error) => {
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
    if (name === "runtime" || name === "price") {
      value = Number(value);
    } else if (name === "release_date") {
      if (Number.isNaN(new Date(event.target.value).getTime())) return;
    }

    setMovie({
      ...movie,
      [name]: value
    });
  };

  const handleCheck = (event, position: number) => {
    let tmpArr = movie.genres;

    tmpArr[position].checked = event.target.checked;

    setMovie({
      ...movie,
      genres: tmpArr
    });
  };

  const confirmDelete = (event) => {
    event.preventDefault();
    setIsOpenDeleteDialog(true);
  };

  const handleChooseVideoFileClick = () => {
    videoFileRef.current.click();
  };

  const handleChooseImageFileClick = () => {
    imageFileRef.current.click();
  };

  const handleFileChange = (event: any, type: FileType) => {
    const fileObj = event.target.files && event.target.files[0];
    if (!fileObj) {
      return;
    }
    // Reset file input
    event.target.value = null;

    const formData = new FormData();
    formData.append("file", fileObj);
    formData.append("fileType", type);

    if (type === FileType.VIDEO) {
      setVideoFile(fileObj);

      uploadFile(formData)
        .then((result) => {
          if (result.fileName) {
            setVideoPath(result.fileName);

            setMovie({
              ...movie,
              video_path: result.fileName.split(".")[0]
            });

            setNotifyState({
              open: true,
              message: "Video Uploaded",
              vertical: "top",
              horizontal: "right",
              severity: "info"
            });
          }
        })
        .catch((error) => {
          setNotifyState({
            open: true,
            message: error.message.message,
            vertical: "top",
            horizontal: "right",
            severity: "error"
          });
        });
    } else if (type === FileType.IMAGE) {
      setImageFile(fileObj);

      uploadFile(formData)
        .then((result) => {
          if (result.fileName) {
            setImageUrl(result.fileName);

            setMovie({
              ...movie,
              image_url: result.fileName
            });

            setNotifyState({
              open: true,
              message: "Image Uploaded",
              vertical: "top",
              horizontal: "right",
              severity: "info"
            });
          }
        })
        .catch((error) => {
          setNotifyState({
            open: true,
            message: error.message.message,
            vertical: "top",
            horizontal: "right",
            severity: "error"
          });
        });
    }
  };

  const handleRemoveFileClick = (fileType: FileType) => {
    let paths;
    if (fileType === FileType.VIDEO) {
      if (videoPath === "") return;

      paths = videoPath.split("/");
    } else if (fileType === FileType.IMAGE) {
      if (movie.image_url!.startsWith("https://image.tmdb.org")) {
        setMovie({
          ...movie,
          image_url: ""
        });
        return;
      }

      if (imageUrl === "") return;

      paths = imageUrl.split("/");
    }

    const fileName = paths[paths.length - 1];
    removeFile({
      fileName: fileName,
      fileType: fileType
    })
      .then((result) => {
        if (result.result === "ok") {
          if (fileType === FileType.VIDEO) {
            setVideoFile(null);
            setVideoPath("");

            setMovie({
              ...movie,
              video_path: ""
            });
          } else if (fileType === FileType.IMAGE) {
            setImageFile(null);
            setImageUrl("");

            setMovie({
              ...movie,
              image_url: ""
            });
          }

          setNotifyState({
            open: true,
            message: `${fileType} removed`,
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
      })
      .catch((error) => {
        setNotifyState({
          open: true,
          message: error.message.message,
          vertical: "top",
          horizontal: "right",
          severity: "error"
        });
      });
  };

  const handleUpdateAveragePrice = () => {
    updateMoviePrice({ id: movie?.id! } as MovieType)
      .then((result) => {
        if (result.message === "ok") {
          setNotifyState({
            open: true,
            message: "Average Price Was Updated",
            vertical: "top",
            horizontal: "right",
            severity: "success"
          });

          fetchMovie()
            .then((movie) => {
              setMovie((prevMovie) => ({
                ...prevMovie,
                price: movie?.price
              }));
            })
            .catch((error) => {
              setNotifyState({
                open: true,
                message: error.message.message,
                vertical: "top",
                horizontal: "right",
                severity: "error"
              });
            });
        }
      })
      .catch((error) => {
        setNotifyState({
          open: true,
          message: error.message.message,
          vertical: "top",
          horizontal: "right",
          severity: "error"
        });
      });
  };

  return (
    <>
      <NotifySnackbar state={notifyState} setState={setNotifyState} />

      {isOpenAlertDialog && (
        <AlertDialog
          open={isOpenAlertDialog}
          setOpen={setIsOpenAlertDialog}
          title={"Error!"}
          description={"You must choose at least one genre!"}
          confirmText={"Agree"}
        />
      )}
      {isOpenDeleteDialog && (
        <AlertDialog
          open={isOpenDeleteDialog}
          setOpen={setIsOpenDeleteDialog}
          title={"Delete Item"}
          description={"You cannot undo this action!"}
          confirmText={"Yes"}
          showCancelButton={true}
          setConfirmDelete={setIsConfirmDelete}
        />
      )}
      <Stack spacing={2}>
        <Box sx={{ p: 1, m: 1 }}>
          <Typography variant="h4">Add/Edit Movie</Typography>
        </Box>
        <Divider />
        <Box sx={{ display: "flex", justifyContent: "center", p: 1, m: 1, width: 1 }}>
          <form onSubmit={handleSubmit}>
            <Grid container spacing={2}>
              <Grid item xs={8}>
                <TextField
                  fullWidth
                  label="Title"
                  variant="outlined"
                  value={movie.title}
                  onChange={(e) => handleChange(e, "title")}
                />
              </Grid>

              <Grid item xs={2}>
                <>
                  {movie.type_code === "MOVIE" && (
                    <TextField
                      fullWidth
                      label="Price"
                      variant="outlined"
                      type="number"
                      name="price"
                      InputProps={{
                        endAdornment: <InputAdornment position="end">USD</InputAdornment>
                      }}
                      value={movie.price || 0}
                      onChange={(e) => handleChange(e, "price")}
                    />
                  )}

                  {movie.type_code === "TV" && (
                    <Stack direction="row" spacing={2}>
                      <IconButton data-testid="updateAvgPrice" color="secondary" onClick={handleUpdateAveragePrice}>
                        <PriceChangeIcon />
                      </IconButton>

                      <TextField
                        disabled
                        fullWidth
                        label="Average Price"
                        variant="outlined"
                        type="number"
                        name="price"
                        InputProps={{
                          endAdornment: <InputAdornment position="end">USD</InputAdornment>
                        }}
                        value={movie.price || 0}
                      />
                    </Stack>
                  )}
                </>
              </Grid>

              <Grid item xs={2}>
                <FormControl>
                  <FormLabel>Movie Type</FormLabel>
                  <RadioGroup row value={movie.type_code} onChange={(e) => handleChange(e, "type_code")}>
                    {movieTypes.map((t, index) => {
                      let label;
                      if (t === "MOVIE") {
                        label = "Movie";
                      } else if (t === "TV") {
                        label = "TV Series";
                      }
                      return <FormControlLabel key={`${t}-${index}`} value={t} control={<Radio />} label={label} />;
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
                  onChange={(e) => handleChange(e, "release_date")}
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
                    endAdornment: <InputAdornment position="end">minutes</InputAdornment>
                  }}
                  value={movie.runtime}
                  onChange={(e) => handleChange(e, "runtime")}
                />
              </Grid>

              <Grid item xs={4}>
                {ratings &&
                  <TextField
                    fullWidth
                    select
                    label="MPAA Rating"
                    variant="outlined"
                    value={rating}
                    onChange={(event) => setRating(event.target.value)}
                  >
                    {
                      ratings.map((o, index) => (
                        <MenuItem key={`${o.id}-${index}`} value={o.code}>
                          {o.name}
                        </MenuItem>
                      ))
                    }
                  </TextField>
                }
              </Grid>


              <Grid container item xs={12} spacing={3} sx={{ display: "flex", alignItems: "center" }}>
                {movie && !movie.image_url && (
                  <>
                    <Grid item xs={6}>
                      <TextField
                        fullWidth
                        label="Image Path"
                        variant="outlined"
                        value={imageUrl}
                        onChange={(e) => setImageUrl(e.target.value)}
                      />
                    </Grid>
                    <Grid container item xs={6} spacing={2} sx={{ display: "flex" }}>
                      <input
                        data-testid="uploadImage"
                        ref={imageFileRef}
                        hidden={true}
                        type="file"
                        name="image"
                        onChange={(event) => handleFileChange(event, FileType.IMAGE)}
                      />
                      <Stack direction="row" spacing={2}>
                        <Box sx={{ display: "flex", alignItems: "center" }}>
                          <Typography variant="subtitle1">Or Upload Image</Typography>
                        </Box>

                        <Button variant="contained" onClick={handleChooseImageFileClick}>
                          Choose File
                        </Button>
                      </Stack>
                    </Grid>
                  </>
                )}

                {movie && movie.image_url && (
                  <>
                    <Grid item>
                      <CardMedia component="img" sx={{ borderRadius: "16px" }} src={movie.image_url} />
                    </Grid>
                    <Grid>
                      <IconButton
                        aria-label="delete"
                        color="error"
                        onClick={() => handleRemoveFileClick(FileType.IMAGE)}
                      >
                        <RemoveCircle />
                      </IconButton>
                    </Grid>
                  </>
                )}
              </Grid>

              {movie && movie.type_code === "MOVIE" && (
                <Grid item xs={12}>
                  <input
                    data-testid="uploadVideo"
                    ref={videoFileRef}
                    hidden={true}
                    type="file"
                    name="video"
                    onChange={(event) => handleFileChange(event, FileType.VIDEO)}
                  />
                  <Stack spacing={2} direction="row">
                    <Box sx={{ display: "flex", alignItems: "center" }}>
                      <Typography variant="subtitle1">Upload Video</Typography>
                    </Box>

                    <Button variant="contained" onClick={handleChooseVideoFileClick}>
                      Choose File
                    </Button>

                    <Box sx={{ display: "flex", alignItems: "center" }}>
                      <Typography>{videoFile?.name}</Typography>
                    </Box>

                    {videoPath !== "" && (
                      <>
                        <Box sx={{ display: "flex", alignItems: "center" }}>
                          <Link
                            href={`${process.env.NEXT_PUBLIC_CLOUDINARY_URL}/video/upload/${videoPath}`}
                            target="_blank"
                          >
                            {videoPath.split("/").reverse()[0]}
                          </Link>
                        </Box>
                        <IconButton
                          aria-label="delete"
                          color="error"
                          onClick={() => handleRemoveFileClick(FileType.VIDEO)}
                        >
                          <RemoveCircle />
                        </IconButton>
                      </>
                    )}
                  </Stack>
                </Grid>
              )}

              <Grid item xs={12}>
                <TextField
                  fullWidth
                  label="Description"
                  variant="outlined"
                  multiline
                  rows={4}
                  value={movie.description}
                  onChange={(e) => handleChange(e, "description")}
                />
              </Grid>

              <Divider component="div" variant="middle" />
            </Grid>

            <Typography sx={{ p: 2 }} variant="h6">
              Genres
            </Typography>
            <Grid item xs={12}>
              <FormGroup>
                <Grid container spacing={1}>
                  {movie.genres &&
                    movie.genres.length > 0 &&
                    movie.genres.map((g, index) => (
                      <Grid key={`${g.id}-${index}`} item xs={2} sx={{ m: 1 }}>
                        <FormControlLabel
                          label={g.name}
                          onChange={(event) => handleCheck(event, index)}
                          value={g.id}
                          control={<Checkbox checked={g.checked === true} />}
                        />
                      </Grid>
                    ))}
                </Grid>
              </FormGroup>
            </Grid>

            <Divider component="div" variant="middle" />

            <Box sx={{ display: "flex", justifyContent: "center", m: 2 }}>
              <Stack direction="row" spacing={2}>
                <Button variant="contained" type="submit">
                  Save
                </Button>
                {movie.id! > 0 && (
                  <Button variant="contained" color="error" href="src/app/core/components#!" onClick={confirmDelete}>
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
