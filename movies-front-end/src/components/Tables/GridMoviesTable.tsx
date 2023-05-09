import {
    ButtonBase,
    Chip,
    FormControl,
    Grid,
    InputLabel,
    MenuItem,
    Pagination,
    Paper,
    Select,
    SelectChangeEvent,
    Stack,
    Typography
} from "@mui/material";
import Link from "next/link";
import {PageType} from "../../types/page";
import {MovieType} from "../../types/movies";
import {styled} from "@mui/material/styles";
import format from "date-fns/format";

const Img = styled("img")({
    margin: "auto",
    display: "block",
    maxWidth: "100%",
    maxHeight: "100%",
});


interface GridTableProps {
    page: PageType<MovieType>;
    pageIndex: number;
    setPageIndex: (value: number) => void;
    pageSize: number;
    setPageSize: (value: number) => void;
}

export function GridMovies({
                               page,
                               pageIndex,
                               pageSize,
                               setPageIndex,
                               setPageSize
                           }: GridTableProps) {
    const handlePageIndexChange = (event: React.ChangeEvent<unknown>, value: number) => {
        setPageIndex(value - 1);
    };

    const handlePageSizeChange = (event: SelectChangeEvent) => {
        const val = event.target.value
        setPageSize(parseInt(val));
    };

    return (
        <Grid container spacing={2}>
            {page && page.data && page.data.map((movie) => (
                <Grid key={movie.id} item xs={4}>
                    <Link href={`/movies/${movie.id}`} style={{textDecoration: "none"}}>
                        <Paper
                            sx={{
                                m: 2,
                                p: 2,
                                flexGrow: 1,
                            }}
                        >
                            <Grid container spacing={2}>
                                <Grid item>
                                    <ButtonBase sx={{width: 128, height: 128}}>
                                        <Img src={`https://image.tmdb.org/t/p/w200/${movie.image}`}/>
                                    </ButtonBase>
                                </Grid>
                                <Grid item xs container direction="column" spacing={2}>
                                    <Grid item xs>
                                        <Typography gutterBottom variant="subtitle1" component="div">
                                            <b>{movie.title}</b>
                                        </Typography>
                                        <Typography variant="body2" gutterBottom>
                                            {format(new Date(movie.release_date!), "MMMM do, yyyy")}
                                        </Typography>
                                        <Chip label={movie.mpaa_rating} color="error"/>
                                    </Grid>
                                </Grid>
                            </Grid>
                        </Paper>
                    </Link>

                </Grid>
            ))}

            <Grid item xs={12} sx={{display: "flex", justifyContent: "center"}}>
                <Stack spacing={2} direction="row">
                    <FormControl>
                        <InputLabel>Size</InputLabel>
                        <Select
                            sx={{display: "flex", alignItems: "center"}}
                            value={pageSize.toString()}
                            label="Size"
                            onChange={handlePageSizeChange}
                        >
                            {process.env.NODE_ENV === "development"
                                && <MenuItem value={1}>1</MenuItem>}
                            <MenuItem value={9}>9</MenuItem>
                            <MenuItem value={18}>18</MenuItem>
                            <MenuItem value={27}>27</MenuItem>
                        </Select>
                    </FormControl>
                    <Pagination
                        sx={{display: "flex", alignItems: "center"}}
                        page={pageIndex + 1}
                        count={page.total_pages}
                        onChange={handlePageIndexChange}
                    />
                </Stack>
            </Grid>
        </Grid>
    );
}