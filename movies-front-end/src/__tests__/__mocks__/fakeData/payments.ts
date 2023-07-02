import { CustomPaymentType } from "@/types/movies";
import { Direction, PageType } from "@/types/page";

export const fakePayments: PageType<CustomPaymentType> = {
  size: 5,
  page: 0,
  sort: {
    orders: [
      {
        property: "created_at",
        direction: Direction.ASC
      }
    ]
  },
  total_elements: 5,
  total_pages: 1,
  content: [
    {
      id: 3,
      type_code: "MOVIE",
      movie_title: "Doctor Strange in the Multiverse of Madness",
      season_name: "",
      episode_name: "",
      provider: "STRIPE",
      payment_method: "card",
      amount: 235,
      currency: "usd",
      status: "succeeded",
      created_at: "2023-06-05T16:32:54.76015Z"
    },
    {
      id: 4,
      type_code: "TV",
      movie_title: "Raiders of the Lost Ark",
      season_name: "Season 1",
      episode_name: "Episode 2",
      provider: "STRIPE",
      payment_method: "card",
      amount: 55,
      currency: "usd",
      status: "succeeded",
      created_at: "2023-06-05T17:22:04.933227Z"
    },
    {
      id: 5,
      type_code: "MOVIE",
      movie_title: "Guy Ritchie's The Covenant",
      season_name: "",
      episode_name: "",
      provider: "STRIPE",
      payment_method: "card",
      amount: 99,
      currency: "usd",
      status: "succeeded",
      created_at: "2023-06-05T18:00:06.867662Z"
    },
    {
      id: 6,
      type_code: "TV",
      movie_title: "Highlander",
      season_name: "Season 1",
      episode_name: "Episode 1",
      provider: "STRIPE",
      payment_method: "card",
      amount: 51,
      currency: "usd",
      status: "succeeded",
      created_at: "2023-06-05T18:00:06.867662Z"
    },
    {
      id: 12,
      type_code: "MOVIE",
      movie_title: "Doctor Strange",
      season_name: "",
      episode_name: "",
      provider: "STRIPE",
      payment_method: "card",
      amount: 230,
      currency: "usd",
      status: "succeeded",
      created_at: "2023-06-08T17:45:42.406271Z"
    }
  ]
}