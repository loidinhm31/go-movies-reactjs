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
    TableSortLabel,
} from "@mui/material";
import { format } from "date-fns";
import { Direction, PageType } from "src/types/page";
import { visuallyHidden } from "@mui/utils";
import { useCallback, useEffect, useState } from "react";
import { UserType } from "src/types/users";
import RoleDialog from "src/components/Dialog/RoleDialog";
import { NotifyState } from "src/components/shared/snackbar";

export interface UserData {
    username: string;
    first_name: string;
    last_name: string;
    email: string;
    is_new: boolean;
    role: string;
    created_at: string;
}

interface HeadCell {
    disablePadding: boolean;
    id: keyof UserData;
    label: string;
}

const headCells: readonly HeadCell[] = [
    {
        id: "username",
        disablePadding: false,
        label: "Username",
    },
    {
        id: "first_name",
        disablePadding: false,
        label: "First Name",
    },
    {
        id: "last_name",
        disablePadding: false,
        label: "Last Name",
    },
    {
        id: "email",
        disablePadding: false,
        label: "Email",
    },
    {
        id: "created_at",
        disablePadding: false,
        label: "Created At",
    },
    {
        id: "is_new",
        disablePadding: false,
        label: "Is New?",
    },
    {
        id: "role",
        disablePadding: false,
        label: "Role",
    },
];

interface UserTableHeadProps {
    onRequestSort: (event: React.MouseEvent<unknown>, newOrderBy: keyof UserData) => void;
    order: Direction;
    orderBy: string;
}

function UserTableHead(props: UserTableHeadProps) {
    const { order, orderBy, onRequestSort } = props;
    const createSortHandler = (newOrderBy: keyof UserData) => (event: React.MouseEvent<unknown>) => {
        onRequestSort(event, newOrderBy);
    };

    return (
        <TableHead>
            <TableRow>
                {headCells.map((headCell) => (
                    <TableCell
                        key={headCell.id}
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

interface UserTableProps {
    page: PageType<UserType> | null;
    pageIndex: number;
    setPageIndex: (value: number) => void;
    rowsPerPage: number;
    setRowsPerPage: (value: number) => void;
    order?: Direction;
    setOrder?: (direction: Direction) => void;
    orderBy?: keyof UserData;
    setOrderBy?: (by: keyof UserData) => void;
    setNotifyState: (state: NotifyState) => void;
    setWasUpdated: (flag: boolean) => void;
}

export default function UserTable({
    order,
    orderBy,
    page,
    pageIndex,
    rowsPerPage,
    setOrder,
    setOrderBy,
    setPageIndex,
    setRowsPerPage,
    setNotifyState,
    setWasUpdated,
}: UserTableProps) {
    const [selectedUser, setSelectedUser] = useState<UserType | null>(null);
    const [isOpenDialog, setIsOpenDialog] = useState(false);

    useEffect(() => {
        if (!isOpenDialog) {
            setSelectedUser(null);
        }
    }, [isOpenDialog]);

    const handleRequestSort = useCallback(
        (event: React.MouseEvent<unknown>, newOrderBy: keyof UserData) => {
            const isAsc = orderBy === newOrderBy && order === "asc";
            const toggledOrder = isAsc ? Direction.DESC : Direction.ASC;
            setOrder?.(toggledOrder);
            setOrderBy?.(newOrderBy);
        },
        [order, orderBy, pageIndex, rowsPerPage]
    );

    const handleChangePageIndex = useCallback(
        (event: unknown, newPageIndex: number) => {
            setPageIndex(newPageIndex);
        },
        [order, orderBy, rowsPerPage]
    );

    const handleChangeRowsPerPage = useCallback(
        (event: React.ChangeEvent<HTMLInputElement>) => {
            const updatedRowsPerPage = parseInt(event.target.value, 10);
            setRowsPerPage(updatedRowsPerPage);
            setPageIndex(0);
        },
        [order, orderBy]
    );

    const handleSelectRow = (user: UserType) => {
        setIsOpenDialog(true);
        setSelectedUser(user);
    };

    return (
        <>
            {isOpenDialog && (
                <RoleDialog
                    user={selectedUser}
                    open={isOpenDialog}
                    setOpen={setIsOpenDialog}
                    setNotifyState={setNotifyState}
                    setWasUpdated={setWasUpdated}
                />
            )}

            <Box sx={{ width: "100%" }}>
                <Paper sx={{ width: "100%", mb: 2 }}>
                    <TableContainer>
                        <Table sx={{ minWidth: 750 }} aria-labelledby="tableTitle">
                            <UserTableHead order={order!} orderBy={orderBy!} onRequestSort={handleRequestSort} />
                            <TableBody>
                                {page
                                    ? page.content?.map((row, index) => {
                                          return (
                                              <TableRow
                                                  hover
                                                  role="checkbox"
                                                  tabIndex={-1}
                                                  key={`${row.id}-${index}`}
                                                  sx={{ cursor: "pointer" }}
                                              >
                                                  <TableCell>
                                                      <Chip label={row.username} color="info" />
                                                  </TableCell>
                                                  <TableCell>{row.first_name}</TableCell>
                                                  <TableCell>{row.last_name}</TableCell>
                                                  <TableCell>{row.email}</TableCell>
                                                  <TableCell>
                                                      {format(new Date(row.created_at!), "yyyy-MM-dd")}
                                                  </TableCell>
                                                  <TableCell>{`${row.is_new ? "YES" : "NO"}`}</TableCell>
                                                  <TableCell>{row.role.role_code}</TableCell>
                                                  <TableCell>
                                                      <Box sx={{ display: "flex", gap: 1 }}>
                                                          <Button
                                                              variant="contained"
                                                              color="success"
                                                              onClick={() => handleSelectRow(row)}
                                                          >
                                                              Provide Access
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
                        rowsPerPageOptions={[5, 10, 25, 50]}
                        component="div"
                        count={page!.total_elements!}
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
