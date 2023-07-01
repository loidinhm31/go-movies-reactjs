import { Checkbox, Grid, Stack, TextField, Typography } from "@mui/material";
import Button from "@mui/material/Button";
import UserTable, { UserData } from "@/components/Tables/UserTable";
import { useEffect, useState } from "react";
import { Direction, PageType } from "@/types/page";
import { UserType } from "@/types/users";
import useSWRMutation from "swr/mutation";
import { post } from "@/libs/api";
import { NotifyState } from "@/components/shared/snackbar";

interface SearchUsersProps {
  setNotifyState: (state: NotifyState) => void;
  wasUpdated: boolean;
  setWasUpdated: (flag: boolean) => void;
}

export default function SearchUsers({ setNotifyState, wasUpdated, setWasUpdated }: SearchUsersProps) {
  const [isNew, setIsNew] = useState(true);
  const [searchKey, setSearchKey] = useState("");

  const [page, setPage] = useState<PageType<UserType> | null>(null);

  const [pageIndex, setPageIndex] = useState(0);
  const [pageSize, setPageSize] = useState(5);
  const [order, setOrder] = useState<Direction>(Direction.ASC);
  const [orderBy, setOrderBy] = useState<keyof UserData>("created_at");

  const { trigger: requestPage } = useSWRMutation(
    `/api/v1/admin/users?pageSize=${pageSize}&pageIndex=${pageIndex}&isNew=${isNew}&query=${searchKey}`,
    post
  );

  useEffect(() => {
    handeRequestPage();
  }, [pageIndex, pageSize, order, orderBy, isNew]);

  useEffect(() => {
    if (wasUpdated) {
      handeRequestPage();
      setWasUpdated(false);
    }
  }, [wasUpdated]);

  // Ensure the page index has been reset when the page size changes
  useEffect(() => {
    setPageIndex(0);
  }, [pageSize]);
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

  const handleSearchClick = () => {
    handeRequestPage();
  };

  const handleKeyPressSearch = (event) => {
    if (event.key === "Enter") {
      handleSearchClick();
    }
  };

  return (
    <Grid container spacing={2} sx={{ display: "flex", alignItems: "center" }}>
      <Grid item xs={10} sx={{ p: 2 }}>
        <TextField
          fullWidth
          label="Keyword"
          variant="outlined"
          value={searchKey}
          onChange={(e) => setSearchKey(e.target.value)}
          onKeyDown={handleKeyPressSearch}
        />
      </Grid>
      <Grid item xs={2}>
        <Stack sx={{ display: "flex", alignItems: "center" }}>
          <Typography>Is New?</Typography>
          <Checkbox data-testid="is-new" checked={isNew} onChange={(event) => setIsNew(event.target.checked)} />
        </Stack>
      </Grid>
      <Grid item xs={12}>
        <Button variant="contained" onClick={handleSearchClick}>
          Search
        </Button>
      </Grid>

      <Grid item xs="auto">
        {page && (
          <UserTable
            page={page}
            pageIndex={pageIndex}
            setPageIndex={setPageIndex}
            rowsPerPage={pageSize}
            setRowsPerPage={setPageSize}
            order={order}
            setOrder={setOrder}
            orderBy={orderBy}
            setOrderBy={setOrderBy}
            setNotifyState={setNotifyState}
            setWasUpdated={setWasUpdated}
          />
        )}
      </Grid>
    </Grid>
  );
}
