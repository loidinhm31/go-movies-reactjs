import { withoutRole } from "@/libs/auth";
import { CollectionType, MovieType } from "@/types/movies";
import { Direction, PageType } from "@/types/page";

const handler = withoutRole("banned", async (req, res, token) => {
  const headers = new Headers();
  headers.append("Content-Type", "application/json");
  headers.append("Authorization", `Bearer ${token.accessToken}`);

  if (req.method === "GET") {
    let { pageIndex, pageSize, type, q } = req.query;

    let page: PageType<any> = {
      sort: {
        orders: [
          {
            property: "created_at",
            direction: Direction.DESC,
          },
        ],
      },
    };
    const data: PageType<MovieType> = req.body;
    if (data.sort) {
      page = data;
    }

    const requestOptions = {
      method: "POST",
      headers: headers,
      body: JSON.stringify(page),
    };

    try {
      const response = await fetch(
        `${process.env.API_BASE_URL}/auth/collections/page?type=${type}&q=${q}&page=${pageIndex}&size=${pageSize}`,
        requestOptions
      );
      res.status(response.status).json(await response.json());
    } catch (error) {
      res.status(500).json({ message: "server error" });
    }
  } else if (req.method === "POST") {
    const data: CollectionType = req.body;

    const requestOptions = {
      method: "POST",
      headers: headers,
      body: JSON.stringify(data),
    };

    try {
      const response = await fetch(`${process.env.API_BASE_URL}/auth/collections`, requestOptions);
      res.status(response.status).json(await response.json());
    } catch (error) {
      res.status(500).json({ message: "server error" });
    }
  } else if (req.method === "DELETE") {
    const { type, refId } = req.query;

    const requestOptions = {
      method: "DELETE",
      headers: headers,
    };

    try {
      const response = await fetch(
        `${process.env.API_BASE_URL}/auth/collections/refs/${refId}?type=${type}`,
        requestOptions
      );
      res.status(response.status).json(await response.json());
    } catch (error) {
      res.status(500).json({ message: "server error" });
    }
  }
});

export default handler;
