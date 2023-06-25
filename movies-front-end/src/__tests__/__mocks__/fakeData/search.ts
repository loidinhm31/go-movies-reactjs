import { MovieType } from "@/types/movies";
import { Direction, PageType } from "@/types/page";

export const fakeSearchData: PageType<MovieType> = {
  size: 10,
  page: 0,
  sort: {
    orders: [
      {
        property: "created_at",
        direction: Direction.DESC
      }
    ]
  },
  total_elements: 7,
  total_pages: 1,
  content: [
    {
      id: 1,
      title: "Test movie 1",
      type_code: "MOVIE",
      release_date: "2014-06-05T00:00:00Z",
      runtime: 102,
      mpaa_rating: "G",
      description: "Test desc movie 1",
      image_url: "https://image.tmdb.org/t/p/w200//d13Uj86LdbDLrfDoHR5aDOFYyJC.jpg"
    },
    {
      id: 2,
      title: "Test movie 2",
      type_code: "MOVIE",
      release_date: "2022-05-24T00:00:00Z",
      runtime: 131,
      mpaa_rating: "NC17",
      description: "Test desc movie 2",
      image_url: "https://image.tmdb.org/t/p/w200//62HCnUTziyWcpDaBO2i1DX17ljH.jpg"
    },
    {
      id: 3,
      title: "Test movie 3",
      type_code: "MOVIE",
      release_date: "2022-05-24T00:00:00Z",
      runtime: 131,
      price: 59,
      mpaa_rating: "NC17",
      description: "Test desc movie 3",
      image_url: "https://image.tmdb.org/t/p/w200//62HCnUTziyWcpDaBO2i1DX17ljH.jpg"
    },
    {
      id: 4,
      title: "Test movie 4",
      type_code: "MOVIE",
      release_date: "2022-05-24T00:00:00Z",
      runtime: 131,
      price: 59,
      mpaa_rating: "NC17",
      description: "Test desc movie 4",
      image_url: "https://image.tmdb.org/t/p/w200//62HCnUTziyWcpDaBO2i1DX17ljH.jpg"
    },
    {
      id: 5,
      title: "Test tv 3",
      type_code: "TV",
      release_date: "2013-09-24T00:00:00Z",
      runtime: 43,
      mpaa_rating: "PG",
      description: "Test desc tv 3",
      image_url: "https://image.tmdb.org/t/p/w200//gHUCCMy1vvj58tzE3dZqeC9SXus.jpg",
      genres: [
        {
          id: 16,
          name: "Action & Adventure",
          type_code: "TV"
        },
        {
          id: 20,
          name: "Drama",
          type_code: "TV"
        },
        {
          id: 31,
          name: "Sci-Fi & Fantasy",
          type_code: "TV"
        }
      ]
    },
    {
      id: 6,
      title: "Test tv 4",
      type_code: "TV",
      release_date: "2021-03-19T00:00:00Z",
      runtime: 50,
      mpaa_rating: "PG13",
      description: "Test desc tv 4",
      image_url: "https://image.tmdb.org/t/p/w200//6kbAMLteGO8yyewYau6bJ683sw7.jpg",
      genres: [
        {
          id: 16,
          name: "Action & Adventure",
          type_code: "TV"
        },
        {
          id: 20,
          name: "Drama",
          type_code: "TV"
        }
      ]
    },
    {
      id: 7,
      title: "Test tv 5",
      type_code: "TV",
      release_date: "2018-04-06T00:00:00Z",
      runtime: 24,
      mpaa_rating: "G",
      description: "Test desc tv 5",
      image_url: "https://image.tmdb.org/t/p/w200//mUVZHkJPKDYgDy1dbDdi0Esj9eB.jpg",
      genres: [
        {
          id: 14,
          name: "Comedy",
          type_code: "TV"
        },
        {
          id: 16,
          name: "Action & Adventure",
          type_code: "TV"
        },
        {
          id: 21,
          name: "Kids",
          type_code: "TV"
        },
        {
          id: 23,
          name: "Animation",
          type_code: "TV"
        }
      ]
    }
  ]
};