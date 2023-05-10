import {useCallback} from "react";
import {
    Box,
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
import {MovieType} from "../../types/movies";
import {Direction, PageType} from "../../types/page";
import format from "date-fns/format";

export interface Data {
    title: string;
    release_date: Date;
    runtime: number
    description: string;
    mpaa_rating: string;
}

export interface HeadCell {
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
        id: "release_date",
        numeric: false,
        disablePadding: false,
        label: "Release Date",
    },
    {
        id: "runtime",
        numeric: true,
        disablePadding: false,
        label: "Runtime",
    },
    {
        id: "description",
        numeric: false,
        disablePadding: false,
        label: "Description",
    },
    {
        id: "mpaa_rating",
        numeric: false,
        disablePadding: false,
        label: "MPAA Rating",
    },
];

interface SearchTableHeadProps {
    onRequestSort: (event: React.MouseEvent<unknown>, newOrderBy: keyof Data) => void;
    order: Direction;
    orderBy: string;
}

function SearchTableHead(props: SearchTableHeadProps) {
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
            </TableRow>
        </TableHead>
    );
}


interface SearchTableProps {
    page: PageType<MovieType>;
    pageIndex: number;
    setPageIndex: (value: number) => void;
    rowsPerPage: number;
    setRowsPerPage: (value: number) => void;
    order?: Direction;
    setOrder?: (direction: Direction) => void;
    orderBy?: keyof Data;
    setOrderBy?: (by: keyof Data) => void;
}

export default function SearchTable({
                                        order,
                                        orderBy,
                                        page,
                                        pageIndex,
                                        rowsPerPage,
                                        setOrder,
                                        setOrderBy,
                                        setPageIndex,
                                        setRowsPerPage,
                                    }: SearchTableProps) {


    const handleRequestSort = useCallback(
        (event: React.MouseEvent<unknown>, newOrderBy: keyof Data) => {
            const isAsc = orderBy === newOrderBy && order === "asc";
            const toggledOrder = isAsc ? Direction.DESC : Direction.ASC;
            setOrder?.(toggledOrder);
            setOrderBy?.(newOrderBy);
        },
        [order, orderBy, pageIndex, rowsPerPage],
    );

    const handleChangePageIndex = useCallback(
        (event: unknown, newPageIndex: number) => {
            setPageIndex(newPageIndex);
        },
        [order, orderBy, rowsPerPage],
    );

    const handleChangeRowsPerPage = useCallback(
        (event: React.ChangeEvent<HTMLInputElement>) => {
            const updatedRowsPerPage = parseInt(event.target.value, 10);
            setRowsPerPage(updatedRowsPerPage);
            setPageIndex(0);
        },
        [order, orderBy],
    );

    return (
        <>
            <Box sx={{width: "100%"}}>
                <Paper sx={{width: "100%", mb: 2}}>
                    <TableContainer>
                        <Table
                            sx={{minWidth: 750}}
                            aria-labelledby="tableTitle"
                        >
                            <SearchTableHead
                                order={order!}
                                orderBy={orderBy!}
                                onRequestSort={handleRequestSort}
                            />
                            <TableBody>
                                {page
                                    ? page.data?.map((row, index) => {
                                        const labelId = `search-table-checkbox-${index}`;

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
                                                    {format(new Date(row.release_date!), "yyyy-MM-dd")}
                                                </TableCell>
                                                <TableCell
                                                    id={labelId}
                                                >
                                                    {row.runtime}
                                                </TableCell>
                                                <TableCell
                                                    id={labelId}
                                                >
                                                    {row.description}
                                                </TableCell>

                                                <TableCell
                                                    id={labelId}
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
                        rowsPerPageOptions={[5, 10, 25, 50]}
                        component="div"
                        count={page.total_elements!}
                        rowsPerPage={rowsPerPage}
                        page={pageIndex}
                        onPageChange={handleChangePageIndex}
                        onRowsPerPageChange={handleChangeRowsPerPage}
                    />
                </Paper>
            </Box>
        </>

    );
}