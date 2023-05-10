import {withAnyRole} from "src/libs/auth";
import {MovieType} from "src/types/movies";

const handler = withAnyRole(["admin", "moderator"], async (req, res, token) => {
    const data: MovieType = req.body;

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Authorization", `Bearer ${token.accessToken}`)

    const requestOptions = {
        method: "POST",
        headers: headers,
        body: JSON.stringify(data)
    };

    const response = await fetch(`${process.env.API_BASE_URL}/analysis/movies/genres/release-date`, requestOptions);
    if (response.ok) {
        res.status(200).json(await response.json());
    } else {
        res.status(response.status).json(await response.json())
    }
});

export default handler