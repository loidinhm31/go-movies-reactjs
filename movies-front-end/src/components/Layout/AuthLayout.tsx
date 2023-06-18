import { Box } from "@mui/material";

export function AuthLayout({ children }) {
    return (
        <Box className="subpixel-antialiased" sx={{ display: "flex", justifyContent: "center", p: 1 }}>
            <Box sx={{ p: 1 }}>{children}</Box>
        </Box>
    );
}
