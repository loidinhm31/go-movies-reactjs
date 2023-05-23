export type SeasonType = {
    id?: number,
    name: string,
    air_date?: string | null,
    description: string,
    movie_id?: number,
}

export type EpisodeType = {
    id?: number,
    name: string,
    air_date: string | null,
    runtime: number,
    video_path?: string,
    season_id?: number,
}