import {SearchRequest} from "../../../../types/search";

const handler = async (req, res) => {
    let searchRequest: SearchRequest = req.body

    const headers = new Headers();
    headers.append("Content-Type", "application/json");

    const requestOptions = {
        method: "POST",
        headers: headers,
        body: JSON.stringify(searchRequest),
    }

    try {
        const response = await fetch(`${process.env.API_BASE_URL}/search`,
            requestOptions
        );
        res.status(response.status).json(await response.json());
    } catch (error) {
        res.status(500).json({message: "server error"});
    }

};

export default handler;