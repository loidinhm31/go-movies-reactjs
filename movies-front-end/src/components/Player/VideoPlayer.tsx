import { Box, Card, CardContent, CardMedia, Typography, useTheme } from "@mui/material";
import "@videojs/themes/dist/sea/index.css";
import { useEffect, useRef } from "react";
import videojs from "video.js";
import "video.js/dist/video-js.css";
export interface VideoJsOption {
    autoplay: boolean;
    controls: boolean;
    sources?: SourceType[];
}


export interface SourceType {
    src: string;
    type: string;
}

export default function VideoPlayer({ options, themeName = "sea", title, duration }) {
    const theme = useTheme();

    const videoRef = useRef<HTMLVideoElement>(null);
    const playerRef = useRef<any>(null);

    useEffect(() => {
        const player = playerRef.current;

        if (!player) {
            const videoElement = videoRef.current;
            if (!videoElement) return;

            playerRef.current = videojs(videoElement, options);
        }

        return () => {
            if (player) {
                player.dispose();
                playerRef.current = null;
            }
        };
    }, [playerRef, videoRef]);

    return (
        <div data-vjs-player>
            <Card >
                <Box sx={{ display: "flex", flexDirection: "column"}}>
                    <CardContent sx={{ flex: "1 0 auto" }}>
                        <Typography component="div" variant="h5">
                            {title}
                        </Typography>
                        <Typography variant="subtitle1" color="text.secondary" component="div">
                            {duration} minutes
                        </Typography>
                        <CardMedia component="video" sx={{aspectRatio: "16:9", height: "650px"}}
                                   ref={videoRef}
                                   className={`video-js vjs-big-play-centered vjs-theme-${themeName}`}
                        />
                    </CardContent>
                </Box>
            </Card>
        </div >
    );
}