import {useHasUsername} from "../../hooks/auth/useHasUsername";
import React, {useState} from "react";
import VideoPlayer, {VideoJsOption} from "../Player/VideoPlayer";
import useSWRMutation from "swr/mutation";
import {get, post} from "../../libs/api";
import useSWR from "swr";
import {MovieType} from "../../types/movies";
import {EpisodeType} from "../../types/seasons";

interface WatchEpisodeProps {
    setMutateView: (flag: boolean) => void;
    author: string;
    movieId: number;
    episodeId: number;
}
export default function WatchEpisode({setMutateView, author, movieId, episodeId}: WatchEpisodeProps) {
    const [videoJsOptions, setVideoJsOptions] = useState<VideoJsOption>({
        autoplay: false,
        controls: true,

    });

    const {data: episode, error} = useSWR<EpisodeType>(`/api/v1/episodes/${episodeId}`, get, {
        onSuccess: (result) => {
            if (result.video_path) {
                setVideoJsOptions({
                    ...videoJsOptions,
                    sources: [
                        {
                            src: `${process.env.NEXT_PUBLIC_URL}/video/upload/${result.video_path}`,
                            type: "video/mp4",
                        }
                    ],
                });
            }
        }
    });

    return (
        <>
            <VideoPlayer
                options={videoJsOptions}
                movieId={movieId}
                author={author}
                title={episode?.name!}
                duration={episode?.runtime!}
                setMutateView={setMutateView}
            />
        </>
    );

}