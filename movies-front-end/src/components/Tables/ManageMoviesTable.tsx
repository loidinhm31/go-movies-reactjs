import React, {useCallback, useEffect, useState} from "react";
import {
    Box,
    Button,
    Chip, IconButton,
    Paper, Skeleton,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TablePagination,
    TableRow,
    TableSortLabel
} from "@mui/material";
import {visuallyHidden} from "@mui/utils";
import {MovieType} from "../../types/movies";
import AlertDialog from "../shared/alert";
import {Direction, PageType} from "../../types/page";
import {format} from "date-fns";
import EditIcon from "@mui/icons-material/Edit";
import DeleteIcon from "@mui/icons-material/Delete";
import VisibilityIcon from "@mui/icons-material/Visibility";
import SeasonDialog from "../Dialog/SeasonDialog";
import useSWRMutation from "swr/mutation";
import {del, post} from "../../libs/api";
import {NotifyState} from "../shared/snackbar";

export interface Data {
    title: string;
    type_code: string;
    release_date: Date;
    rating: number;
}

interface HeadCell {
    disablePadding: boolean;
    id: keyof Data;
    label: string;
    numeric: boolean;
}

const headCells: readonly HeadCell[] = [
    {
        id: "title",
        numeric: false,
        disablePadding: false,
        label: "Movie",
    },
    {
        id: "type_code",
        numeric: false,
        disablePadding: false,
        label: "Type",
    },
    {
        id: "release_date",
        numeric: false,
        disablePadding: false,
        label: "Release Date",
    },
    {
        id: "rating",
        numeric: false,
        disablePadding: false,
        label: "Rating",
    },
];

interface EnhancedTableHeadProps {
    onRequestSort: (event: React.MouseEvent<unknown>, newOrderBy: keyof Data) => void;
    order: Direction;
    orderBy: string;
}

function EnhancedTableHead(props: EnhancedTableHeadProps) {
    const {order, orderBy, onRequestSort} = props;
    const createSortHandler =
        (newOrderBy: keyof Data) => (event: React.MouseEvent<unknown>) => {
            onRequestSort(event, newOrderBy);
        };

    return (
        <TableHead>
            <TableRow>
                {headCells.map((headCell) => (
                    <TableCell
                        key={headCell.id}
                        align={headCell.numeric ? "right" : "left"}
                        padding={headCell.disablePadding ? "none" : "normal"}
                        sortDirection={orderBy === headCell.id ? order : false}
                    >
                        <TableSortLabel
                            active={orderBy === headCell.id}
                            direction={orderBy === headCell.id ? order : "asc"}
                            onClick={createSortHandler(headCell.id)}
                        >
                            <b>{headCell.label}</b>
                            {orderBy === headCell.id ? (
                                <Box component="span" sx={visuallyHidden}>
                                    {order === "desc" ? "sorted descending" : "sorted ascending"}
                                </Box>
                            ) : null}
                        </TableSortLabel>
                    </TableCell>
                ))}
                <TableCell
                    aria-label="last"
                    style={{width: "var(--Table-lastColumnWidth)"}}/>
            </TableRow>
        </TableHead>
    );
}


interface ManageMoviesTableProps {
    selectedMovie: MovieType | null;
    setSelectedMovie: (obj: MovieType | null) => void
    setOpenSeasonDialog: (flag: boolean) => void;
    setNotifyState: (state: NotifyState) => void;
}

export default function ManageMoviesTable({selectedMovie, setSelectedMovie, setOpenSeasonDialog, setNotifyState}: ManageMoviesTableProps) {
    const [page, setPage] = useState<PageType<MovieType> | null>(null);

    const [isConfirmDelete, setIsConfirmDelete] = useState<boolean>(false);
    const [deleteId, setDeleteId] = useState<number | null>();

    const [pageIndex, setPageIndex] = useState(0);
    const [pageSize, setPageSize] = useState(5)
    const [order, setOrder] = useState<Direction>(Direction.ASC);
    const [orderBy, setOrderBy] = useState<keyof Data>("release_date");

    const [isOpenDeleteDialog, setIsOpenDeleteDialog] = useState(false);

    // Get Tables
    const {trigger: requestPage} =
        useSWRMutation(`../api/v1/movies?pageIndex=${pageIndex}&pageSize=${pageSize}`, post);
    const {trigger: deleteMovie} = useSWRMutation(`../api/v1/admin/movies/delete/${deleteId}`, del);

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
                    if (data.message === "ok") {
                        setNotifyState({
                            open: true,
                            message: "Movie deleted",
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
                        message: error.message.message,
                        vertical: "top",
                        horizontal: "right",
                        severity: "error"
                    });
                })
                .finally(() => {
                    setSelectedMovie(null);
                });
        }
    }, [deleteId]);

    useEffect(() => {
        if (isConfirmDelete) {
            setDeleteId(selectedMovie?.id);
            setIsConfirmDelete(true);
        }
    }, [isConfirmDelete]);

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
                message: error.message.message,
                vertical: "top",
                horizontal: "right",
                severity: "error"
            });
        });
    }

    const handleRequestSort = useCallback(
        (event: React.MouseEvent<unknown>, newOrderBy: keyof Data) => {
            const isAsc = orderBy === newOrderBy && order === "asc";
            const toggledOrder = isAsc ? Direction.DESC : Direction.ASC;
            setOrder?.(toggledOrder);
            setOrderBy?.(newOrderBy);
        },
        [order, orderBy, pageIndex, pageSize],
    );

    const handleChangePageIndex = useCallback(
        (event: unknown, newPageIndex: number) => {
            setPageIndex(newPageIndex);
        },
        [order, orderBy, pageSize],
    );

    const handleChangeRowsPerPage = useCallback(
        (event: React.ChangeEvent<HTMLInputElement>) => {
            const updatedRowsPerPage = parseInt(event.target.value, 10);
            setPageSize(updatedRowsPerPage);
            setPageIndex(0);
        },
        [order, orderBy],
    );

    const handleDeleteRow = (movie: MovieType) => {
        setIsOpenDeleteDialog(true);
        setSelectedMovie(movie);
    };

    const handleViewTv = (movie: MovieType) => {
        setOpenSeasonDialog(true);
        setSelectedMovie(movie);
    }

    return (
        <>
            {isOpenDeleteDialog &&
                <AlertDialog
                    open={isOpenDeleteDialog}
                    setOpen={setIsOpenDeleteDialog}
                    title={`Delete "${selectedMovie?.title}"`}
                    description={"You cannot undo this action!"}
                    confirmText={"Yes"}
                    showCancelButton={true}
                    setConfirmDelete={setIsConfirmDelete}/>
            }

            {!page &&
                <>
                    <Skeleton/>
                    <Skeleton animation="wave"/>
                    <Skeleton animation={false}/>
                </>
            }

            {page && page.content &&
                <Box sx={{width: "100%"}}>
                    <Paper sx={{width: "100%", mb: 2}}>
                        <TableContainer>
                            <Table
                                sx={{minWidth: 750}}
                                aria-labelledby="tableTitle"
                            >
                                <EnhancedTableHead
                                    order={order!}
                                    orderBy={orderBy!}
                                    onRequestSort={handleRequestSort}
                                />
                                <TableBody>
                                    {page
                                        ? page.content?.map((row, index) => {
                                            return (
                                                <TableRow
                                                    hover
                                                    role="checkbox"
                                                    tabIndex={-1}
                                                    key={row.id}
                                                    sx={{cursor: "pointer"}}
                                                >
                                                    <TableCell>
                                                        <Chip label={row.title} color="info" component="a"
                                                              href={`/movies/${row.id}`} clickable/>
                                                    </TableCell>
                                                    <TableCell>
                                                        {row.type_code}
                                                    </TableCell>
                                                    <TableCell>
                                                        {format(new Date(row.release_date!), "yyyy-MM-dd")}
                                                    </TableCell>
                                                    <TableCell>
                                                        {row.mpaa_rating}
                                                    </TableCell>
                                                    <TableCell>
                                                        <Box sx={{display: "flex", gap: 1}}>
                                                            <IconButton
                                                                color="inherit"
                                                                href={`/admin/movies?id=${row.id}`}
                                                            >
                                                                <EditIcon/>
                                                            </IconButton>
                                                            <IconButton
                                                                color="error"
                                                                onClick={() => handleDeleteRow(row)}
                                                            >
                                                                <DeleteIcon/>
                                                            </IconButton>
                                                            {row.type_code === "TV" &&
                                                                <IconButton
                                                                    color="inherit"
                                                                    onClick={() => handleViewTv(row)}
                                                                >
                                                                    <VisibilityIcon/>
                                                                </IconButton>
                                                            }
                                                        </Box>
                                                    </TableCell>
                                                </TableRow>
                                            );
                                        })
                                        : null}
                                </TableBody>
                            </Table>
                        </TableContainer>
                        <TablePagination
                            rowsPerPageOptions={[5, 10, 25, 50]}
                            component="div"
                            count={page.total_elements!}
                            rowsPerPage={pageSize}
                            page={pageIndex}
                            onPageChange={handleChangePageIndex}
                            onRowsPerPageChange={handleChangeRowsPerPage}
                        />
                    </Paper>
                </Box>
            }
            
        </>

    );
}