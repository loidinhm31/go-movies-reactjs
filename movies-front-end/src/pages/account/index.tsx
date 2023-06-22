import Head from "next/head";
import Link from "next/link";
import { useSession } from "next-auth/react";
import React from "react";
import { Box, Container, Grid, IconButton, Stack, Typography } from "@mui/material";
import Divider from "@mui/material/Divider";
import EditIcon from "@mui/icons-material/Edit";

export default function Account() {
  const { data: session } = useSession();

  if (!session) {
    return;
  }
  return (
    <>
      <Head>
        <title>My Account</title>
      </Head>
      <main className="oa-basic-theme p-6">
        <Stack spacing={2}>
          <Box display="block" fontSize="2xl" py={2}>
            <b>Your Account</b>
          </Box>
          <Divider />
          <Container maxWidth="sm" sx={{ p: 1, m: 1 }}>
            <Grid container spacing={3} sx={{ m: 2 }}>
              <Grid item xs={12} md={6}>
                <Box>
                  <Typography>
                    <b>ID</b>
                  </Typography>
                </Box>
              </Grid>
              <Grid item xs={12} md={6}>
                <Box>
                  <Typography>{session.user.id ?? "(No ID)"}</Typography>
                </Box>
              </Grid>

              <Grid item xs={12} md={6}>
                <Box>
                  <Typography>
                    <b>Username</b>
                  </Typography>
                </Box>
              </Grid>

              <Grid item xs={12} md={6}>
                <Box>
                  {session.user.name ?? "(No username)"}
                  <Link href="/account/edit">
                    <IconButton color="primary">
                      <EditIcon />
                    </IconButton>
                  </Link>
                </Box>
              </Grid>

              <Grid item xs={12} md={6}>
                <Box>
                  <Typography>
                    <b>Role</b>
                  </Typography>
                </Box>
              </Grid>
              <Grid item xs={12} md={6}>
                <Box>
                  <Typography>
                    <b>{session.user.role}</b>
                  </Typography>
                </Box>
              </Grid>
              <Grid item xs={12} md={6}>
                <Box>
                  <Typography>
                    <b>Email</b>
                  </Typography>
                </Box>
              </Grid>
              <Grid item xs={12} md={6}>
                <Box>
                  <Typography>{session.user.email ?? "(No Email)"}</Typography>
                </Box>
              </Grid>
            </Grid>
          </Container>
        </Stack>
      </main>
    </>
  );
}
