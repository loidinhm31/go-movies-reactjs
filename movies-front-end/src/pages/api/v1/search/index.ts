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

    const response = await fetch(`${process.env.API_BASE_URL}/search`,
        requestOptions
    );
    const pageResult = await response.json();
    res.status(200).json(pageResult);
};

export default handler;