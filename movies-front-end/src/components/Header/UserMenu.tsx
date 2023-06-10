import React, {ElementType, useCallback} from "react";
import {signOut, useSession} from "next-auth/react";
import {Avatar, Box, Divider, IconButton, Menu, MenuItem, MenuList, Tooltip, Typography} from "@mui/material";
import LogoutIcon from "@mui/icons-material/Logout";
import SettingsIcon from "@mui/icons-material/Settings";
import CameraRollIcon from "@mui/icons-material/CameraRoll";
import {useHasAnyRole} from "src/hooks/auth/useHasAnyRole";
import {useRouter} from "next/router";
import {MilitaryTech} from "@mui/icons-material";
import PaidIcon from "@mui/icons-material/Paid";

interface MenuOption {
    name: string;
    href: string;
    desc: string;
    icon: ElementType;
    isExternal: boolean;
}

export function UserMenu() {
    const router = useRouter();

    const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
    const open = Boolean(anchorEl);

    const handleSignOut = useCallback(() => {
        signOut({callbackUrl: "/"});
    }, []);

    const {data: session, status} = useSession();
    const isAdminOrMod = useHasAnyRole(["admin", "moderator"]);

    if (!session || status !== "authenticated") {
        return null;
    }

    const options: MenuOption[] = [
        {
            name: "Account Settings",
            href: "/account",
            desc: "Account Settings",
            icon: SettingsIcon,
            isExternal: false
        },
        {
            name: "Your Collections",
            href: "/collections",
            desc: "Your Collections",
            icon: CameraRollIcon,
            isExternal: false
        },
        {
            name: "Your Payments",
            href: "/payments",
            desc: "Your Payments",
            icon: PaidIcon,
            isExternal: false
        }
    ];

    const handleClick = (event: React.MouseEvent<HTMLButtonElement>) => {
        setAnchorEl(event.currentTarget);
    };
    const createHandleClose = () => {
        setAnchorEl(null);
    };

    const handeRoute = (href: string) => {
        router.push(href).then(r => createHandleClose());
    }

    return (
        <Box>
            <Tooltip title="Account settings">
                <IconButton
                    onClick={handleClick}
                    size="small"
                    sx={{ml: 2}}
                    aria-controls={open ? "account-menu" : undefined}
                    aria-haspopup="true"
                    aria-expanded={open ? "true" : undefined}
                >
                    {session.user.name &&
                        <Avatar sx={{width: 32, height: 32}}>{session.user.name!.at(0)!.toUpperCase()}</Avatar>
                    }
                </IconButton>
            </Tooltip>
            <Menu
                anchorEl={anchorEl}
                open={open}
                onClose={createHandleClose}
            >
                <MenuList>
                    <Box sx={{
                        display: "flex", flexDirection: "column", alignItems: "center",
                    }}>
                        {session.user.name ??
                            <MenuItem>
                                <Avatar
                                    sx={{m: 1}}>{session.user.name!.at(0)!.toUpperCase()}</Avatar> {session.user.name}
                            </MenuItem>
                        }
                        {isAdminOrMod &&
                            <MenuItem>
                                <MilitaryTech/>{session.user.role}
                            </MenuItem>
                        }
                    </Box>
                    <Divider/>
                    <MenuList>
                        {options.map((item) => (
                            <Box
                                key={item.name}
                            >
                                <MenuItem onClick={() => handeRoute(item.href)}>
                                    <item.icon className="text-blue-500" aria-hidden="true"/>
                                    <Typography sx={{pl: 1}}>{item.name}</Typography>
                                </MenuItem>
                            </Box>
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