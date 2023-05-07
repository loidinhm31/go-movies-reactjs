import {withAnyRole} from "src/libs/auth";
import moment from "moment";
import {MovieType} from "src/types/movies";

const handler = withAnyRole(["admin", "moderator"], async (req, res, token) => {
    const data: MovieType = req.body;

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Authorization", `Bearer ${token.accessToken}`)

    // Define date format
    data.release_date = moment(data.release_date).toISOString();
    let method = "PUT";

    if (data.id !== undefined) {
        method = "PATCH";
    }
    const requestOptions = {
        method: method,
        headers: headers,
        body: JSON.stringify(data)
    };

    const response = await fetch(`${process.env.API_BASE_URL}/private/movies`, requestOptions);
    if (response.ok) {
        res.status(200).json({message: "Movie saved"});
    } else {
        const message = await response.json()
        res.status(response.status).json(message.message! || "Failed to save movie");
    }
});

export default handler;
