import {
    Button,
    CardMedia,
    FormControl,
    Grid,
    InputLabel,
    MenuItem,
    Pagination,
    Paper,
    Select,
    SelectChangeEvent,
    Stack,
    Typography,
} from "@mui/material";
import Link from "next/link";
import format from "date-fns/format";
import { useEffect, useState } from "react";
import useSWR from "swr";
import { get } from "src/libs/api";
import SearchKey from "src/components/Search/SearchKey";
import { PageType } from "src/types/page";
import { MovieType } from "src/types/movies";

export function CollectionMovieTab() {
    const [searchKey, setSearchKey] = useState<string>("");

    const [pageIndex, setPageIndex] = useState(0);
    const [pageSize, setPageSize] = useState(9);

    const { data: page } = useSWR<PageType<MovieType>>(
        `/api/v1/collections?type=MOVIE&q=${searchKey}&pageIndex=${pageIndex}&pageSize=${pageSize}`,
        get
    );

    useEffect(() => {
        setPageIndex(0);
    }, [pageSize, searchKey]);

    const handlePageIndexChange = (event: React.ChangeEvent<unknown>, value: number) => {
        setPageIndex(value - 1);
    };

    const handlePageSizeChange = (event: SelectChangeEvent) => {
        const val = event.target.value;
        setPageSize(parseInt(val));
    };

    return (
        <Grid container spacing={1}>
            <Grid item xs={12}>
                <SearchKey keyword={searchKey} setKeyword={setSearchKey} />
            </Grid>
            <Grid item container xs={12} sx={{ p: 1, m: 1 }}>
                {page &&
                    page.content &&
                    page.content.map((movie) => (
                        <Grid key={movie.id} item xs={3}>
                            <Link href={`/movies/${movie.id}`} style={{ textDecoration: "none" }}>
                                <Paper
                                    sx={{
                                        m: 2,
                                        p: 2,
                                        flexGrow: 1,
                                    }}
                                >
                                    <Stack spacing={3}>
                                        <Grid item xs>
                                            <Stack direction="row" justifyContent="space-between">
                                                <Typography
                                                    gutterBottom
                                                    variant="subtitle1"
                                                    component="div"
                                                    sx={{ m: 1 }}
                                                >
                                                    <b>{movie.title}</b>
                                                </Typography>
                                                <Button color="success" sx={{ m: 1 }}>
                                                    <b>{`${movie.price ? movie.price + " USD" : "FREE"}`}</b>
                                                </Button>
                                            </Stack>
                                        </Grid>
                                        <Grid item container spacing={1}>
                                            <Grid item xs={6}>
                                                <CardMedia
                                                    sx={{ borderRadius: "16px" }}
                                                    component="img"
                                                    src={movie.image_url}
                                                />
                                            </Grid>
                                            <Grid item xs={6}>
                                                <Typography variant="body2" gutterBottom>
                                                    {format(new Date(movie.release_date!), "MMMM do, yyyy")}
                                                </Typography>
                                                <Typography gutterBottom variant="inherit" component="div">
                                                    {movie.description}
                                                </Typography>
                                            </Grid>
                                        </Grid>
                                    </Stack>
                                </Paper>
                            </Link>
                        </Grid>
                    ))}
            </Grid>
            <Grid item xs={12} sx={{ display: "flex", justifyContent: "center" }}>
                <Stack spacing={2} direction="row">
                    <FormControl>
                        <InputLabel>Size</InputLabel>
                        <Select
                            sx={{ display: "flex", alignItems: "center" }}
                            value={pageSize.toString()}
                            label="Size"
                            onChange={handlePageSizeChange}
                        >
                            {process.env.NODE_ENV === "development" && <MenuItem value={1}>1</MenuItem>}
                            <MenuItem value={9}>9</MenuItem>
                            <MenuItem value={18}>18</MenuItem>
                            <MenuItem value={27}>27</MenuItem>
                        </Select>
                    </FormControl>
                    <Pagination
                        sx={{ display: "flex", alignItems: "center" }}
                        page={pageIndex + 1}
                        count={page?.total_pages}
                        onChange={handlePageIndexChange}
                    />
                </Stack>
            </Grid>
        </Grid>
    );
}
