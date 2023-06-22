import { MovieType } from "@/types/movies";
import { withAnyRole } from "@/libs/auth";

const handler = withAnyRole(["admin", "moderator"], async (req, res, token) => {
  const data: MovieType = req.body;

  const headers = new Headers();
  headers.append("Content-Type", "application/json");
  headers.append("Authorization", `Bearer ${token.accessToken}`);

  const requestOptions = {
    method: "POST",
    headers: headers,
    body: JSON.stringify(data),
  };

  try {
    const response = await fetch(`${process.env.API_BASE_URL}/auth/references/tmdb`, requestOptions);
    res.status(response.status).json(await response.json());
  } catch (error) {
    res.status(500).json({ message: "server error" });
  }
});

export default handler;
