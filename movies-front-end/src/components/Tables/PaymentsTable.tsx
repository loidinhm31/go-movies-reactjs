import React, { useCallback, useEffect, useState } from "react";
import {
  Box,
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
import { CustomPaymentType } from "@/types/movies";
import { Direction, PageType } from "@/types/page";
import { format } from "date-fns";
import useSWRMutation from "swr/mutation";
import { post } from "@/libs/api";
import SearchKey from "@/components/Search/SearchKey";

interface PaymentData {
  type_code: string;
  movie_title: string;
  season_name: string;
  epsisode_name: string;
  provider: string;
  amount: number;
  currency: string;
  payment_method: string;
  status: string;
  created_at: Date;
}

interface HeadCell {
  disablePadding: boolean;
  id: keyof PaymentData;
  label: string;
  numeric: boolean;
}

const headCells: readonly HeadCell[] = [
  {
    id: "type_code",
    numeric: false,
    disablePadding: false,
    label: "Type",
  },
  {
    id: "movie_title",
    numeric: false,
    disablePadding: false,
    label: "Title",
  },
  {
    id: "season_name",
    numeric: false,
    disablePadding: false,
    label: "Season",
  },
  {
    id: "epsisode_name",
    numeric: false,
    disablePadding: false,
    label: "Episode",
  },
  {
    id: "amount",
    numeric: true,
    disablePadding: true,
    label: "Amount",
  },
  {
    id: "currency",
    numeric: false,
    disablePadding: false,
    label: "Currency",
  },
  {
    id: "payment_method",
    numeric: false,
    disablePadding: false,
    label: "Payment Method",
  },
  {
    id: "status",
    numeric: false,
    disablePadding: false,
    label: "Status",
  },
  {
    id: "created_at",
    numeric: false,
    disablePadding: false,
    label: "Created Date",
  },
];

interface PaymentTableHeadProps {
  onRequestSort: (event: React.MouseEvent<unknown>, newOrderBy: keyof PaymentData) => void;
  order: Direction;
  orderBy: string;
}

function PaymentTableHead(props: PaymentTableHeadProps) {
  const { order, orderBy, onRequestSort } = props;
  const createSortHandler = (newOrderBy: keyof PaymentData) => (event: React.MouseEvent<unknown>) => {
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

export default function PaymentsTable() {
  const [page, setPage] = useState<PageType<CustomPaymentType> | null>(null);

  const [pageIndex, setPageIndex] = useState(0);
  const [pageSize, setPageSize] = useState(5);
  const [order, setOrder] = useState<Direction>(Direction.ASC);
  const [orderBy, setOrderBy] = useState<keyof PaymentData>("created_at");

  const [searchKey, setSearchKey] = useState<string>("");

  // Get Tables
  const { trigger: requestPage } = useSWRMutation(
    `/api/v1/payments?q=${searchKey}&pageIndex=${pageIndex}&pageSize=${pageSize}`,
    post
  );

  useEffect(() => {
    handeRequestPage();
  }, [pageIndex, pageSize, order, orderBy, searchKey]);

  // Ensure the page index has been reset when the page size changes
  useEffect(() => {
    setPageIndex(0);
  }, [pageSize, searchKey]);

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
        console.log(error);
      });
  };

  const handleRequestSort = useCallback(
    (event: React.MouseEvent<unknown>, newOrderBy: keyof PaymentData) => {
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

  return (
    <>
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
                <PaymentTableHead order={order!} orderBy={orderBy!} onRequestSort={handleRequestSort} />
                <TableBody>
                  {page
                    ? page.content?.map((row, index) => {
                        return (
                          <TableRow hover role="checkbox" tabIndex={-1} key={row.id} sx={{ cursor: "pointer" }}>
                            <TableCell>{row.type_code}</TableCell>
                            <TableCell>{row.movie_title}</TableCell>
                            <TableCell>{row.season_name}</TableCell>
                            <TableCell>{row.episode_name}</TableCell>
                            <TableCell align="right">{row.amount}</TableCell>
                            <TableCell>{row.currency.toUpperCase()}</TableCell>
                            <TableCell>{row.payment_method}</TableCell>
                            <TableCell>{row.status}</TableCell>
                            <TableCell>{format(new Date(row.created_at!), "yyyy-MM-dd")}</TableCell>
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
