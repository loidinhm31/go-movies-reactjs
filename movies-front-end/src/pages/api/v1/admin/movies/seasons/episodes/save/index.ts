import { withAnyRole } from "@/libs/auth";
import { EpisodeType } from "@/types/seasons";

const handler = withAnyRole(["admin", "moderator"], async (req, res, token) => {
  const data: EpisodeType = req.body;

  const headers = new Headers();
  headers.append("Content-Type", "application/json");
  headers.append("Authorization", `Bearer ${token.accessToken}`);

  // Define date format
  data.air_date = new Date(data.air_date!).toISOString();
  let method = "POST";

  if (data.id !== undefined) {
    method = "PUT";
  }
  const requestOptions = {
    method: method,
    headers: headers,
    body: JSON.stringify(data),
  };

  try {
    const response = await fetch(`${process.env.API_BASE_URL}/auth/episodes`, requestOptions);
    res.status(response.status).json(await response.json());
  } catch (error) {
    res.status(500).json({ message: "server error" });
  }
});

export default handler;
