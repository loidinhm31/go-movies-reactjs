import { withAnyRole } from "@/libs/auth";
import { MovieType } from "@/types/movies";

const handler = withAnyRole(["admin", "moderator"], async (req, res, token) => {
  const data: MovieType = req.body;

  const headers = new Headers();
  headers.append("Content-Type", "application/json");
  headers.append("Authorization", `Bearer ${token.accessToken}`);

  const requestOptions = {
    method: "PATCH",
    headers: headers,
  };

  try {
    const response = await fetch(`${process.env.API_BASE_URL}/auth/movies/${data.id}/price`, requestOptions);
    res.status(response.status).json(await response.json());
  } catch (error) {
    res.status(500).json({ message: "server error" });
  }
});

export default handler;
