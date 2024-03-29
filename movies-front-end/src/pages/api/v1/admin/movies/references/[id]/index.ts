import { withAnyRole } from "@/libs/auth";

const handler = withAnyRole(["admin", "moderator"], async (req, res, token) => {
  const { id, type } = req.query;

  const headers = new Headers();
  headers.append("Content-Type", "application/json");
  headers.append("Authorization", `Bearer ${token.accessToken}`);

  const requestOptions = {
    method: "GET",
    headers: headers,
  };

  try {
    const response = await fetch(`${process.env.API_BASE_URL}/auth/references/tmdb/${id}?type=${type}`, requestOptions);
    res.status(response.status).json(await response.json());
  } catch (error) {
    res.status(500).json({ message: "server error" });
  }
});

export default handler;
