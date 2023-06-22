const handler = async (req, res) => {
  let { movieId } = req.query;

  const headers = new Headers();
  headers.append("Content-Type", "application/json");

  const requestOptions = {
    method: "GET",
    headers: headers,
  };

  const response = await fetch(`${process.env.API_BASE_URL}/views/${movieId}`, requestOptions);
  const views = await response.json();
  if (response.ok) {
    res.status(200).json(views);
  } else {
    res.status(response.status);
  }
};

export default handler;
