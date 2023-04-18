import Link from "next/link";
import {UserMenu} from "./UserMenu";
import React from "react";
import {useSession} from "next-auth/react";
import {AppBarProps, Box, Button, Toolbar, Typography} from "@mui/material";
import {styled} from "@mui/material/styles";
import MuiAppBar from "@mui/material/AppBar";
import PersonIcon from "@mui/icons-material/Person";
function AccountButton() {
    const {data: session} = useSession();
    if (session) {
        return;
    }
    return (
        <Link href="/auth/signin" aria-label="Home" style={{textDecoration: "none", color: "white"}}>
            <Box alignItems="center">
                <Button startIcon={<PersonIcon/>} sx={{color: "white"}}>
                    Sign in
                </Button>
            </Box>
        </Link>
    );
}

const AppBar = styled(MuiAppBar, {
    shouldForwardProp: (prop) => prop !== 'open',
})<AppBarProps>(({theme}) => ({
    zIndex: theme.zIndex.drawer + 1,
}));

export function Header() {
    const {data: session} = useSession();
    const homeURL = session ? "/dashboard" : "/";

    return (
        <AppBar position="fixed" sx={{background: "orange"}}>
            <Toolbar component="div" sx={{display: "flex", justifyContent: "space-between"}}>
                <Link href={homeURL} style={{textDecoration: "none", color: "white"}}>
                    <Box sx={{justifyContent: "center", alignItems: "center"}}>
                        <Typography variant="h6" noWrap component="div">
                            Movies App
                        </Typography>
                    </Box>
                </Link>
                <Box sx={{display: "flex", alignItems: "center"}}>
                    <AccountButton/>
                    <UserMenu/>
                </Box>
            </Toolbar>
        </AppBar>

    );
}
