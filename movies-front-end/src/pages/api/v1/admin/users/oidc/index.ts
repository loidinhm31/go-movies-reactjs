import { withRole } from "@/libs/auth";

const handler = withRole("admin", async (req, res, token) => {
  const headers = new Headers();
  headers.append("Content-Type", "application/json");
  headers.append("Authorization", `Bearer ${token.accessToken}`);

  if (req.method === "GET") {
    const { username } = req.query;

    const requestOptions = {
      method: "GET",
      headers: headers,
    };

    try {
      const response = await fetch(`${process.env.API_BASE_URL}/auth/oidc?username=${username}`, requestOptions);
      res.status(response.status).json(await response.json());
    } catch (error) {
      res.status(500).json({ message: "server error" });
    }
  } else if (req.method === "PUT") {
    const data = req.body;
    const requestOptions = {
      method: "PUT",
      headers: headers,
      body: JSON.stringify(data),
    };

    try {
      const response = await fetch(`${process.env.API_BASE_URL}/auth/users/oidc`, requestOptions);
      res.status(response.status).json(await response.json());
    } catch (error) {
      res.status(500).json({ message: "server error" });
    }
  }
});

export default handler;
