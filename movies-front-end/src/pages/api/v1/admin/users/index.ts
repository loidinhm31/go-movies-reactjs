import { withRole } from "@/libs/auth";
import { Direction, PageType } from "@/types/page";

const handler = withRole("admin", async (req, res, token) => {
  const headers = new Headers();
  headers.append("Content-Type", "application/json");
  headers.append("Authorization", `Bearer ${token.accessToken}`);

  if (req.method === "POST") {
    let { pageIndex, pageSize, isNew, query } = req.query;

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
    const data: PageType<any> = req.body;
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
        `${process.env.API_BASE_URL}/auth/users/page?page=${pageIndex}&size=${pageSize}&isNew=${isNew}&q=${query}`,
        requestOptions
      );
      res.status(response.status).json(await response.json());
    } catch (error) {
      res.status(500).json({ message: "server error" });
    }
  } else if (req.method === "PATCH") {
    const data = req.body;

    const requestOptions = {
      method: "PATCH",
      headers: headers,
      body: JSON.stringify(data),
    };

    try {
      const response = await fetch(`${process.env.API_BASE_URL}/auth/users/role`, requestOptions);
      res.status(response.status).json(await response.json());
    } catch (error) {
      res.status(500).json({ message: "server error" });
    }
  }
});

export default handler;
