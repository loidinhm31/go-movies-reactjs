import {Accordion, AccordionDetails, AccordionSummary, Grid, IconButton, Paper, Typography} from "@mui/material";
import EditIcon from "@mui/icons-material/Edit";
import {format} from "date-fns";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import {EpisodesTable} from "../Tables/EpisodesTable";
import React from "react";
import useSWRMutation from "swr/mutation";
import {SeasonType} from "../../types/seasons";
import {get} from "../../libs/api";
import useSWR from "swr";
interface SeasonInformationProps {
    movieId: number;
}

export default function SeasonInformation({movieId}: SeasonInformationProps) {
    const {data: seasons} = useSWR<SeasonType[]>(`/api/v1/seasons?movieId=${movieId}`, get)

    return (
        <Grid item xs={12} md={10}>
            <Paper elevation={5} sx={{p: 1}}>
                {seasons && seasons.map((s, index) => (
                    <Paper
                        key={`${s.id}-${index}`}
                        elevation={3}
                        sx={{
                            m: 2,
                            p: 2,
                            flexGrow: 1,
                        }}
                    >
                        <Grid container spacing={2}>
                            <Grid item xs={6}>
                                <Typography variant="subtitle1"><b>{s.name}</b></Typography>
                            </Grid>
                            <Grid item xs={4}>
                                <Typography
                                    variant="subtitle2">{format(new Date(s.air_date!), "MMMM do, yyyy")}</Typography>
                            </Grid>
                            <Grid item xs={12}>
                                <Typography variant="body1">{s.description}</Typography>
                            </Grid>
                            <Grid item xs={12}>
                                <Accordion TransitionProps={{unmountOnExit: true}}>
                                    <AccordionSummary
                                        expandIcon={<ExpandMoreIcon/>}
                                    >
                                        <Typography variant="caption">Episodes</Typography>
                                    </AccordionSummary>
                                    <AccordionDetails>
                                        <EpisodesTable
                                            season={s}
                                        />
                                    </AccordionDetails>
                                </Accordion>
                            </Grid>
                        </Grid>
                    </Paper>
                ))}
            </Paper>
        </Grid>
    );
}