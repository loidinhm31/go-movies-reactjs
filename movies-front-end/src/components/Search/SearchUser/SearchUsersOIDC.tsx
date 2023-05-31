import {
    Container,
    FormControl,
    FormControlLabel,
    FormLabel,
    Grid,
    IconButton,
    Input,
    InputLabel,
    Radio,
    RadioGroup,
    TextField,
    Typography
} from "@mui/material";
import Button from "@mui/material/Button";
import React, {useState} from "react";
import useSWRMutation from "swr/mutation";
import {get, put} from "src/libs/api";
import {NotifyState} from "src/components/shared/snackbar";
import {RoleType, UserType} from "src/types/users";
import PersonAddIcon from "@mui/icons-material/PersonAdd";
import useSWR from "swr";

interface SearchUsersOIDCProps {
    setNotifyState: (state: NotifyState) => void;
    wasUpdated: boolean;
    setWasUpdated: (flag: boolean) => void;
}

export default function SearchUsersOIDC({setNotifyState, wasUpdated, setWasUpdated}: SearchUsersOIDCProps) {

    const [username, setUsername] = useState("");

    const [oidcUser, setOidcUser] = useState<UserType | null>(null);

    const [selectedRole, setSelectedRole] = useState<string>("");

    const {trigger: fetchUser} = useSWRMutation<UserType>(`/api/v1/admin/users/oidc?username=${username}`, get);
    const {trigger: putUser} = useSWRMutation(`/api/v1/admin/users/oidc`, put);
    const {data: roles} = useSWR<RoleType[]>("/api/v1/admin/roles", get);

    const handleSearchClick = () => {
        if (username !== "") {
            fetchUser().then((result) => {
                setOidcUser(result!);
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
    }

    const handleAddOidcUserClick = () => {
        if (selectedRole === undefined || selectedRole === "") {
            setNotifyState({
                open: true,
                message: "OIDC User need to set role",
                vertical: "bottom",
                horizontal: "center",
                severity: "warning"
            });
            return;
        }

        if (oidcUser) {
            oidcUser.role.role_code = selectedRole;
            putUser(oidcUser!).then((result) => {
                if (result.message === "ok") {
                    setNotifyState({
                        open: true,
                        message: "OIDC User added",
                        vertical: "top",
                        horizontal: "right",
                        severity: "success"
                    });

                    // Update table
                    setWasUpdated(true);

                    // Clear information
                    setOidcUser(null);
                    setSelectedRole("");
                    setUsername("");
                }
            }).catch((error) => {
                setNotifyState({
                    open: true,
                    message: error.message.message,
                    vertical: "top",
                    horizontal: "right",
                    severity: "error"
                });
            })
        }
    }

    return (
        <Grid container spacing={2}>
            <Grid item xs={12} sx={{p: 2}}>
                <TextField
                    fullWidth
                    label="OIDC Username"
                    variant="outlined"
                    value={username}
                    onChange={e => setUsername(e.target.value)}
                />
            </Grid>

            <Grid item>
                <Button
                    variant="contained"
                    onClick={handleSearchClick}
                >
                    Search
                </Button>
            </Grid>

            {oidcUser &&
                <Grid item xs={12}>
                    <Container maxWidth="sm">
                        <Grid container spacing={2}
                              component="form"
                              noValidate
                              autoComplete="off"
                        >
                            <Grid item xs={12}>
                                <Typography variant="overline">OIDC User</Typography>
                            </Grid>
                            <Grid item xs={6}>
                                <FormControl
                                    disabled
                                    variant="standard"
                                >
                                    <InputLabel htmlFor="component-disabled">Username</InputLabel>
                                    <Input value={oidcUser?.username}/>
                                </FormControl>
                            </Grid>
                            <Grid item xs={6}>
                                <FormControl
                                    disabled
                                    variant="standard"
                                >
                                    <InputLabel htmlFor="component-disabled">Email</InputLabel>
                                    <Input value={oidcUser?.email}/>
                                </FormControl>
                            </Grid>
                            <Grid item xs={6}>
                                <FormControl
                                    disabled
                                    variant="standard"
                                >
                                    <InputLabel htmlFor="component-disabled">First Name</InputLabel>
                                    <Input value={oidcUser?.first_name}/>
                                </FormControl>
                            </Grid>
                            <Grid item xs={6}>
                                <FormControl
                                    disabled
                                    variant="standard"
                                >
                                    <InputLabel htmlFor="component-disabled">Last Name</InputLabel>
                                    <Input value={oidcUser?.first_name}/>
                                </FormControl>
                            </Grid>

                            <Grid item xs={12}>
                                <FormControl>
                                    <FormLabel>Select Role</FormLabel>
                                    <RadioGroup
                                        row
                                        value={selectedRole}
                                        onChange={(event) => setSelectedRole(event.target.value)}
                                    >
                                        {roles &&
                                            roles.map((role, index) => {
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
                            </Grid>
                            <IconButton
                                color="primary"
                                size="large"
                                onClick={handleAddOidcUserClick}
                            >
                                <PersonAddIcon fontSize="inherit"/>
                            </IconButton>
                        </Grid>
                    </Container>
                </Grid>
            }

        </Grid>
    );
}