import {withoutRole} from "src/libs/auth";

const handler = withoutRole("banned", async (req, res, jwt) => {
    let {id, movieId} = req.query;

    const username = jwt.id;

    const requestOptions = {
        method: "GET",
    }

    try {
        const response = await fetch(`${process.env.API_BASE_URL}/payments/stripe/${id}/verification?username=${username}&movieId=${movieId}`,
            requestOptions
        );

        res.status(response.status).json(await response.json());
    } catch (error) {
        res.status(500).json({message: "server error"});
    }

});

export default handler;