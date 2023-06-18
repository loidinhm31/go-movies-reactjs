import {
    Box,
    CssBaseline,
    Divider,
    IconButton,
    List,
    ListItem,
    ListItemButton,
    ListItemIcon,
    ListItemText,
} from "@mui/material";
import Link from "next/link";
import { NextRouter, useRouter } from "next/router";
import { useTheme } from "@mui/material/styles";
import React, { ElementType, useState } from "react";
import ChevronLeftIcon from "@mui/icons-material/ChevronLeft";
import ChevronRightIcon from "@mui/icons-material/ChevronRight";
import MenuIcon from "@mui/icons-material/Menu";
import { blue } from "@mui/material/colors";
import { Drawer, DrawerHeader } from "src/components/shared/drawer";
import { useHasAnyRole } from "src/hooks/auth/useHasAnyRole";

function MiniDrawer(props: SideMenuLayoutProps) {
    const isPrivilege = useHasAnyRole(["admin", "moderator"]);
    const router = useRouter();
    const theme = useTheme();
    const [open, setOpen] = useState(false);

    const handleDrawerOpen = () => {
        setOpen(true);
    };

    const handleDrawerClose = () => {
        setOpen(false);
    };

    return (
        <Box component="div">
            <Drawer variant="permanent" open={open}>
                <DrawerHeader />
                <DrawerHeader>
                    <IconButton onClick={handleDrawerClose}>
                        {theme.direction === "rtl" ? <ChevronRightIcon /> : <ChevronLeftIcon />}
                    </IconButton>
                    <IconButton
                        color="inherit"
                        aria-label="open drawer"
                        onClick={handleDrawerOpen}
                        sx={{
                            ...(open && { display: "none" }),
                        }}
                    >
                        <MenuIcon />
                    </IconButton>
                </DrawerHeader>
                <Divider />
                <List>
                    {props.menuButtonOptions.map((item, itemIndex) => (
                        <SideMenuItem
                            key={`${item.label}-${itemIndex}`}
                            open={open}
                            router={router}
                            item={item}
                            itemIndex={itemIndex}
                        />
                    ))}
                </List>
                <Divider />
                <List>
                    {isPrivilege &&
                        props.adminMenuButtonOptions!.map((item, itemIndex) => (
                            <SideMenuItem
                                key={`${item.label}-${itemIndex}`}
                                open={open}
                                router={router}
                                item={item}
                                itemIndex={itemIndex}
                            />
                        ))}
                </List>
            </Drawer>
        </Box>
    );
}

export interface MenuButtonOption {
    label: string;
    pathname: string;
    desc: string;
    icon: ElementType;
}

interface SideMenuLayoutProps {
    menuButtonOptions: MenuButtonOption[];
    adminMenuButtonOptions?: MenuButtonOption[];
    children?: React.ReactNode;
}

interface SideMenuItemProps {
    open: boolean;
    router: NextRouter;
    item: MenuButtonOption;
    itemIndex: number;
}

function SideMenuItem({ open, router, item }: SideMenuItemProps) {
    return (
        <Link href={item.pathname} style={{ textDecoration: "none", color: "black" }}>
            <ListItem disablePadding sx={{ display: "block" }}>
                <ListItemButton
                    sx={{
                        minHeight: 48,
                        justifyContent: open ? "initial" : "center",
                        px: 2.5,
                        ...(router.pathname === item.pathname && { backgroundColor: blue[500] }),
                    }}
                >
                    <ListItemIcon
                        sx={{
                            minWidth: 0,
                            mr: open ? 3 : "auto",
                            justifyContent: "center",
                            ...(router.pathname === item.pathname && { color: "white" }),
                        }}
                    >
                        <item.icon />
                    </ListItemIcon>
                    <ListItemText
                        primary={item.label}
                        sx={{
                            opacity: open ? 1 : 0,
                            ...(router.pathname === item.pathname && { color: "white" }),
                        }}
                    />
                </ListItemButton>
            </ListItem>
        </Link>
    );
}

export function SideMenuLayout(props: SideMenuLayoutProps) {
    return (
        <Box sx={{ display: "flex", width: 1 }}>
            <CssBaseline />
            <MiniDrawer
                menuButtonOptions={props.menuButtonOptions}
                adminMenuButtonOptions={props.adminMenuButtonOptions}
            />

            <Box component="main" sx={{ flexGrow: 1, p: 2, width: "1000px" }}>
                <DrawerHeader />
                <Box>{props.children}</Box>
            </Box>
        </Box>
    );
}
