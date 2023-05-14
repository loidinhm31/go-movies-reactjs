export type MovieType = {
    id?: number,
    title: string,
    description: string,
    release_date: string | null,
    runtime: number,
    mpaa_rating: string,
    image_path?: string,
    video_path?: string,
    genres: GenreType[],
    vote_average?: number,
    vote_count?: number,
}

export type GenreType = {
    id: number,
    genre: string,
    checked: boolean,
}