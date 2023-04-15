import {Box, Container, Grid, MenuItem, MenuList, Paper} from "@mui/material";
import Link from "next/link";

export interface MenuButtonOption {
    label: string;
    pathname: string;
    desc: string;
}
interface SideMenuLayoutProps {
    menuButtonOptions: MenuButtonOption[];
    children: React.ReactNode;
}

export function SideMenuLayout(props: SideMenuLayoutProps) {
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
                            <MenuList>
                                {props.menuButtonOptions.map((item, itemIndex) => (
                                    <MenuItem>
                                        <Link href="/home" className="list-group-item list-group-item-action">
                                            {item.label}
                                        </Link>
                                    </MenuItem>
                                ))}

                                <MenuItem>

                                    <Link
                                        href="/admin/movie/0"
                                        className="list-group-item list-group-item-action"
                                    >
                                        Add Movie
                                    </Link>
                                </MenuItem>
                                <MenuItem>
                                    <Link
                                        href="/index"
                                        className="list-group-item list-group-item-action"
                                    >
                                        Manage Catalogue
                                    </Link>
                                </MenuItem>
                                <MenuItem>
                                    <Link
                                        href="/graphql"
                                        className="list-group-item list-group-item-action"
                                    >
                                        GraphQL
                                    </Link>
                                </MenuItem>
                            </MenuList>
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