import Head from "next/head";
import Link from "next/link";
import {useSession} from "next-auth/react";
import React from "react";
import {Box, Grid, Typography} from "@mui/material";
import Divider from "@mui/material/Divider";
import EditIcon from '@mui/icons-material/Edit';


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
				<Box m="auto" className="max-w-7xl" alignContent="center">
					<Box className="p-4 sm:p-6 w-full">
						<Box display="block" fontSize="2xl" py={2}>
							<b>Your Account</b>
						</Box>
						<Divider />
						<Grid gridTemplateColumns="repeat(2, max-content)" alignItems="center" gap={6} py={4}>
							<Typography><b>ID</b></Typography>
							<Typography>{session.user.id ?? "(No ID)"}</Typography>
							<Typography><b>Username</b></Typography>
							<Box gap={2}>
								{session.user.name ?? "(No username)"}
								<Link href="/account/edit">
									<EditIcon />
								</Link>
							</Box>
							<Typography><b>Email</b></Typography>
							<Typography>{session.user.email ?? "(No Email)"}</Typography>
						</Grid>
						<p></p>
					</Box>
				</Box>
			</main>
		</>
	);
}
