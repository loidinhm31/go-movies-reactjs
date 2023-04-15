// https://nextjs.org/docs/basic-features/layouts

import type {NextPage} from "next";
import {Box} from "@mui/material";
import {SideMenuLayout} from "./SideMenuLayout";

export type NextPageWithLayout<P = unknown, IP = P> = NextPage<P, IP> & {
    getLayout?: (page: React.ReactElement) => React.ReactNode;
};

export const getDefaultLayout = (page: React.ReactElement) => (
    <SideMenuLayout
        menuButtonOptions={[
            {
                label: "Home",
                pathname: "/home",
                desc: "Home",
            },
            {
                label: "Movies",
                pathname: "/movies",
                desc: "Movies",
            },
            {
                label: "Genres",
                pathname: "/genres",
                desc: "Genres",
            },
        ]}
    >
        <Box>
            {page}
        </Box>
    </SideMenuLayout>
);
export const noLayout = (page: React.ReactElement) => page;
