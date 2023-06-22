const handler = async (req, res) => {
  const { type } = req.query;

  const headers = new Headers();
  headers.append("Content-Type", "application/json");

  const requestOptions = {
    method: "GET",
    headers: headers,
  };

  try {
    let response;
    if (type) {
      response = await fetch(`${process.env.API_BASE_URL}/genres?type=${type}`, requestOptions);
    } else {
      response = await fetch(`${process.env.API_BASE_URL}/genres`, requestOptions);
    }

    res.status(response.status).json(await response.json());
  } catch (error) {
    res.status(500).json({ message: "server error" });
  }
};

export default handler;
