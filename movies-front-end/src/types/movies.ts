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
    season_name?: string,
    episode_name?: string,
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
    user_id: string,
    movie_id?: number,
    episode_id?: number,
    type_code: string,
}

export type PaymentType = {
    id: number,
    user_id: number,
    ref_id: number,
    type_code: string,
    provider: string,
    amount: number,
    currency: string,
    payment_method: string,
    status: string,
    created_at: number,
}

export type CustomPaymentType = {
    id: number,
    type_code: string,
    movie_title: string,
    season_name: string,
    episode_name: string,
    provider: string,
    payment_method: string,
    amount: number,
    currency: string,
    status: string,
    created_at: number,
}