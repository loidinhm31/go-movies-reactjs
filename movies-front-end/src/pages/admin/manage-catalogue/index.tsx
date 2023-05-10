import React, {useEffect, useState} from "react";
import {useRouter} from "next/router";
import {del, post} from "src/libs/api";
import {Box, Button, Divider, Skeleton, Stack, Typography} from "@mui/material";
import {useSession} from "next-auth/react";
import EnhancedTable, {Data} from "../../../components/Tables/EnhancedMoviesTable";
import useSWRMutation from "swr/mutation";
import {Direction, PageType} from "../../../types/page";
import {MovieType} from "../../../types/movies";
import NotifySnackbar, {NotifyState} from "../../../components/shared/snackbar";

const ManageCatalogue = () => {
    const {data: session, status} = useSession();
    const router = useRouter();

    const [page, setPage] = useState<PageType<MovieType> | null>(null);

    const [isConfirmDelete, setIsConfirmDelete] = useState<boolean>(false);
    const [deleteId, setDeleteId] = useState<number | null>();

    const [pageIndex, setPageIndex] = useState(0);
    const [pageSize, setPageSize] = useState(5)
    const [order, setOrder] = useState<Direction>(Direction.ASC);
    const [orderBy, setOrderBy] = useState<keyof Data>("release_date");

    const [notifyState, setNotifyState] = useState<NotifyState>({open: false, vertical: "top", horizontal: "right"});

    // Get Tables
    const {trigger: requestPage} =
        useSWRMutation(`../api/v1/movies?pageIndex=${pageIndex}&pageSize=${pageSize}`, post);
    const {trigger: deleteMovie} = useSWRMutation(`../api/v1/admin/movies/delete/${deleteId}`, del);

    useEffect(() => {
        if (status === "loading") {
            return;
        }
        const role = session?.user.role;

        if (role === "admin" || role === "moderator") {
            return;
        }
        router.push("/");
    }, [router, session, status])

    useEffect(() => {
        handeRequestPage();
    }, [pageIndex, pageSize, order, orderBy])

    // Ensure the page index has been reset when the page size changes
    useEffect(() => {
        setPageIndex(0);
    }, [pageSize])

    useEffect(() => {
        if (deleteId && isConfirmDelete) {
            deleteMovie()
                .then((data) => {
                    if (data) {
                        setNotifyState({
                            open: true,
                            message: data.message,
                            vertical: "top",
                            horizontal: "right",
                            severity: "info"
                        });

                        setIsConfirmDelete(false);
                        setDeleteId(null);

                        handeRequestPage();
                    }
                })
                .catch((error) => {
                    setNotifyState({
                        open: true,
                        message: error.message,
                        vertical: "top",
                        horizontal: "right",
                        severity: "error"
                    });
                });
        }
    }, [deleteId]);

    const handeRequestPage = () => {
        requestPage({
            sort: {
                orders: [
                    {
                        property: orderBy,
                        direction: order
                    }
                ]
            }
        }).then((data) => {
            setPage(data);
        }).catch((error) => {
            setNotifyState({
                open: true,
                message: error.message,
                vertical: "top",
                horizontal: "right",
                severity: "error"
            });
        });
    }

    return (
        <>
            <NotifySnackbar state={notifyState} setState={setNotifyState}/>
            <Stack spacing={2}>
                <Box sx={{display: "flex", p: 1, m: 1}}>
                    <Typography variant="h4">Manage Catalogue</Typography>
                </Box>
                <Divider/>
                <Box>
                    <Button variant="contained" href="/admin/movies">Add Movie</Button>
                </Box>

                {!page &&
                    <>
                        <Skeleton/>
                        <Skeleton animation="wave"/>
                        <Skeleton animation={false}/>
                    </>
                }

                {page && page.data &&
                    <EnhancedTable
                        page={page}
                        setDeleteId={setDeleteId}
                        confirmDelete={isConfirmDelete}
                        setConfirmDelete={setIsConfirmDelete}
                        pageIndex={pageIndex}
                        setPageIndex={setPageIndex}
                        rowsPerPage={pageSize}
                        setRowsPerPage={setPageSize}
                        order={order}
                        setOrder={setOrder}
                        orderBy={orderBy}
                        setOrderBy={setOrderBy}
                    />
                }
            </Stack>
        </>
    )
}

export default ManageCatalogue;