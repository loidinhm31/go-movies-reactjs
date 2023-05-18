import {withRole} from "../../../../../libs/auth";

const handler = withRole("admin", async (req, res, token) => {
    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Authorization", `Bearer ${token.accessToken}`)

    const requestOptions = {
        method: "GET",
        headers: headers,
    }

    const response = await fetch(`${process.env.API_BASE_URL}/auth/roles`,
        requestOptions
    );
    if (response.ok) {
        const pageResult = await response.json();
        res.status(200).json(pageResult);
    } else {
        res.status(response.status).json(await response.json())
    }
});

export default handler;