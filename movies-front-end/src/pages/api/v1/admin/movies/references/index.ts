import {MovieType} from "src/types/movies";
import {withAnyRole} from "src/libs/auth";

const handler = withAnyRole(["admin", "moderator"], async (req, res, token) => {

    const data: MovieType = req.body;

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Authorization", `Bearer ${token.accessToken}`)

    const requestOptions = {
        method: "POST",
        headers: headers,
        body: JSON.stringify(data),
    }

    const response = await fetch(`${process.env.API_BASE_URL}/integration/tmdb`,
        requestOptions
    );
    if (response.ok) {
        const pageResult = await response.json();
        res.status(200).json(pageResult);
    } else {
        res.status(response.status).json(await response.json())
    }
});

export default handler;