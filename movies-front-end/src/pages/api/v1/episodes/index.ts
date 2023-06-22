import { withOptionalRole } from "@/libs/auth";

const handler = withOptionalRole("banned", async (req, res, token) => {
  const { seasonId } = req.query;

  const headers = new Headers();
  headers.append("Content-Type", "application/json");

  if (token !== null) {
    headers.append("Authorization", `Bearer ${token.accessToken}`);
  }

  const requestOptions = {
    method: "GET",
    headers: headers,
  };

  try {
    const response = await fetch(`${process.env.API_BASE_URL}/episodes?seasonID=${seasonId}`, requestOptions);
    res.status(response.status).json(await response.json());
  } catch (error) {
    res.status(500).json({ message: "server error" });
  }
});

export default handler;
