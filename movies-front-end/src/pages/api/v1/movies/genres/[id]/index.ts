import {useState} from "react";
import {Direction, PageType} from "../../../../../../types/page";
import {MovieType} from "../../../../../../types/movies";

const handler =  async (req, res) => {
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

    const response = await fetch(`${process.env.API_BASE_URL}/movies/genres/${id}?page=${pageIndex}&size=${pageSize}`,
        requestOptions
    );
    const pageResult = await response.json();

    res.status(200).json(pageResult);
};

export default handler;