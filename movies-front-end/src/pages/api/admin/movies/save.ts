import {withRole} from "../../../../libs/auth";

const handler = withRole("admin", async (req, res, token) => {
    let {id} = req.query;

    const data = req.body;

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Authorization", `Bearer ${token.accessToken}`);

    // Assume we are adding a new movies
    let method = "PUT";

    if (id !== undefined) {
        method = "PATCH";
    }
    const requestOptions = {
        method: method,
        headers: headers,
        body: JSON.stringify(data)
    };

    const response = await fetch(`${process.env.API_BASE_URL}/genres`, requestOptions);

    res.status(200).json({});
});

export default handler;
