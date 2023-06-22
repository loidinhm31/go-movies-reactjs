const handler = async (req, res) => {
  const { movieId } = req.query;

  const headers = new Headers();
  headers.append("Content-Type", "application/json");

  const requestOptions = {
    method: "GET",
    headers: headers,
  };

  try {
    const response = await fetch(`${process.env.API_BASE_URL}/seasons?movieID=${movieId}`, requestOptions);
    res.status(response.status).json(await response.json());
  } catch (error) {
    res.status(500).json({ message: "server error" });
  }
};

export default handler;
