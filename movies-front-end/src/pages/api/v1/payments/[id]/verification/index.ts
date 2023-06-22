import { withoutRole } from "@/libs/auth";

const handler = withoutRole("banned", async (req, res, token) => {
  let { id, type, refId } = req.query;

  const username = token.id;

  const requestOptions = {
    method: "GET",
  };

  try {
    const response = await fetch(
      `${process.env.API_BASE_URL}/payments/stripe/${id}/verification?username=${username}&type=${type}&refId=${refId}`,
      requestOptions
    );

    res.status(response.status).json(await response.json());
  } catch (error) {
    res.status(500).json({ message: "server error" });
  }
});

export default handler;
