import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogContentText from "@mui/material/DialogContentText";
import DialogTitle from "@mui/material/DialogTitle";
import {Box, Grid} from "@mui/material";


interface AlertDialogProps {
    open: boolean;
    setOpen: (flag: boolean) => void;
    title: string;
    description: string;
    confirmText: string;
    showCancelButton?: boolean;
    setConfirmDelete?: (flag: boolean) => void;
}

export default function AlertDialog({
                                        confirmText,
                                        description,
                                        open,
                                        setConfirmDelete,
                                        setOpen,
                                        showCancelButton,
                                        title
                                    }: AlertDialogProps) {
    const handleClose = () => {
        setOpen(false);
    };

    const handleConfirmDelete = (flag: boolean) => {
        setConfirmDelete?.(flag);
        handleClose();
    };

    return (
        <div>
            <Dialog
                open={open}
                aria-labelledby="alert-dialog-title"
                aria-describedby="alert-dialog-description"
            >
                <DialogTitle id="alert-dialog-title">
                    {title}
                </DialogTitle>
                <DialogContent>
                    <DialogContentText id="alert-dialog-description">
                        {description}
                    </DialogContentText>
                </DialogContent>
                <DialogActions>
                    {
                        !showCancelButton ? (
                            <Button onClick={handleClose} autoFocus>
                                {confirmText}
                            </Button>
                        ) : (
                            <>
                                <Button onClick={() => handleConfirmDelete(true)} variant="contained" color="success">
                                    {confirmText}
                                </Button>
                                <Button onClick={handleClose} variant="contained" color="error">
                                    No
                                </Button>
                            </>
                        )
                    }
                </DialogActions>
            </Dialog>
        </div>
    );
}