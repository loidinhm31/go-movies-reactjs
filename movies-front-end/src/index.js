import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import "bootstrap/dist/css/bootstrap.min.css";
import {createBrowserRouter, RouterProvider} from "react-router-dom";
import ErrorPage from "./app/core/components/ErrorPage";
import Home from "./app/core/components/Home";
import Movies from "./app/core/components/Movies";
import Genres from "./app/core/components/Genres";
import EditMovie from "./app/core/components/EditMovie";
import GraphQL from "./app/core/components/GraphQL";
import Login from "./app/core/components/Login";
import Movie from "./app/core/components/Movie";
import ManageCatalogue from "./app/core/components/ManageCatalogue";
import OneGenre from "./app/core/components/OneGenre";

const router = createBrowserRouter([
    {
        path: "/",
        element: <App/>,
        errorElement: <ErrorPage/>,
        children: [
            {index: true, element: <Home/>},
            {
                path: "/movies",
                element: <Movies/>,
            },
            {
                path: "/movies/:id",
                element: <Movie/>,
            },
            {
                path: "/genres",
                element: <Genres/>,
            },
            {
                path: "/genres/:id",
                element: <OneGenre/>,
            },
            {
                path: "/admin/movie/0",
                element: <EditMovie/>,
            },
            {
                path: "/admin/movie/:id",
                element: <EditMovie/>,
            },
            {
                path: "/manage-catalogue",
                element: <ManageCatalogue/>,
            },
            {
                path: "/graphql",
                element: <GraphQL/>,
            },
            {
                path: "/login",
                element: <Login/>,
            }
        ]
    }
])

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(
    <React.StrictMode>
        <RouterProvider router={router}/>
    </React.StrictMode>
);
