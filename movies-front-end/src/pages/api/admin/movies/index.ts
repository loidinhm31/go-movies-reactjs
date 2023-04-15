const handler =  async (req, res) => {
    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    // headers.append("Authorization", "Bearer " + jwtToken);

    const requestOptions = {
        method: "GET",
        headers: headers,
    }

    const response = await fetch(`${process.env.API_BASE_URL}/admin/movies`,
        requestOptions
    );
    const genres = await response.json();

    res.status(200).json(genres);
};

export default handler;