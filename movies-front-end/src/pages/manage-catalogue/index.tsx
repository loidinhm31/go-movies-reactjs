import React, {useEffect, useState} from "react";
import {useRouter} from "next/router";
import useSWR from "swr";
import {get} from "../../libs/api";
import {Box, Button, Divider, Stack, Typography} from "@mui/material";
import {useSession} from "next-auth/react";
import EnhancedTable from "../../components/movies/table/MoviesTable";
import useSWRMutation from "swr/mutation";

const ManageCatalogue = () => {
    const session = useSession();
    const router = useRouter();

    const [isConfirmDelete, setIsConfirmDelete] = useState<boolean>(false);
    const [deleteId, setDeleteId] = useState<number>();

    const { data: movies, isLoading } = useSWR(`../api/v1/movies`, get);
    const {trigger: deleteMovie} = useSWRMutation(`../api/v1/admin/movies/delete/${deleteId}`, get);


    useEffect( () => {
        // if (!session) {
        //     router.push("/auth/signin");
        //     return
        // }

    }, [router]);

    useEffect(() => {
        if (deleteId && isConfirmDelete) {
            deleteMovie()
                .then((data) => {
                    if (data.error) {
                        console.log(data.error);
                    } else {
                        router.push("/manage-catalogue");
                    }
                })
                .catch(err => {
                    console.log(err);
                })
                .finally(() => {
                    setIsConfirmDelete(false);
                    setDeleteId(null);
                })
        }
    }, [deleteId])

    return(
        <Stack spacing={2}>
            <Box sx={{display: "flex", p: 1, m: 1}}>
                <Typography variant="h4">Manage Catalogue</Typography>
            </Box>
            <Divider/>
            <Box>
                <Button variant="contained" href="/admin/movies" >Add Movie</Button>
            </Box>

            {movies &&
                <EnhancedTable
                    rows={movies}
                    setDeleteId={setDeleteId}
                    confirmDelete={isConfirmDelete}
                    setConfirmDelete={setIsConfirmDelete} />
            }
        </Stack>
    )
}

export default ManageCatalogue;