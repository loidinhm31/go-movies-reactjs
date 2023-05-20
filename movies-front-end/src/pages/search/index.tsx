import {Box, Divider, Skeleton, Stack, Typography} from "@mui/material";
import {useEffect, useState} from "react";
import {SearchField} from "src/components/Search/SearchField";
import {FieldData, SearchRequest} from "src/types/search";
import useSWRMutation from "swr/mutation";
import SearchTable, {Data} from "../../components/Tables/SearchTable";
import {post} from "../../libs/api";
import {MovieType} from "../../types/movies";
import {Direction, PageType} from "../../types/page";
import NotifySnackbar, {NotifyState} from "../../components/shared/snackbar";


function Search() {
    const [page, setPage] = useState<PageType<MovieType> | null>(null);

    const [pageIndex, setPageIndex] = useState(0);
    const [pageSize, setPageSize] = useState(5)
    const [order, setOrder] = useState<Direction>(Direction.ASC);
    const [orderBy, setOrderBy] = useState<keyof Data>("release_date");

    const [fieldDataMap, setFieldDataMap] = useState<Map<string, FieldData>>(new Map());
    const searchRequest: SearchRequest = {};

    const [notifyState, setNotifyState] = useState<NotifyState>({open: false, vertical: "top", horizontal: "right"});

    // Get Tables
    const {trigger: requestPage} = useSWRMutation(`../api/v1/search`, post);

    useEffect(() => {
        searchRequest.filters = [];
        handleChangeSearchRequest(searchRequest!.filters!);
    }, [pageIndex, pageSize, order, orderBy])

    const handleChangeSearchRequest = (fieldData: FieldData[]) => {
        console.log(fieldData)
        if (fieldData) {
            searchRequest!.page_request = {
                page: pageIndex,
                size: pageSize,
                sort: {
                    orders: [
                        {
                            property: orderBy,
                            direction: order,
                        }
                    ]
                }
            };
            searchRequest!.filters = fieldData;

            requestPage(
                searchRequest
            ).then((data) => {
                setPage(data);
            }).catch((error) => {
                setNotifyState({
                    open: true,
                    message: error.message.message,
                    vertical: "top",
                    horizontal: "right",
                    severity: "error"
                });
            })
        }
    }

    return (
        <>
            <NotifySnackbar state={notifyState} setState={setNotifyState}/>
            <Stack spacing={2}>
                <Box sx={{display: "flex", p: 1, m: 1}}>
                    <Typography variant="h4">Search</Typography>
                </Box>
                <Divider/>

                <SearchField
                    trigger={handleChangeSearchRequest}
                    fieldDataMap={fieldDataMap}
                    setFieldDataMap={setFieldDataMap}
                />

                {!page &&
                    <>
                        <Skeleton/>
                        <Skeleton animation="wave"/>
                        <Skeleton animation={false}/>
                    </>
                }

                {page && page.content &&
                    <SearchTable
                        page={page}
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

export default Search;