import {useCallback, useEffect, useState} from "react";
import {
    Box,
    Button,
    Chip,
    Paper,
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
import {MovieType} from "../../../types/movies";
import moment from "moment";
import {boolean} from "boolean";
import AlertDialog from "../../shared/alert";

export interface Data {
    movie: string;
    release_date: Date;
    rating: number;
}

type Order = "asc" | "desc";

interface HeadCell {
    disablePadding: boolean;
    id: keyof Data;
    label: string;
    numeric: boolean;
}

const headCells: readonly HeadCell[] = [
    {
        id: "movie",
        numeric: false,
        disablePadding: false,
        label: "Movie",
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
    order: Order;
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


interface EnhancedTableProps {
    rows: MovieType[];
    setDeleteId: (selectId: number) => void;
    confirmDelete: boolean;
    setConfirmDelete: (flag: boolean) => void;
}

export default function EnhancedTable(props: EnhancedTableProps) {
    const [order, setOrder] = useState<Order>("asc");
    const [orderBy, setOrderBy] = useState<keyof Data>("release_date");
    const [page, setPage] = useState(0);
    const [rowsPerPage, setRowsPerPage] = useState(5);

    const [isOpenDeleteDialog, setIsOpenDeleteDialog] = useState(false);
    const [selectedTitle, setSelectedTitle] = useState("");
    const [selectedId, setSelectedId] = useState<number>();

    useEffect(() => {
        if (props.confirmDelete) {
            props.setDeleteId(selectedId);
            props.setConfirmDelete(true);
        }
    }, [props.confirmDelete])

    const handleRequestSort = useCallback(
        (event: React.MouseEvent<unknown>, newOrderBy: keyof Data) => {
            const isAsc = orderBy === newOrderBy && order === "asc";
            const toggledOrder = isAsc ? "desc" : "asc";
            setOrder(toggledOrder);
            setOrderBy(newOrderBy);
        },
        [order, orderBy, page, rowsPerPage],
    );

    const handleChangePage = useCallback(
        (event: unknown, newPage: number) => {
            setPage(newPage);
            // Avoid a layout jump when reaching the last page with empty rows.
            const numEmptyRows =
                newPage > 0 ? Math.max(0, (1 + newPage) * rowsPerPage - props.rows.length) : 0;

        },
        [order, orderBy, rowsPerPage],
    );

    const handleChangeRowsPerPage = useCallback(
        (event: React.ChangeEvent<HTMLInputElement>) => {
            const updatedRowsPerPage = parseInt(event.target.value, 10);
            setRowsPerPage(updatedRowsPerPage);

            setPage(0);
        },
        [order, orderBy],
    );

    const handleDeleteRow = (movie: MovieType) => {
        setIsOpenDeleteDialog(true);
        setSelectedId(movie.id);
        setSelectedTitle(movie.title);
    };

    return (
        <>
            {isOpenDeleteDialog &&
                <AlertDialog
                    open={isOpenDeleteDialog}
                    setOpen={setIsOpenDeleteDialog}
                    title={`Delete "${selectedTitle}"`}
                    description={"You cannot undo this action!"}
                    confirmText={"Yes"}
                    showCancelButton={true}
                    setConfirmDelete={props.setConfirmDelete}/>
            }

            <Box sx={{width: "100%"}}>
                <Paper sx={{width: "100%", mb: 2}}>
                    <TableContainer>
                        <Table
                            sx={{minWidth: 750}}
                            aria-labelledby="tableTitle"
                        >
                            <EnhancedTableHead
                                order={order}
                                orderBy={orderBy}
                                onRequestSort={handleRequestSort}
                            />
                            <TableBody>
                                {props.rows
                                    ? props.rows.map((row, index) => {
                                        const labelId = `enhanced-table-checkbox-${index}`;

                                        return (
                                            <TableRow
                                                hover
                                                role="checkbox"
                                                tabIndex={-1}
                                                key={row.id}
                                                sx={{cursor: "pointer"}}
                                            >

                                                <TableCell
                                                    id={labelId}

                                                >
                                                    <Chip label={row.title} color="info" component="a"
                                                          href={`/movies/${row.id}`} clickable/>
                                                </TableCell>
                                                <TableCell
                                                    id={labelId}
                                                >
                                                    {moment(row.release_date).format("yyyy-MM-DD")}
                                                </TableCell>
                                                <TableCell
                                                    id={labelId}
                                                >
                                                    {row.mpaa_rating}
                                                </TableCell>
                                                <TableCell>
                                                    <Box sx={{display: "flex", gap: 1}}>
                                                        <Button variant="contained" color="inherit"
                                                                href={`/admin/movies?id=${row.id}`}>
                                                            Edit
                                                        </Button>
                                                        <Button variant="contained" color="error"
                                                                onClick={() => handleDeleteRow(row)}>
                                                            Delete
                                                        </Button>
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
                        rowsPerPageOptions={[5, 10, 25]}
                        component="div"
                        count={props.rows.length}
                        rowsPerPage={rowsPerPage}
                        page={page}
                        onPageChange={handleChangePage}
                        onRowsPerPageChange={handleChangeRowsPerPage}
                    />
                </Paper>
            </Box>
        </>

    );
}