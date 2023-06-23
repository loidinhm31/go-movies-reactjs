import { rest } from "msw";

import { fakeEpisodes } from "@/__tests__/__mocks__/fakeData/episodes";
import { fakeGenres } from "@/__tests__/__mocks__/fakeData/genres";
import { fakeMoviePage, fakeTvPage } from "@/__tests__/__mocks__/fakeData/movies";

export const handlers = [
  rest.get("/api/v1/movies", (req, res, ctx) => {
    const type = req.url.searchParams.get("type");

    let page;
    if (type === "MOVIE") {
      page = fakeMoviePage;
    } else if (type === "TV") {
      page = fakeTvPage;
    }

    return res(
      ctx.status(200),
      ctx.body(JSON.stringify(page))
    );
  }),

  rest.get("/api/v1/movies/:id", (req, res, ctx) => {
    const { id } = req.params;

    let movie;
    if (Number(id) === 1 || Number(id) === 2) {
      movie = fakeMoviePage.content.filter((value) => value.id === Number(id))[0];
    } else {
      movie = fakeTvPage.content.filter((value) => value.id === Number(id))[0];
    }

    return res(
      ctx.status(200),
      ctx.body(JSON.stringify(movie))
    );
  }),

  rest.get("/api/v1/episodes/:id", (req, res, ctx) => {
    const { id } = req.params;

    let episode = fakeEpisodes.filter((e) => e.id === Number(id))[0];

    return res(
      ctx.status(200),
      ctx.body(JSON.stringify(episode))
    );
  }),

  rest.get("/api/v1/views/:id", (req, res, ctx) => {
    const { id } = req.params;

    return res(
      ctx.status(200),
      ctx.json({ "message": "OK", "views": 5 })
    );
  }
  ),

  rest.get("api/v1/collections/check", (req, res, ctx) => {
    const type = req.url.searchParams.get("type");
    const refId = req.url.searchParams.get("refId");

    let resJson;
    if (Number(refId) === 1) { // In collection
      resJson = { "user_id": 1, "movie_id": 1 };
    } else if (Number(refId) === 2) { // Not in collection
      resJson = {};
    }

    return res(
      ctx.status(200),
      ctx.json(resJson)
    );
  }),

  rest.post("/api/v1/collections", (req, res, ctx) => {
    return res(
      ctx.status(200),
      ctx.json({ message: "ok" }
      )
    );
  }),

  rest.delete("/api/v1/collections", (req, res, ctx) => {
    return res(
      ctx.status(200),
      ctx.json({ message: "ok" }
      )
    );
  }),


  rest.get("/api/v1/payments/check", (req, res, ctx) => {
    const type = req.url.searchParams.get("type");
    const refId = req.url.searchParams.get("refId");

    let resJson;
    if (Number(refId) === 1) { // In Payment
      resJson = {
        "ref_id": 1,
        "type_code": "MOVIE"
      };
    } else if (Number(refId) === 2) { // Not in Payment
      resJson = {
        "ref_id": 0,
        "type_code": ""
      };
    }

    return res(
      ctx.status(200),
      ctx.json(resJson)
    );
  }),

  rest.get("/api/v1/admin/dashboard/movies/genres", (req, res, ctx) => {

    const data = [
      { name: "G1", type_code: "MOVIE", count: 10 },
      { name: "G2", type_code: "MOVIE", count: 15 },
      { name: "G3", type_code: "MOVIE", count: 5 },
    ];
    return res(
      ctx.status(200),
      ctx.json({ data })
    );
  }),

  rest.post("api/v1/admin/dashboard/views/genres", (req, res, ctx) => {
    return res(
      ctx.status(200),
      ctx.json({
        "data": [
          {
            "year": "2023",
            "month": "3",
            "name": "Adventure",
            "count": 2,
            "cumulative": 2
          },
          {
            "year": "2023",
            "month": "4",
            "name": "Adventure",
            "count": 4,
            "cumulative": 6
          },
          {
            "year": "2023",
            "month": "6",
            "name": "Adventure",
            "count": 1,
            "cumulative": 7
          }
        ]
      })
    );
  }),

  rest.post("api/v1/admin/dashboard/movies/genres/release-date", (req, res, ctx) => {
    return res(
      ctx.status(200),
      ctx.json({
        "data": [
          {
            "year": "2022",
            "month": "7",
            "name": "Action",
            "count": 1,
            "cumulative": 1
          },
          {
            "year": "2022",
            "month": "11",
            "name": "Action",
            "count": 1,
            "cumulative": 2
          },
          {
            "year": "2023",
            "month": "3",
            "name": "Action",
            "count": 1,
            "cumulative": 3
          }
        ]
      })
    );
  }),

  rest.post("/api/v1/admin/dashboard/views", (req, res, ctx) => {
    return res(
      ctx.status(200),
      ctx.json({
        "data": [
            {
                "year": "2023",
                "month": "3",
                "count": 2
            },
            {
                "year": "2023",
                "month": "4",
                "count": 7
            },
            {
                "year": "2023",
                "month": "6",
                "count": 11
            }
        ]
    }),
    )
  }),

  rest.get("/api/v1/genres", (req, res, ctx) => {
    const type = req.url.searchParams.get("type");

    let genres;
    if (type === "MOVIE") {
      genres = fakeGenres;
    }

    return res(
      ctx.status(200),
      ctx.body(JSON.stringify(genres))
    );
  })
];
