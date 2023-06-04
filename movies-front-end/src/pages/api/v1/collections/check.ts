import {withOptionalRole} from "src/libs/auth";

const handler = withOptionalRole("banned", async (req, res, token) => {
    const {refId, type} = req.query;

    const headers = new Headers();
    headers.append("Content-Type", "application/json");

    const requestOptions = {
        method: "GET",
        headers: headers,
    }

    try {
        const response = await fetch(`${process.env.API_BASE_URL}/collections?username=${token.id}&type=${type}&refId=${refId}`,
            requestOptions
        );

        res.status(response.status).json(await response.json());
    } catch (error) {
        res.status(500).json({message: "server error"});
    }

});

export default handler;