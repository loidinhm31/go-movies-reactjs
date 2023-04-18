import {useCallback, useState} from "react";
import {
    Box,
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
import Link from "next/link";

export interface Data {
    movie: string;
    release_date: Date;
    rating: number;
}

type Order = 'asc' | 'desc';

interface HeadCell {
    disablePadding: boolean;
    id: keyof Data;
    label: string;
    numeric: boolean;
}

const headCells: readonly HeadCell[] = [
    {
        id: 'movie',
        numeric: false,
        disablePadding: false,
        label: 'Movie',
    },
    {
        id: 'release_date',
        numeric: false,
        disablePadding: false,
        label: 'Release Date',
    },
    {
        id: 'rating',
        numeric: false,
        disablePadding: false,
        label: 'Rating',
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
        <TableHead component="div">
            <TableRow component="div">
                {headCells.map((headCell) => (
                    <TableCell component="div"
                        key={headCell.id}
                        align={headCell.numeric ? 'right' : 'left'}
                        padding={headCell.disablePadding ? 'none' : 'normal'}
                        sortDirection={orderBy === headCell.id ? order : false}
                    >
                        <TableSortLabel
                            active={orderBy === headCell.id}
                            direction={orderBy === headCell.id ? order : 'asc'}
                            onClick={createSortHandler(headCell.id)}
                        >
                            <b>{headCell.label}</b>
                            {orderBy === headCell.id ? (
                                <Box component="span" sx={visuallyHidden}>
                                    {order === 'desc' ? 'sorted descending' : 'sorted ascending'}
                                </Box>
                            ) : null}
                        </TableSortLabel>
                    </TableCell>
                ))}
            </TableRow>
        </TableHead>
    );
}


interface EnhancedTableProps {
    rows: MovieType[];
}

export default function EnhancedTable(props: EnhancedTableProps) {
    const [order, setOrder] = useState<Order>("asc");
    const [orderBy, setOrderBy] = useState<keyof Data>("release_date");
    const [page, setPage] = useState(0);
    const [rowsPerPage, setRowsPerPage] = useState(5);

    const handleRequestSort = useCallback(
        (event: React.MouseEvent<unknown>, newOrderBy: keyof Data) => {
            const isAsc = orderBy === newOrderBy && order === 'asc';
            const toggledOrder = isAsc ? 'desc' : 'asc';
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

    return (
        <Box sx={{width: '100%'}}>
            <Paper sx={{width: '100%', mb: 2}}>
                <TableContainer>
                    <Table component="div"
                        sx={{minWidth: 750}}
                        aria-labelledby="tableTitle"
                    >
                        <EnhancedTableHead
                            order={order}
                            orderBy={orderBy}
                            onRequestSort={handleRequestSort}
                        />
                        <TableBody component="div">
                            {props.rows
                                ? props.rows.map((row, index) => {
                                    const labelId = `enhanced-table-checkbox-${index}`;

                                    return (
                                        <TableRow
                                            hover
                                            role="checkbox"
                                            tabIndex={-1}
                                            key={row.id}
                                            sx={{cursor: 'pointer'}}
                                            component={Link}
                                            href={`/movies/${row.id}`}
                                            style={{textDecoration: 'none'}}
                                        >
                                            <TableCell
                                                component="div"
                                                id={labelId}
                                                scope="row"
                                            >
                                                {row.title}
                                            </TableCell>
                                            <TableCell
                                                component="div"
                                                id={labelId}
                                                scope="row"
                                            >
                                                {new Date(row.release_date).toDateString()}
                                            </TableCell>
                                            <TableCell
                                                component="div"
                                                id={labelId}
                                                scope="row"
                                            >
                                                {row.mpaa_rating}
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
    );
}