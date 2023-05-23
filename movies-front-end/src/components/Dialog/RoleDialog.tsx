import Dialog from "@mui/material/Dialog";
import DialogContent from "@mui/material/DialogContent";
import React, {useState} from "react";
import useSWR from "swr";
import {get, patch} from "src/libs/api";
import {Button, DialogActions, FormControl, FormControlLabel, FormLabel, Radio, RadioGroup} from "@mui/material";
import {RoleType, UserType} from "../../types/users";
import DialogTitle from "@mui/material/DialogTitle";
import useSWRMutation from "swr/mutation";
import {NotifyState} from "../shared/snackbar";

interface RoleDialogProps {
    user: UserType | null;
    open: boolean;
    setOpen: (flag: boolean) => void;
    setNotifyState: (state: NotifyState) => void;
    setWasUpdated: (flag: boolean) => void;
}

const RoleDialog = ({user, open, setOpen, setNotifyState, setWasUpdated}: RoleDialogProps) => {
    const [selectedRole, setSelectedRole] = useState<string>(user!.role.role_code);

    const {data} = useSWR<RoleType[]>("/api/v1/admin/roles", get);
    const {trigger: updateRole} = useSWRMutation("/api/v1/admin/users", patch)

    const handleUpdateRole = () => {

        updateRole({
            ...user,
            role: {
                id: undefined,
                role_name: undefined,
                role_code: selectedRole,
            }
        } as UserType).then((result) => {
            if (result.message === "ok") {
                setWasUpdated(true);

                handleClose();

                setNotifyState({
                    open: true,
                    message: "Update role successfully",
                    vertical: "top",
                    horizontal: "right",
                    severity: "success"
                });
            } else {
                setNotifyState({
                    open: true,
                    message: "Cannot update role",
                    vertical: "top",
                    horizontal: "right",
                    severity: "error"
                });
            }
        }).catch((error) => {
            setNotifyState({
                open: true,
                message: error.message.message,
                vertical: "top",
                horizontal: "right",
                severity: "error"
            });
        });
    }

    const handleClose = () => {
        setOpen(false);
    };

    return (
        <>
            <Dialog
                fullWidth={true}
                maxWidth={"lg"}
                open={open}
                onClose={handleClose}
            >

                <DialogTitle>
                    {`Choose Role for \"${user?.username}\"`}
                </DialogTitle>
                <DialogContent>
                    <FormControl>
                        <FormLabel>Role</FormLabel>
                        <RadioGroup
                            row
                            value={selectedRole}
                            onChange={(event) => setSelectedRole(event.target.value)}
                        >
                            {data &&
                                data.map((role, index) => {
                                    return (
                                        <FormControlLabel
                                            key={`${role.id}-${index}`}
                                            value={role.role_code}
                                            control={<Radio/>}
                                            label={role.role_code}
                                        />
                                    );
                                })
                            }
                        </RadioGroup>
                    </FormControl>

                </DialogContent>
                <DialogActions>
                    <Button variant="outlined" color="error" onClick={handleClose}>CANCEL</Button>
                    <Button variant="contained" color="success" onClick={handleUpdateRole} autoFocus>
                        APPROVE
                    </Button>
                </DialogActions>
            </Dialog>
        </>
    );
}

export default RoleDialog;