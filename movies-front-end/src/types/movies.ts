export type MovieType = {
    id: number,
    genres: GenreType[],
    title: string,
    description: string,
    release_date: Date,
    runtime: number,
    mpaa_rating: string,
    image: string,
}

export type GenreType = {
    id: any,
    genre: string,
}