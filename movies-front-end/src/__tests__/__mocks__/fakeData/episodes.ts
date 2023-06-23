import { EpisodeType } from "@/types/seasons";


export const fakeEpisodes: EpisodeType[] = [
  {
    id: 1,
    name: "E1: Test episode 1",
    air_date: "2016-11-04T00:00:00Z",
    runtime: 58,
    video_path: "test/videos/1685896967",
    season_id: 5,
    season: {name: "Season 1", description: "test desc s1"},
    price: 55
  },
  {
    id: 2,
    name: "E2: Test episode 2",
    air_date: "2016-11-04T00:00:00Z",
    runtime: 25,
    video_path: "",
    season_id: 5,
    price: 55
  },
  {
    id: 3,
    name: "E3: Test episode 3",
    air_date: "2016-11-04T00:00:00Z",
    runtime: 60,
    video_path: "",
    season_id: 5,
    price: 57
  },
  {
    id: 4,
    name: "E4: Test episode 4",
    air_date: "2016-11-04T00:00:00Z",
    runtime: 59,
    video_path: "",
    season_id: 5,
    price: 60
  },
  {
    id: 5,
    name: "E5: Test episode 5",
    air_date: "2016-11-04T00:00:00Z",
    runtime: 56,
    video_path: "",
    season_id: 5,
    price: 59
  },
];
