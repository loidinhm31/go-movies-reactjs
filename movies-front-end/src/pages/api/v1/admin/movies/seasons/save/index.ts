import {withAnyRole} from "src/libs/auth";
import {SeasonType} from "../../../../../../../types/seasons";

const handler = withAnyRole(["admin", "moderator"], async (req, res, token) => {
    const data: SeasonType = req.body;

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Authorization", `Bearer ${token.accessToken}`)

    // Define date format
    data.air_date = new Date(data.air_date!).toISOString();
    let method = "PUT";

    if (data.id !== undefined) {
        method = "PATCH";
    }
    const requestOptions = {
        method: method,
        headers: headers,
        body: JSON.stringify(data)
    };

    try {
        const response = await fetch(`${process.env.API_BASE_URL}/auth/seasons`, requestOptions);
        res.status(response.status).json(await response.json());
    } catch (error) {
        res.status(500).json({message: "server error"});
    }
});

export default handler;
