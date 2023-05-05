import {withRole} from "../../../../../../../libs/auth";

const handler = withRole("admin", async (req, res, token) => {
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
        res.status(200).json({});
    } else {
        res.status(response.status).json(await response.json())
    }
});

export default handler;