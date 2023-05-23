import {Box, Card, CardMedia, Paper, Typography} from "@mui/material";
import "@videojs/themes/dist/sea/index.css";
import {useEffect, useRef, useState} from "react";
import videojs from "video.js";
import "video.js/dist/video-js.css";
import useSWRMutation from "swr/mutation";
import {post} from "../../libs/api";

export interface VideoJsOption {
    autoplay: boolean;
    controls: boolean;
    sources?: SourceType[];
}

export interface SourceType {
    src: string;
    type: string;
}

interface VideoPlayerProps {
    options: VideoJsOption;
    movieId: number;
    author: string;
    title: string;
    duration: number;
    themeName?: string;
    setMutateView: (flag: boolean) => void;
}

export default function VideoPlayer({
                                        options,
                                        movieId,
                                        author,
                                        themeName = "sea",
                                        title,
                                        duration,
                                        setMutateView
                                    }: VideoPlayerProps) {
    const videoRef = useRef<HTMLVideoElement>(null);
    const playerRef = useRef<any>(null);

    const [isRecognize, setIsRecognize] = useState(false);

    const {trigger} = useSWRMutation(`/api/v1/users`, post);

    const [wasPlaying, setWasPlaying] = useState(false)

    useEffect(() => {
        if (options.sources?.some((s) => s.src)) {
            const player = playerRef.current;

            if (!player) {
                const videoElement = videoRef.current;
                if (!videoElement) return;

                playerRef.current = videojs(videoElement, options);

                playerRef.current.on(["waiting", "pause"], function () {
                    setWasPlaying(false);
                });

                playerRef.current.on("playing", function () {
                    setWasPlaying(true);
                });
            }
            return () => {
                if (player) {
                    player.dispose();
                    playerRef.current = null;
                    setIsRecognize(false);
                }
            };
        }
    }, [playerRef, videoRef, options]);

    useEffect(() => {
        if (movieId) {
            if (!isRecognize && wasPlaying) {
                const request: any = {
                    movie_id: movieId,
                    viewer: author
                };

                trigger(request).catch((error) => {
                    console.log(error);
                }).finally(() => {
                    setIsRecognize(true);
                    setMutateView(true);
                });
            }
        }
    }, [wasPlaying, movieId, isRecognize]);

    return (
        <div data-vjs-player>
            <Card>
                <Box sx={{display: "flex", flexDirection: "column"}}>
                    <Paper
                        sx={{flex: "1 0 auto", p: 2}}
                    >
                        <Typography component="div" variant="h5">
                            {title}
                        </Typography>
                        <Typography variant="subtitle1" color="text.secondary" component="div">
                            {duration} minutes
                        </Typography>
                        <CardMedia
                            component="video"
                            sx={{aspectRatio: "16:9", height: "650px"}}
                            ref={videoRef}
                            className={`video-js vjs-big-play-centered vjs-theme-${themeName}`}
                        />
                    </Paper>
                </Box>
            </Card>
        </div>
    );
}