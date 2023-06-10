import {Direction, PageType} from "src/types/page";
import {CustomPaymentType} from "src/types/movies";
import {withoutRole} from "src/libs/auth";

const handler = withoutRole("banned", async (req, res, token) => {
    let {pageIndex, pageSize, type, q} = req.query;

    let page: PageType<any> = {
        sort: {
            orders: [
                {
                    property: "created_at",
                    direction: Direction.DESC
                }
            ]
        }
    }
    const data: PageType<CustomPaymentType> = req.body;
    if (data.sort) {
        page = data;
    }

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Authorization", `Bearer ${token.accessToken}`);

    const requestOptions = {
        method: "POST",
        headers: headers,
        body: JSON.stringify(page),
    }

    try {
        let response = await fetch(`${process.env.API_BASE_URL}/auth/payments?q=${q}&page=${pageIndex}&size=${pageSize}`,
            requestOptions
        );
        res.status(response.status).json(await response.json());
    } catch (error) {
        res.status(500).json({message: "server error"});
    }

});

export default handler;