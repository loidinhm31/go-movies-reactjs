import {withAnyRole} from "src/libs/auth";

const handler = withAnyRole(["admin", "moderator"], async (req, res, token) => {
    const {id} = req.query;    

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Authorization", `Bearer ${token.accessToken}`)

    const requestOptions = {
        method: "GET",
        headers: headers,
    };

    const response = await fetch(`${process.env.API_BASE_URL}/auth/integration/tmdb/${id}`, requestOptions);
    if (response.ok) {
        res.status(200).json(await response.json());
    } else {
        res.status(response.status).json(await response.json())
    }
});

export default handler;
