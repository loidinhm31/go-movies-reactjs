export const enum FileType {
    VIDEO = "video",
    IMAGE = "image",
}

export type MovieType = {
    id?: number,
    title: string,
    type_code: string,
    description: string,
    release_date: string | null,
    runtime: number,
    mpaa_rating?: string,
    image_url?: string,
    video_path?: string,
    price?: number,
    genres: GenreType[],
    vote_average?: number,
    vote_count?: number,
}

export type GenreType = {
    id: number,
    name: string,
    type_code: string,
    checked: boolean,
}

export type RatingType = {
    id: number,
    code: string,
    name: string,
}

export type CollectionType = {
    username: string,
    movie_id: number,
}