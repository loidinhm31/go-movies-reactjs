const handler = async (req, res) => {
    let searchRequest = req.body;

    const headers = new Headers();
    headers.append("Content-Type", "application/json");

    const requestOptions = {
        method: "POST",
        headers: headers,
        body: JSON.stringify(searchRequest),
    }

    const response = await fetch(`${process.env.API_BASE_URL}/views`,
        requestOptions
    );

    if (response.ok) {
        res.status(200).json({});
    } else {
        res.status(response.status).json({})
    }
};

export default handler;