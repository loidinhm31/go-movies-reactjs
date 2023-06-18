import { withAnyRole } from "src/libs/auth";
import { MovieType } from "src/types/movies";

const handler = withAnyRole(["admin", "moderator"], async (req, res, token) => {
    const data: MovieType = req.body;

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Authorization", `Bearer ${token.accessToken}`);

    // Define date format
    data.release_date = new Date(data.release_date!).toISOString();
    let method = "PUT";

    if (data.id !== undefined) {
        method = "PATCH";
    }
    const requestOptions = {
        method: method,
        headers: headers,
        body: JSON.stringify(data),
    };

    try {
        const response = await fetch(`${process.env.API_BASE_URL}/auth/movies`, requestOptions);
        res.status(response.status).json(await response.json());
    } catch (error) {
        res.status(500).json({ message: "server error" });
    }
});

export default handler;
