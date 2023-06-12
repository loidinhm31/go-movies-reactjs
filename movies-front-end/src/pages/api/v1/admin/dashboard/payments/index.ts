import {withAnyRole} from "src/libs/auth";

const handler = withAnyRole(["admin", "moderator"], async (req, res, token) => {
    const {type} = req.query;

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Authorization", `Bearer ${token.accessToken}`);

    const requestOptions = {
        method: "GET",
        headers: headers,
    };

    try {
        const response = await fetch(`${process.env.API_BASE_URL}/auth/analysis/payments?type=${type}`, requestOptions);
        res.status(response.status).json(await response.json());
    } catch (error) {
        res.status(500).json({message: "server error"});
    }
});

export default handler;