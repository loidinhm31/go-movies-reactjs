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

    await fetch(`${process.env.API_BASE_URL}/movies/${id}`,
        requestOptions
    );

    res.status(200).json({});
});

export default handler;