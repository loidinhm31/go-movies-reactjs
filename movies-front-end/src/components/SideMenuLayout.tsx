import {Box, Container, Divider, Grid, Paper} from "@mui/material";
import Link from "next/link";
import {NextRouter, useRouter} from "next/router";
import {SelectButton, UnselectButton} from "./shared/buttons";

export interface MenuButtonOption {
    label: string;
    pathname: string;
    desc: string;
}

interface SideMenuLayoutProps {
    menuButtonOptions: MenuButtonOption[];
    children: React.ReactNode;
}

interface SideMenuItemProps {
    router: NextRouter;
    item: MenuButtonOption;

}

function SideMenuItem({router, item}: SideMenuItemProps) {
    return (
        <Link href={item.pathname} style={{textDecoration: "none"}}>
            {router.pathname === item.pathname ? (
                <>
                    <SelectButton sx={{width: "100%", borderRadius: 0}}
                                  size="medium">
                        <Box width="100%" display="flex" justifyContent="flex-start">
                            {item.label}
                        </Box>
                    </SelectButton>
                    <Divider variant="middle" />
                </>
            ) : (
                <>
                    <UnselectButton sx={{width: "100%", borderRadius: 0}}>
                        <Box width="100%" display="flex" justifyContent="flex-start">
                            {item.label}
                        </Box>
                    </UnselectButton>
                    <Divider variant="middle" />
                </>
            )}
        </Link>
    );
}

export function SideMenuLayout(props: SideMenuLayoutProps) {
    const router = useRouter();

    const adminMenuButtonOptions: MenuButtonOption[] = [
        {
            label: "Add Movie",
            pathname: "/admin/movie/0",
            desc: "Add Movie",
        },
        {
            label: "Manage Catalogue",
            pathname: "/manage-catalogue",
            desc: "Manage Catalogue",
        },
        {
            label: "GraphQL",
            pathname: "/graphql",
            desc: "GraphQL",
        },
    ];

    return (
        <Container maxWidth="lg">
            <Grid container spacing={2}>
                <Grid item xs={4}>

                    <Box sx={{
                        my: 4,
                        display: 'flex',
                        flexDirection: 'column',
                        justifyContent: 'center',
                        alignItems: 'center',
                    }}>
                        <Paper sx={{width: 230}}>
                            <nav className="grid grid-cols-3 col-span-3 sm:flex sm:flex-col gap-2">
                                {props.menuButtonOptions.map((item, itemIndex) => (
                                    <SideMenuItem key={`${item.label}-${itemIndex}`}
                                                  router={router} item={item} />
                                ))}

                                {adminMenuButtonOptions.map((item, itemIndex) => (
                                    <SideMenuItem key={`${item.label}-${itemIndex}`}
                                                  router={router} item={item} />
                                ))}
                            </nav>
                        </Paper>
                    </Box>
                </Grid>
                <Grid item xs={8}>
                    <Box>
                        {props.children}
                    </Box>
                </Grid>
            </Grid>
        </Container>
    );
}