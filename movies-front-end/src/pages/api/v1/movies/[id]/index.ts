const handler =  async (req, res) => {
    let {id} = req.query;

    const headers = new Headers();
    headers.append("Content-Type", "application/json");

    const requestOptions = {
        method: "GET",
        headers: headers,
    }

    const response = await fetch(`${process.env.API_BASE_URL}/movies/${id}`,
        requestOptions
    );
    const movie = await response.json();
    res.status(200).json(movie);
};

export default handler;