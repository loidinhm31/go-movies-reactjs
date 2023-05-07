import {withAnyRole} from "src/libs/auth";

const handler = withAnyRole(["admin", "moderator"], async (req, res, token) => {
    let {id} = req.query;
    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Authorization", `Bearer ${token.accessToken}`);

    const requestOptions = {
        method: "DELETE",
        headers: headers,
    }

    const response = await fetch(`${process.env.API_BASE_URL}/private/movies/${id}`,
        requestOptions
    );

    if (response.ok) {
        res.status(200).json({message: "Movie deleted"});
    } else {
        const message = await response.json()
        res.status(response.status).json(message.message! || "Failed to delete movie");
    }
});

export default handler;