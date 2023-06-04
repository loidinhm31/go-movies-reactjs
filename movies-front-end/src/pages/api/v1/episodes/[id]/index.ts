import {withOptionalRole} from "src/libs/auth";

const handler = withOptionalRole("banned", async (req, res, token) => {
    let {id} = req.query;

    const headers = new Headers();
    headers.append("Content-Type", "application/json");

    if (token !== null) {
        headers.append("Authorization", `Bearer ${token.accessToken}`);
    }

    const requestOptions = {
        method: "GET",
        headers: headers,
    }

    const response = await fetch(`${process.env.API_BASE_URL}/episodes/${id}`,
        requestOptions
    );
    try {
        res.status(response.status).json(await response.json());
    } catch (error) {
        res.status(response.status).json({message: "server error"});
    }
});

export default handler;