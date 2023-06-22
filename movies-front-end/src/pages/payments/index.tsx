import { Box, Divider, Paper, Stack, Typography } from "@mui/material";
import React, { useEffect } from "react";
import { useCheckTokenAndRole } from "@/hooks/auth/useCheckTokenAndRole";
import { signIn } from "next-auth/react";
import PaymentsTable from "@/components/Tables/PaymentsTable";

export default function Payments() {
  const isInvalid = useCheckTokenAndRole(["admin", "moderator"]);

  useEffect(() => {
    if (isInvalid) {
      signIn();
      return;
    }
  }, [isInvalid]);

  return (
    <Stack spacing={2}>
      <Box sx={{ p: 1, m: 1 }}>
        <Typography variant="h4">Your Payments</Typography>
      </Box>
      <Divider />

      <Paper elevation={6} sx={{ p: 2 }}>
        <Box sx={{ m: 1 }}>
          <PaymentsTable />
        </Box>
      </Paper>
    </Stack>
  );
}
