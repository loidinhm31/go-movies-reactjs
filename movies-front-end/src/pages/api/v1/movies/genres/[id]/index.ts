import {Direction, PageType} from "../../../../../../types/page";
import {MovieType} from "../../../../../../types/movies";

const handler = async (req, res) => {
    let {id} = req.query;

    let {pageIndex} = req.query;
    let {pageSize} = req.query;

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
    const data: PageType<MovieType> = req.body;
    if (data.sort) {
        page = data;
    }

    const headers = new Headers();
    headers.append("Content-Type", "application/json")

    const requestOptions = {
        method: "POST",
        headers: headers,
        body: JSON.stringify(page),
    }

    try {
        const response = await fetch(`${process.env.API_BASE_URL}/movies/genres/${id}?page=${pageIndex}&size=${pageSize}`,
            requestOptions
        );
        res.status(response.status).json(await response.json());
    } catch (error) {
        res.status(500).json({message: "server error"});
    }
};

export default handler;