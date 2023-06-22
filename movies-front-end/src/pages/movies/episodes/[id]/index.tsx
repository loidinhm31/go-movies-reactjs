import { Stack } from "@mui/material";
import { useRouter } from "next/router";
import React from "react";
import WatchEpisode from "@/components/Movie/WatchEpisode";
import { useHasUsername } from "@/hooks/auth/useHasUsername";

function Episode() {
  const router = useRouter();
  let { id, movieId } = router.query;

  const author = useHasUsername();

  return (
    <Stack spacing={2}>
      <WatchEpisode author={author} movieId={parseInt(movieId as string)} episodeId={parseInt(id as string)} />
    </Stack>
  );
}

export default Episode;
