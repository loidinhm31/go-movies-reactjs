import {useCallback, useEffect, useState} from "react";
import {
    Box,
    Button,
    Chip, IconButton,
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
import AlertDialog from "../shared/alert";
import {Direction, PageType} from "../../types/page";
import {format} from "date-fns";
import EditIcon from "@mui/icons-material/Edit";
import DeleteIcon from "@mui/icons-material/Delete";

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


interface EnhancedTableProps {
    page: PageType<MovieType>;
    setDeleteId?: (selectId: number | null | undefined) => void;
    confirmDelete?: boolean;
    setConfirmDelete?: (flag: boolean) => void;
    pageIndex: number;
    setPageIndex: (value: number) => void;
    rowsPerPage: number;
    setRowsPerPage: (value: number) => void;
    order?: Direction;
    setOrder?: (direction: Direction) => void;
    orderBy?: keyof Data;
    setOrderBy?: (by: keyof Data) => void;
}

export default function ManageMoviesTable({
                                              confirmDelete,
                                              order,
                                              orderBy,
                                              page,
                                              pageIndex,
                                              rowsPerPage,
                                              setConfirmDelete,
                                              setDeleteId,
                                              setOrder,
                                              setOrderBy,
                                              setPageIndex,
                                              setRowsPerPage
                                          }: EnhancedTableProps) {
    const [isOpenDeleteDialog, setIsOpenDeleteDialog] = useState(false);
    const [selectedTitle, setSelectedTitle] = useState("");
    const [selectedId, setSelectedId] = useState<number | null>();

    useEffect(() => {
        if (confirmDelete) {
            setDeleteId?.(selectedId);
            setConfirmDelete?.(true);
        }
    }, [confirmDelete])

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
                    setConfirmDelete={setConfirmDelete}/>
            }

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