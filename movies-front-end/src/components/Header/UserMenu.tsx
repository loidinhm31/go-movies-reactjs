import Link from "next/link";
import React, {ElementType, useCallback} from "react";
import {signOut, useSession} from "next-auth/react";
import {Box, Button, Divider, Menu, MenuItem, MenuList, Typography} from "@mui/material";
import PersonIcon from "@mui/icons-material/Person";
import LogoutIcon from "@mui/icons-material/Logout";
import SettingsIcon from "@mui/icons-material/Settings";
import DashboardIcon from "@mui/icons-material/Dashboard";

interface MenuOption {
    name: string;
    href: string;
    desc: string;
    icon: ElementType;
    isExternal: boolean;
}

export function UserMenu() {
    const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
    const [selectedIndex, setSelectedIndex] = React.useState<number>(1);
    const open = Boolean(anchorEl);

    const handleSignOut = useCallback(() => {
        signOut({callbackUrl: "/"});
    }, []);
    const {data: session, status} = useSession();

    if (!session || status !== "authenticated") {
        return null;
    }

    const options: MenuOption[] = [
        {
            name: "Dashboard",
            href: "/dashboard",
            desc: "Dashboard",
            icon: DashboardIcon,
            isExternal: false
        },
        {
            name: "Account Settings",
            href: "/account",
            desc: "Account Settings",
            icon: SettingsIcon,
            isExternal: false
        }
    ];

    const handleClick = (event: React.MouseEvent<HTMLButtonElement>) => {
        setAnchorEl(event.currentTarget);
    };
    const createHandleClose = (index: number) => () => {
        setAnchorEl(null);
        if (typeof index === "number") {
            setSelectedIndex(index);
        }
    };

    return (
        <Box>
            <Button
                onClick={handleClick}
                startIcon={<PersonIcon/>}
                sx={{color: "white"}}
            >
                {session.user.name || "New User"}
            </Button>
            <Menu
                anchorEl={anchorEl}
                open={open}
                onClose={createHandleClose(-1)}>
                <MenuList>
                    <Box sx={{
                        display: "flex", flexDirection: "column", alignItems: "center",
                        borderRadius: "md", p: 4
                    }}>
                        {session.user.name}
                    </Box>
                    <Divider/>
                    <MenuList>
                        {options.map((item) => (
                            <Link
                                key={item.name}
                                href={item.href}
                                style={{textDecoration: "none", color: "black"}}
                            >
                                <MenuItem>
                                    <item.icon className="text-blue-500" aria-hidden="true"/>
                                    <Typography sx={{pl: 1}}>{item.name}</Typography>
                                </MenuItem>
                            </Link>
                        ))}
                    </MenuList>
                    <Divider/>
                    <MenuItem onClick={handleSignOut}>
                        <LogoutIcon/>
                        <Typography sx={{pl: 1}}>Logout</Typography>
                    </MenuItem>
                </MenuList>
            </Menu>
        </Box>
    );
}