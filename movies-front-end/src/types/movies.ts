export type MovieType = {
    id?: number,
    title: string,
    description: string,
    release_date: Date,
    runtime: number,
    mpaa_rating: string,
    image?: string,
    genres: GenreType[],
    genres_array?: any[],
}

export type GenreType = {
    id: number,
    genre: string,
    checked: boolean,
}