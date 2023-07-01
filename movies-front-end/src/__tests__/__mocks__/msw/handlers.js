import { rest } from "msw";

import { fakeEpisodes } from "@/__tests__/__mocks__/fakeData/episodes";
import { fakeGenres } from "@/__tests__/__mocks__/fakeData/genres";
import { fakeMoviePage, fakeTvPage } from "@/__tests__/__mocks__/fakeData/movies";
import { fakeSearchData } from "@/__tests__/__mocks__/fakeData/search";
import { fakeSeasons } from "@/__tests__/__mocks__/fakeData/seasons";
import useSWRMutation from "swr/mutation";
import { patch } from "@/libs/api";
import { fakeUsers } from "@/__tests__/__mocks__/fakeData/users";
import { boolean } from "boolean";

export const handlers = [rest.get("/api/v1/movies", (req, res, ctx) => {
  const type = req.url.searchParams.get("type");
  let page;
  if (type === "MOVIE") {
    page = fakeMoviePage;
  } else if (type === "TV") {
    page = fakeTvPage;
  }

  return res(ctx.status(200), ctx.body(JSON.stringify(page)));
}),

  rest.post("/api/v1/movies", (req, res, ctx) => {
    const type = req.url.searchParams.get("type");
    let page = fakeSearchData;

    return res(ctx.status(200), ctx.body(JSON.stringify(page)));
  }),


  rest.get("/api/v1/movies/:id", (req, res, ctx) => {
    const { id } = req.params;
    if (id === "checkout") {
      return res(ctx.status(200), ctx.json({
        id: 1, title: "Test movie 1", image_url: "test.jpg"
      }));
    } else {
      let movie;
      if (Number(id) === 1 || Number(id) === 2 || Number(id) === 3 || Number(id) === 4) {
        movie = fakeMoviePage.content.filter((value) => value.id === Number(id))[0];
      } else {
        movie = fakeTvPage.content.filter((value) => value.id === Number(id))[0];
      }

      return res(ctx.status(200), ctx.body(JSON.stringify(movie)));
    }
  }),

  rest.delete("/api/v1/admin/movies/delete/:id", (req, res, ctx) => {
    const { id } = req.params;
    if (Number(id) === 1) {
      return res(ctx.status(200), ctx.json({ message: "ok" }));
    } else {
      return res(ctx.status(500));
    }
  }),

  rest.post("/api/v1/admin/movies/save", (req, res, ctx) => {
    return res(ctx.json({ message: "ok" }));
  }),

  rest.put("/api/v1/admin/movies/price", async (req, res, ctx) => {
    const body = await req.json();
    if (body.id === 5) {
      return res(ctx.status(200), ctx.json({ message: "ok" }));
    } else {
      return res(ctx.status(500));
    }
  }),

  rest.post("/api/v1/admin/movies/files/upload", (req, res, ctx) => {
    return res(ctx.status(200), ctx.json({ message: "ok", fileName: "example-file" }));
  }),

  rest.get("/api/v1/episodes/:id", (req, res, ctx) => {
    const { id } = req.params;

    let episode = fakeEpisodes.filter((e) => e.id === Number(id))[0];

    return res(ctx.status(200), ctx.body(JSON.stringify(episode)));
  }),

  rest.get("/api/v1/episodes", (req, res, ctx) => {
    const seasonId = req.url.searchParams.get("seasonId");
    const episodes = fakeEpisodes.filter((e) => e.season_id === Number(seasonId));
    return res(ctx.status(200), ctx.json(episodes));
  }),

  rest.post("/api/v1/admin/movies/seasons/episodes/save", (req, res, ctx) => {
    // Return success message
    return res(ctx.status(200), ctx.json({
      message: "ok"
    }));
  }),

  rest.delete("/api/v1/admin/movies/seasons/episodes/delete/:id", (req, res, ctx) => {
    // Return success message
    return res(ctx.status(200), ctx.json({
      message: "ok"
    }));
  }),

  rest.get("/api/v1/views/:id", (req, res, ctx) => {
    const { id } = req.params;

    return res(ctx.status(200), ctx.json({ "message": "OK", "views": 5 }));
  }),

  rest.get("api/v1/collections/check", (req, res, ctx) => {
    const type = req.url.searchParams.get("type");
    const refId = req.url.searchParams.get("refId");

    let resJson;
    if (Number(refId) === 1) { // In collection
      resJson = { "user_id": 1, "movie_id": 1 };
    } else if (Number(refId) === 2) { // Not in collection
      resJson = {};
    }

    return res(ctx.status(200), ctx.json(resJson));
  }),

  rest.post("/api/v1/collections", (req, res, ctx) => {
    return res(ctx.status(200), ctx.json({ message: "ok" }));
  }),

  rest.delete("/api/v1/collections", (req, res, ctx) => {
    return res(ctx.status(200), ctx.json({ message: "ok" }));
  }),


  rest.get("/api/v1/payments/check", (req, res, ctx) => {
    const type = req.url.searchParams.get("type");
    const refId = req.url.searchParams.get("refId");

    let resJson;
    if (Number(refId) === 1 || Number(refId) === 4) { // In Payment
      resJson = {
        "ref_id": Number(refId), "type_code": type, "status": "succeeded"
      };
    } else if (Number(refId) === 2 || Number(refId) === 3) { // Not in Payment
      resJson = {
        "ref_id": 0, "type_code": ""
      };
    }

    return res(ctx.status(200), ctx.json(resJson));
  }),

  rest.post("/api/v1/payments/intents", (req, res, ctx) => {
    return res(ctx.json({
      clientSecret: "test_client_secret", paymentId: 123456
    }));
  }),

  rest.get("/api/v1/payments/:providerPaymentId/verification", (req, res, ctx) => {
    const { providerPaymentId } = req.params;
    if (Number(providerPaymentId) === 123) {
      return res(ctx.json({
        message: "ok"
      }));
    } else if (Number(providerPaymentId) === 234) {
      return res(ctx.status(500), ctx.json({
        message: "error"
      }));
    }
  }),

  rest.get("/api/v1/admin/dashboard/movies/genres", (req, res, ctx) => {

    const data = [{ name: "G1", type_code: "MOVIE", count: 10 }, {
      name: "G2", type_code: "MOVIE", count: 15
    }, { name: "G3", type_code: "MOVIE", count: 5 }];
    return res(ctx.status(200), ctx.json({ data }));
  }),

  rest.post("api/v1/admin/dashboard/views/genres", (req, res, ctx) => {
    return res(ctx.status(200), ctx.json({
      "data": [{
        "year": "2023", "month": "3", "name": "Adventure", "count": 2, "cumulative": 2
      }, {
        "year": "2023", "month": "4", "name": "Adventure", "count": 4, "cumulative": 6
      }, {
        "year": "2023", "month": "6", "name": "Adventure", "count": 1, "cumulative": 7
      }]
    }));
  }),

  rest.post("api/v1/admin/dashboard/movies/genres/release-date", (req, res, ctx) => {
    return res(ctx.status(200), ctx.json({
      "data": [{
        "year": "2022", "month": "7", "name": "Action", "count": 1, "cumulative": 1
      }, {
        "year": "2022", "month": "11", "name": "Action", "count": 1, "cumulative": 2
      }, {
        "year": "2023", "month": "3", "name": "Action", "count": 1, "cumulative": 3
      }]
    }));
  }),

  rest.post("/api/v1/admin/dashboard/views", (req, res, ctx) => {
    return res(ctx.status(200), ctx.json({
      "data": [{
        "year": "2023", "month": "3", "count": 2
      }, {
        "year": "2023", "month": "4", "count": 7
      }, {
        "year": "2023", "month": "6", "count": 11
      }]
    }));
  }),

  rest.get("api/v1/admin/dashboard/payments", (req, res, ctx) => {
    return res(ctx.status(200), ctx.json({ "total_amount": 670, "total_received": 670 }));
  }),

  rest.get("/api/v1/genres", (req, res, ctx) => {
    const type = req.url.searchParams.get("type");
    let genres;
    if (type !== "") {
      genres = fakeGenres;
    } else {
      genres = fakeGenres.filter((g) => g.type_code === type);
    }

    return res(ctx.status(200), ctx.body(JSON.stringify(genres)));
  }),

  rest.get("/api/v1/ratings", (req, res, ctx) => {
    return res(ctx.status(200), ctx.json([{
      "id": 1, "code": "G", "name": "G"
    }, {
      "id": 2, "code": "PG", "name": "PG"
    }, {
      "id": 3, "code": "PG13", "name": "PG-13"
    }, {
      "id": 4, "code": "R", "name": "R"
    }, {
      "id": 5, "code": "NC17", "name": "NC-17"
    }, {
      "id": 6, "code": "18A", "name": "18A"
    }]));
  }),

  rest.post("/api/v1/search", async (req, res, ctx) => {
    const body = req.json();
    let data = fakeSearchData;
    return res(ctx.status(200), ctx.body(JSON.stringify(data)));
  }),

  rest.get("/api/v1/seasons", (req, res, ctx) => {
    const mockSeasons = fakeSeasons;
    return res(ctx.json(mockSeasons));
  }),

  rest.get("/api/v1/seasons/:id", (req, res, ctx) => {
    const { id } = req.params;

    let season;
    if (Number(id) === 1) {
      season = fakeSeasons.filter((s) => s.id === Number(id))[0];
      return res(ctx.status(200), ctx.json(season));
    } else {
      return res(ctx.status(500), ctx.json({ message: "server error" }));
    }
  }),

  rest.post("/api/v1/admin/movies/seasons/save", (req, res, ctx) => {
    return res(ctx.status(200), ctx.json({ message: "Season saved" }));
  }),

  rest.delete("/api/v1/admin/movies/seasons/delete/:id", (req, res, ctx) => {
    return res(ctx.status(200), ctx.json({ message: "Season deleted" }));
  }),

  rest.post("http://localhost/api/v1/admin/movies/references", async (req, res, ctx) => {
    const body = await req.json();
    let page;
    if (body.type_code === "MOVIE") {
      page = fakeMoviePage.content;
    } else if (body.type_code === "TV") {
      page = fakeTvPage.content;
    }
    return res(ctx.status(200), ctx.body(JSON.stringify(page)));
  }),

  rest.get("/api/v1/collections", (req, res, ctx) => {
    const type = req.url.searchParams.get("type");

    let content;
    if (type === "MOVIE") {
      content = [{
        id: 1,
        title: "Movie 1",
        price: 9.99,
        image_url: "image1.jpg",
        release_date: "2023-01-01",
        description: "Movie 1 description "
      }, {
        id: 2, title: "Movie 2", image_url: "image2.jpg", release_date: "2023-02-01", description: "Movie 2 description"
      }];
    } else if (type === "TV") {
      content = [{
        id: 1,
        title: "Episode 1",
        price: 9.99,
        image_url: "image1.jpg",
        release_date: "2023-01-01",
        description: "Description 1",
        season_name: "Season 1",
        episode_name: "Episode 1"
      }, {
        id: 2,
        title: "Episode 2",
        price: 19.99,
        image_url: "image2.jpg",
        release_date: "2023-02-01",
        description: "Description 2",
        season_name: "Season 2",
        episode_name: "Episode 2"
      }];
    }

    return res(ctx.status(200), ctx.json({
      content: content, total_pages: 2
    }));
  }),

  rest.get("/api/v1/admin/users/oidc", (req, res, ctx) => {
    const username = req.url.searchParams.get("username");
    if (username === "existingUser") {
      return res(ctx.json({
        username: "existingUser",
        email: "existingUser@example.com",
        first_name: "John",
        last_name: "Doe",
        role: { role_code: "" }
      }));
    } else {
      return res(ctx.status(404), ctx.json({ message: "User not found" }));
    }
  }),

  rest.put("/api/v1/admin/users/oidc", (req, res, ctx) => {
    return res(ctx.status(200), ctx.json({ message: "ok" }));
  }),

  rest.get("/api/v1/admin/roles", (req, res, ctx) => {
    return res(ctx.status(200), ctx.json([{ id: 1, role_code: "admin" }, { id: 2, role_code: "general" }]));
  }),

  rest.post("/api/v1/admin/users", (req, res, ctx) => {
    const isNew = req.url.searchParams.get("isNew");
    const users = fakeUsers;
    users.content = fakeUsers.content.filter((u) => u.is_new === boolean(isNew));
    users.total_elements = users.content.length;
    return res(ctx.status(200), ctx.body(JSON.stringify(users)));
  })

];
