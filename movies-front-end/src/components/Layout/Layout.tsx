// https://nextjs.org/docs/basic-features/layouts

import type {NextPage} from "next";
import {SideMenuLayout} from "./SideMenuLayout";
import {Header} from "../Header";
import {Footer} from "../Footer";
import {Box, Grid} from "@mui/material";
import DashboardIcon from "@mui/icons-material/Dashboard";
import LibraryBooksIcon from "@mui/icons-material/LibraryBooks";
import SearchIcon from "@mui/icons-material/Search";
import HomeIcon from "@mui/icons-material/Home";
import LiveTvIcon from "@mui/icons-material/LiveTv";
import ClassIcon from "@mui/icons-material/Class";

export type NextPageWithLayout<P = unknown, IP = P> = NextPage<P, IP> & {
    getLayout?: (page: React.ReactElement) => React.ReactNode;
};

export const getDefaultLayout = (page: React.ReactElement) => (
    <Grid container>
        <Header/>
        <SideMenuLayout
            menuButtonOptions={[
                {
                    label: "Home",
                    pathname: "/home",
                    desc: "Home",
                    icon: HomeIcon,
                },
                {
                    label: "Movies",
                    pathname: "/movies",
                    desc: "Movies",
                    icon: LiveTvIcon,
                },
                {
                    label: "Genres",
                    pathname: "/genres",
                    desc: "Genres",
                    icon: ClassIcon,
                },
            ]}
            adminMenuButtonOptions={[
                {
                    label: "Dashboard",
                    pathname: "/dashboard",
                    desc: "Dashboard",
                    icon: DashboardIcon,
                },
                {
                    label: "Manage Catalogue",
                    pathname: "/manage-catalogue",
                    desc: "Manage Catalogue",
                    icon: LibraryBooksIcon,
                },
                {
                    label: "Search",
                    pathname: "/search",
                    desc: "Search",
                    icon: SearchIcon,
                },
            ]}
        >
            <Box>
                {page}
            </Box>
            <Footer/>
        </SideMenuLayout>
    </Grid>
);