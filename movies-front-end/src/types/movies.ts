export type MovieType = {
    id?: number,
    title: string,
    description: string,
    release_date: string | null,
    runtime: number,
    mpaa_rating: string,
    image?: string,
    genres: GenreType[],
}

export type GenreType = {
    id: number,
    genre: string,
    checked: boolean,
}