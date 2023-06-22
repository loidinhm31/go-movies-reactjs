import DeleteIcon from "@mui/icons-material/Delete";
import EditIcon from "@mui/icons-material/Edit";
import PriceChangeIcon from "@mui/icons-material/PriceChange";
import VisibilityIcon from "@mui/icons-material/Visibility";
import {
  Box,
  Chip,
  IconButton,
  Paper,
  Skeleton,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TablePagination,
  TableRow,
  TableSortLabel,
} from "@mui/material";
import { visuallyHidden } from "@mui/utils";
import { format } from "date-fns";
import React, { useCallback, useEffect, useState } from "react";
import SearchKey from "@/components/Search/SearchKey";
import AlertDialog from "@/components/shared/alert";
import { NotifyState } from "@/components/shared/snackbar";
import { del, patch, post } from "@/libs/api";
import { MovieType } from "@/types/movies";
import { Direction, PageType } from "@/types/page";
import useSWRMutation from "swr/mutation";

export interface Data {
  title: string;
  price: number;
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
    id: "price",
    numeric: true,
    disablePadding: true,
    label: "Price (USD)",
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
  const { order, orderBy, onRequestSort } = props;
  const createSortHandler = (newOrderBy: keyof Data) => (event: React.MouseEvent<unknown>) => {
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
        <TableCell aria-label="last" style={{ width: "var(--Table-lastColumnWidth)" }} />
      </TableRow>
    </TableHead>
  );
}

interface ManageMoviesTableProps {
  selectedMovie: MovieType | null;
  setSelectedMovie: (obj: MovieType | null) => void;
  setOpenSeasonDialog: (flag: boolean) => void;
  setNotifyState: (state: NotifyState) => void;
}

export default function ManageMoviesTable({
  selectedMovie,
  setSelectedMovie,
  setOpenSeasonDialog,
  setNotifyState,
}: ManageMoviesTableProps) {
  const [page, setPage] = useState<PageType<MovieType> | null>(null);

  const [isConfirmDelete, setIsConfirmDelete] = useState<boolean>(false);
  const [deleteId, setDeleteId] = useState<number | null>();

  const [pageIndex, setPageIndex] = useState(0);
  const [pageSize, setPageSize] = useState(5);
  const [order, setOrder] = useState<Direction>(Direction.ASC);
  const [orderBy, setOrderBy] = useState<keyof Data>("release_date");

  const [isOpenDeleteDialog, setIsOpenDeleteDialog] = useState(false);

  const [searchKey, setSearchKey] = useState<string>("");

  // Get Tables
  const { trigger: requestPage } = useSWRMutation(
    `/api/v1/movies?q=${searchKey}&pageIndex=${pageIndex}&pageSize=${pageSize}`,
    post
  );
  const { trigger: deleteMovie } = useSWRMutation(`/api/v1/admin/movies/delete/${deleteId}`, del);
  const { trigger: updateMoviePrice } = useSWRMutation(`/api/v1/admin/movies/price`, patch);

  useEffect(() => {
    handeRequestPage();
  }, [pageIndex, pageSize, order, orderBy, searchKey]);

  // Ensure the page index has been reset when the page size changes
  useEffect(() => {
    setPageIndex(0);
  }, [pageSize, searchKey]);

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
              severity: "info",
            });

            handeRequestPage();
          }
        })
        .catch((error) => {
          setNotifyState({
            open: true,
            message: error.message.message,
            vertical: "top",
            horizontal: "right",
            severity: "error",
          });
        })
        .finally(() => {
          setSelectedMovie(null);
          setIsConfirmDelete(false);
          setDeleteId(null);
        });
    }
  }, [deleteId, isConfirmDelete]);

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
            direction: order,
          },
        ],
      },
    })
      .then((data) => {
        setPage(data);
      })
      .catch((error) => {
        setNotifyState({
          open: true,
          message: error.message.message,
          vertical: "top",
          horizontal: "right",
          severity: "error",
        });
      });
  };

  const handleRequestSort = useCallback(
    (event: React.MouseEvent<unknown>, newOrderBy: keyof Data) => {
      const isAsc = orderBy === newOrderBy && order === "asc";
      const toggledOrder = isAsc ? Direction.DESC : Direction.ASC;
      setOrder?.(toggledOrder);
      setOrderBy?.(newOrderBy);
    },
    [order, orderBy, pageIndex, pageSize]
  );

  const handleChangePageIndex = useCallback(
    (event: unknown, newPageIndex: number) => {
      setPageIndex(newPageIndex);
    },
    [order, orderBy, pageSize]
  );

  const handleChangeRowsPerPage = useCallback(
    (event: React.ChangeEvent<HTMLInputElement>) => {
      const updatedRowsPerPage = parseInt(event.target.value, 10);
      setPageSize(updatedRowsPerPage);
      setPageIndex(0);
    },
    [order, orderBy]
  );

  const handleDeleteRow = (movie: MovieType) => {
    setIsOpenDeleteDialog(true);
    setSelectedMovie(movie);
  };

  const handleViewTv = (movie: MovieType) => {
    setOpenSeasonDialog(true);
    setSelectedMovie(movie);
  };

  const handleUpdateAveragePrice = (movie: MovieType) => {
    updateMoviePrice({ id: movie.id } as MovieType)
      .then((result) => {
        if (result.message === "ok") {
          setNotifyState({
            open: true,
            message: "Average Price Was Updated",
            vertical: "top",
            horizontal: "right",
            severity: "success",
          });
        }
      })
      .catch((error) => {
        setNotifyState({
          open: true,
          message: error.message.message,
          vertical: "top",
          horizontal: "right",
          severity: "error",
        });
      });
  };

  return (
    <>
      {isOpenDeleteDialog && (
        <AlertDialog
          open={isOpenDeleteDialog}
          setOpen={setIsOpenDeleteDialog}
          title={`Delete "${selectedMovie?.title}"`}
          description={"You cannot undo this action!"}
          confirmText={"Yes"}
          showCancelButton={true}
          setConfirmDelete={setIsConfirmDelete}
        />
      )}

      {!page && (
        <>
          <Skeleton />
          <Skeleton animation="wave" />
          <Skeleton animation={false} />
        </>
      )}
      <SearchKey keyword={searchKey} setKeyword={setSearchKey} />
      {page && page.content && (
        <Box sx={{ width: "100%" }}>
          <Paper sx={{ width: "100%", mb: 2 }}>
            <TableContainer>
              <Table sx={{ minWidth: 750 }} aria-labelledby="tableTitle">
                <EnhancedTableHead order={order!} orderBy={orderBy!} onRequestSort={handleRequestSort} />
                <TableBody>
                  {page
                    ? page.content?.map((row, index) => {
                        return (
                          <TableRow hover role="checkbox" tabIndex={-1} key={row.id} sx={{ cursor: "pointer" }}>
                            <TableCell>
                              <Chip label={row.title} color="info" component="a" href={`/movies/${row.id}`} clickable />
                            </TableCell>
                            <TableCell>{row.type_code}</TableCell>
                            <TableCell align="right">
                              {`${row.type_code === "MOVIE" ? (row.price ? row.price : "FREE") : ""}`}
                            </TableCell>
                            <TableCell>{format(new Date(row.release_date!), "yyyy-MM-dd")}</TableCell>
                            <TableCell>{row.mpaa_rating}</TableCell>
                            <TableCell>
                              <Box sx={{ display: "flex", gap: 1 }}>
                                <IconButton color="inherit" href={`/admin/manage-catalogue/movies?id=${row.id}`}>
                                  <EditIcon />
                                </IconButton>
                                <IconButton color="error" onClick={() => handleDeleteRow(row)}>
                                  <DeleteIcon />
                                </IconButton>
                                {row.type_code === "TV" && (
                                  <>
                                    <IconButton color="inherit" onClick={() => handleViewTv(row)}>
                                      <VisibilityIcon />
                                    </IconButton>
                                    <IconButton color="secondary" onClick={() => handleUpdateAveragePrice(row)}>
                                      <PriceChangeIcon />
                                    </IconButton>
                                  </>
                                )}
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
      )}
    </>
  );
}
