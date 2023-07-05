import {
  Box,
  Button,
  Container,
  Divider,
  Grid,
  IconButton,
  InputAdornment,
  Link,
  Stack,
  TextField,
  Typography,
} from "@mui/material";
import { format } from "date-fns";
import { RemoveCircle } from "@mui/icons-material";
import React, { useEffect, useRef, useState } from "react";
import { NotifyState } from "@/components/shared/snackbar";
import { EpisodeType } from "@/types/seasons";
import useSWRMutation from "swr/mutation";
import { del, get, post, postForm } from "@/libs/api";
import EditIcon from "@mui/icons-material/Edit";
import CloseIcon from "@mui/icons-material/Close";
import AlertDialog from "@/components/shared/alert";
import AddIcon from "@mui/icons-material/Add";
import DeleteIcon from "@mui/icons-material/Delete";
import { FileType } from "@/types/movies";

interface EditEpisodeProps {
  id?: number;
  seasonId: number;
  setNotifyState: (state: NotifyState) => void;
  setWasUpdated: (flag: boolean) => void;
}

export default function EditEpisode({ id, seasonId, setNotifyState, setWasUpdated }: EditEpisodeProps) {
  const [isOpenDeleteDialog, setIsOpenDeleteDialog] = useState<boolean>(false);
  const [isConfirmDelete, setIsConfirmDelete] = useState<boolean>(false);

  const [openEdit, setOpenEdit] = useState(false);

  const clearObj: EpisodeType = {
    name: "",
    air_date: format(new Date(), "yyyy-MM-dd"),
    runtime: 0,
    season_id: seasonId,
    price: 0,
  };

  const [episode, setEpisode] = useState<EpisodeType>(clearObj);

  const videoFileRef = useRef<any>(null);
  const [videoFile, setVideoFile] = useState<HTMLInputElement | null>(null);
  const [videoPath, setVideoPath] = useState<string>("");

  const { trigger: fetchEpisode } = useSWRMutation<EpisodeType>(`/api/v1/episodes/${id}`, get);
  const { trigger: triggerEpisode } = useSWRMutation(`/api/v1/admin/movies/seasons/episodes/save`, post);
  const { trigger: deleteEpisode } = useSWRMutation(`/api/v1/admin/movies/seasons/episodes/delete/${id}`, del);
  const { trigger: uploadVideo } = useSWRMutation(`/api/v1/admin/movies/files/upload`, postForm);
  const { trigger: removeVideo } = useSWRMutation(`/api/v1/admin/movies/files/remove`, post);

  useEffect(() => {
    if (id !== undefined) {
      if (openEdit) {
        fetchEpisode()
          .then((result) => {
            setEpisode(result!);

            // Set file video
            if (result?.video_path) {
              setVideoPath(result?.video_path!);
            }
          })
          .catch((error) => {
            setNotifyState({
              open: true,
              message: error.message.message,
              vertical: "top",
              horizontal: "right",
              severity: "error",
            });
          });
      }
    }
  }, [id, seasonId, openEdit]);

  useEffect(() => {
    if (isConfirmDelete) {
      deleteEpisode()
        .then((data) => {
          if (data) {
            setNotifyState({
              open: true,
              message: data.message,
              vertical: "top",
              horizontal: "right",
              severity: "info",
            });
            setWasUpdated(true);
          }
        })
        .catch((error) => {
          setNotifyState({
            open: true,
            message: error.message.message,
            vertical: "top",
            horizontal: "right",
            severity: "error",
          });
        });
    }
  }, [isConfirmDelete]);

  const handleSubmit = () => {
    let errors: any = [];
    let required = [
      { field: episode.name, name: "name", label: "Name" },
      { field: episode.air_date, name: "air_date", label: "Air Date" },
      { field: episode.runtime, name: "runtime", label: "Runtime" },
    ];

    required.forEach(function ({ field, name, label }: any) {
      if (field === "" || field === undefined) {
        errors.push(label);
      }

      if (name === "runtime" && field === 0) {
        errors.push(label);
      }
    });

    if (errors.length > 0) {
      setNotifyState({
        open: true,
        message: `Fill value for ${errors.join(", ")}`,
        vertical: "bottom",
        horizontal: "center",
        severity: "warning",
      });
      return false;
    }

    triggerEpisode(episode)
      .then((data) => {
        if (data) {
          setNotifyState({
            open: true,
            message: "Episode Saved",
            vertical: "top",
            horizontal: "right",
            severity: "success",
          });
          setWasUpdated(true);

          // Clear form for adding new object
          setEpisode(clearObj);
          setVideoFile(null);
          setVideoPath("");
          setOpenEdit(false);
        }
      })
      .catch((error) => {
        setNotifyState({
          open: true,
          message: error.message.message,
          vertical: "top",
          horizontal: "right",
          severity: "error",
        });
      });
  };

  const handleChange = (event, name: string) => {
    let value: string | number = event.target.value;
    if (name === "runtime" || name === "price") {
      value = Number(value);
    } else if (name === "air_date") {
      if (Number.isNaN(new Date(event.target.value).getTime())) return;
    }
    setEpisode({
      ...episode,
      [name]: value,
    });
  };

  const considerDelete = (event) => {
    event.preventDefault();
    setIsOpenDeleteDialog(true);
  };

  const handleChooseVideoFileClick = () => {
    videoFileRef.current.click();
  };

  const handleVideoFileChange = (event) => {
    const fileObj = event.target.files && event.target.files[0];
    if (!fileObj) {
      return;
    }

    // Reset file input
    event.target.value = null;

    setVideoFile(fileObj);

    const formData = new FormData();
    formData.append("file", fileObj);
    formData.append("fileType", FileType.VIDEO);

    uploadVideo(formData)
      .then((result) => {
        if (result.fileName) {
          setVideoPath(result.fileName);

          setEpisode({
            ...episode,
            video_path: result.fileName.split(".")[0],
          });

          setNotifyState({
            open: true,
            message: "Video Uploaded",
            vertical: "top",
            horizontal: "right",
            severity: "info",
          });
        }
      })
      .catch((error) => {
        setNotifyState({
          open: true,
          message: error.message.message,
          vertical: "top",
          horizontal: "right",
          severity: "error",
        });
      });
  };

  const handleRemoveVideoFileClick = () => {
    if (videoPath !== "") {
      const paths = videoPath.split("/");
      const fileName = paths[paths.length - 1];
      removeVideo({
        fileName: fileName,
        fileType: FileType.VIDEO,
      })
        .then((result) => {
          if (result.result === "ok") {
            setVideoFile(null);
            setVideoPath("");

            setEpisode({
              ...episode,
              video_path: "",
            });

            setNotifyState({
              open: true,
              message: "Video Removed",
              vertical: "top",
              horizontal: "right",
              severity: "info",
            });
          } else {
            setNotifyState({
              open: true,
              message: `Cannot remove video, ${result.result}`,
              vertical: "top",
              horizontal: "right",
              severity: "info",
            });
          }
        })
        .catch((error) => {
          setNotifyState({
            open: true,
            message: error.message.message,
            vertical: "top",
            horizontal: "right",
            severity: "error",
          });
        });
    }
  };

  const handleOpenEdit = () => {
    setOpenEdit(!openEdit);
  };

  return (
    <>
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
      <Grid container spacing={2}>
        <Grid item xs={12}>
          <Stack direction="row" spacing={2}>
            {openEdit ? (
              <Button color="warning" variant="contained" onClick={handleOpenEdit}>
                <CloseIcon />
              </Button>
            ) : id ? (
              <Button data-testid="edit-episode" variant="contained" onClick={handleOpenEdit}>
                <EditIcon />
              </Button>
            ) : (
              <Button data-testid="add-episode" variant="contained" onClick={handleOpenEdit}>
                <AddIcon />
              </Button>
            )}

            {id! > 0 && (
              <Button data-testid="delete-episode" variant="contained" color="error" onClick={considerDelete}>
                <DeleteIcon />
              </Button>
            )}
          </Stack>
        </Grid>

        <Grid item xs={12}>
          {openEdit && (
            <Container>
              <Grid container spacing={2}>
                <input type="hidden" name="id" defaultValue={id} id="id" readOnly={true}></input>
                <Grid item xs={8}>
                  <TextField
                    fullWidth
                    label="Name"
                    variant="outlined"
                    value={episode.name}
                    onChange={(e) => handleChange(e, "name")}
                  />
                </Grid>

                <Grid item xs={4}>
                  <TextField
                    fullWidth
                    variant="outlined"
                    label="Air Date"
                    type="date"
                    name="air_date"
                    value={format(new Date(episode.air_date!), "yyyy-MM-dd")}
                    onChange={(e) => handleChange(e, "air_date")}
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
                    value={episode.runtime}
                    onChange={(e) => handleChange(e, "runtime")}
                  />
                </Grid>

                <Grid item xs={3}>
                  <TextField
                    fullWidth
                    label="Price"
                    variant="outlined"
                    type="number"
                    name="price"
                    InputProps={{
                      endAdornment: <InputAdornment position="end">USD</InputAdornment>,
                    }}
                    value={episode.price}
                    onChange={(e) => handleChange(e, "price")}
                  />
                </Grid>

                <Grid item xs={12}>
                  <input data-testid="upload-video" ref={videoFileRef} hidden={true} type="file" name="video" onChange={handleVideoFileChange} />
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
                        <IconButton aria-label="delete" color="error" onClick={handleRemoveVideoFileClick}>
                          <RemoveCircle />
                        </IconButton>
                      </>
                    )}
                  </Stack>
                </Grid>

                <Divider component="div" variant="middle" />

                <Grid item xs={12}>
                  <Box sx={{ display: "flex", justifyContent: "center", m: 2 }}>
                    <Stack direction="row" spacing={2}>
                      <Button variant="contained" onClick={handleSubmit}>
                        Save
                      </Button>
                      {episode.id! > 0 && (
                        <Button variant="contained" color="error" onClick={considerDelete}>
                          Delete Episode
                        </Button>
                      )}
                    </Stack>
                  </Box>
                </Grid>
              </Grid>
            </Container>
          )}
        </Grid>
      </Grid>
    </>
  );
}
