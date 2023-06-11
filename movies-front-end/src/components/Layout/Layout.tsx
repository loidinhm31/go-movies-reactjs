// https://nextjs.org/docs/basic-features/layouts

import type {NextPage} from "next";
import {SideMenuLayout} from "src/components/Layout/SideMenuLayout";
import {Header} from "src/components/Header";
import {Footer} from "src/components/Footer";
import {Box, Grid} from "@mui/material";
import DashboardIcon from "@mui/icons-material/Dashboard";
import LibraryBooksIcon from "@mui/icons-material/LibraryBooks";
import SearchIcon from "@mui/icons-material/Search";
import HomeIcon from "@mui/icons-material/Home";
import LiveTvIcon from "@mui/icons-material/LiveTv";
import ClassIcon from "@mui/icons-material/Class";
import LinkIcon from "@mui/icons-material/Link";
import SupervisedUserCircleIcon from "@mui/icons-material/SupervisedUserCircle";

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
                {
                    label: "Advanced Search",
                    pathname: "/search",
                    desc: "Search",
                    icon: SearchIcon,
                },
            ]}
            adminMenuButtonOptions={[
                {
                    label: "Dashboard",
                    pathname: "/admin/dashboard",
                    desc: "Dashboard",
                    icon: DashboardIcon,
                },
                {
                    label: "User Management",
                    pathname: "/admin/manage-user",
                    desc: "User Management",
                    icon: SupervisedUserCircleIcon,
                },
                {
                    label: "Catalogue Management",
                    pathname: "/admin/manage-catalogue",
                    desc: "Catalogue Management",
                    icon: LibraryBooksIcon,
                },
                {
                    label: "Reference Search",
                    pathname: "/admin/references",
                    desc: "Reference Search",
                    icon: LinkIcon,
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