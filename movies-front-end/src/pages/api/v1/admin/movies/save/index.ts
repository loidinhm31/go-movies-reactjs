import {withRole} from "../../../../../../libs/auth";
import moment from "moment";
import {MovieType} from "../../../../../../types/movies";

const handler = withRole("admin", async (req, res, token) => {
    let {id} = req.query;

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

    const response = await fetch(`${process.env.API_BASE_URL}/movies`, requestOptions);

    res.status(200).json({});
});

export default handler;
