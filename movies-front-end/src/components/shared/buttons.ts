import {Button, ButtonProps, styled} from "@mui/material";
import {blue, blueGrey} from "@mui/material/colors";

export const SelectButton = styled(Button)<ButtonProps>(({ theme}) => ({
    color: theme.palette.getContrastText(blue[500]),
    backgroundColor: blue[500],
    "&:hover": {
        backgroundColor: blue[700],
    },
}));

export const UnselectButton = styled(Button)<ButtonProps>(({ theme}) => ({
    color: theme.palette.getContrastText(blueGrey[500]),
    backgroundColor: blueGrey[500],
    "&:hover": {
        backgroundColor: blueGrey[700],
    },
}));
