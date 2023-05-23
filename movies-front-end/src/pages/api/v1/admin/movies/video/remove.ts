import {withAnyRole} from "src/libs/auth";

const handler = withAnyRole(["admin", "moderator"], async (req, res, token) => {

    let data = req.body;

    const objectKey = data.fileName.split(".")[0];

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Authorization", `Bearer ${token.accessToken}`);

    const requestOptions = {
        method: "DELETE",
        headers: headers,
    }

    const response = await fetch(`${process.env.API_BASE_URL}/auth/integration/videos/${objectKey}`,
        requestOptions
    );

    if (response.ok) {
        res.status(200).json(await response.json());
    } else {
        const message = await response.json()
        res.status(response.status).json(message.message! || "Failed to delete video");
    }
});

export default handler;