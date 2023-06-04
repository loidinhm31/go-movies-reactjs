import {withoutRole} from "src/libs/auth";

const handler = withoutRole("banned", async (req, res, token) => {
    let {episodeId} = req.query;

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Authorization", `Bearer ${token.accessToken}`);

    const requestOptions = {
        method: "GET",
        headers: headers,
    }

    try {
        const response = await fetch(`${process.env.API_BASE_URL}/auth/movies?episodeId=${episodeId}`,
            requestOptions
        );

        res.status(response.status).json(await response.json());
    } catch (error) {
        res.status(500).json({message: "server error"});
    }
});

export default handler;