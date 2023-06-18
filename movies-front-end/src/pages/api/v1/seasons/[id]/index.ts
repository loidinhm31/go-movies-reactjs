const handler = async (req, res) => {
    let { id } = req.query;

    const headers = new Headers();
    headers.append("Content-Type", "application/json");

    const requestOptions = {
        method: "GET",
        headers: headers,
    };

    const response = await fetch(`${process.env.API_BASE_URL}/seasons/${id}`, requestOptions);
    try {
        res.status(response.status).json(await response.json());
    } catch (error) {
        res.status(response.status).json({ message: "server error" });
    }
};

export default handler;
