import { withoutRole } from "@/libs/auth";

const handler = withoutRole("banned", async (req, res, token) => {
  let data = req.body;

  const headers = new Headers();
  headers.append("Content-Type", "application/json");
  headers.append("Authorization", `Bearer ${token.accessToken}`);

  const requestOptions = {
    method: "POST",
    headers: headers,
    body: JSON.stringify(data),
  };

  try {
    const response = await fetch(`${process.env.API_BASE_URL}/auth/views`, requestOptions);

    res.status(response.status).json(await response.json());
  } catch (error) {
    console.log(error);
  }
});

export default handler;
