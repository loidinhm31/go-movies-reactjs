import Head from "next/head";
import {useRouter} from "next/router";
import React, {useEffect, useState} from "react";
import {useForm} from "react-hook-form";
import {Footer} from "src/components/Footer";
import {Header} from "src/components/Header";
import {AuthLayout} from "../../components/layout/AuthLayout";
import {ClientSafeProvider, getProviders, signIn} from "next-auth/react";
import {Role, RoleSelect} from "../../components/RoleSelect";
import {GetServerSideProps} from "next";
import {
    Box,
    Button,
    ButtonProps,
    Chip,
    Divider, FormHelperText,
    InputAdornment,
    OutlinedInput,
    Stack,
    TextField,
    Typography
} from "@mui/material";
import PersonIcon from "@mui/icons-material/Person";
import BugReportIcon from "@mui/icons-material/BugReport";
import {DrawerHeader} from "../../components/shared/drawer";

export type SignInErrorTypes =
    | "Signin"
    | "EmailCreateAccount"
    | "Callback"
    | "CredentialsSignin"
    | "SessionRequired"
    | "default";

const errorMessages: Record<SignInErrorTypes, string> = {
    Signin: "Try signing in with a different account.",
    EmailCreateAccount: "Try signing in with a different account.",
    Callback: "Try signing in with a different account.",
    CredentialsSignin: "Sign in failed. Check the details you provided are correct.",
    SessionRequired: "Please sign in to access this page.",
    default: "Unable to sign in."
};

interface SigninProps {
    providers: Awaited<ReturnType<typeof getProviders>>;
}

function Signin({providers}: SigninProps) {
    const router = useRouter();
    const {debug, server} = providers;
    const [error, setError] = useState("");

    useEffect(() => {
        const err = router?.query?.error;
        if (err) {
            if (typeof err === "string") {
                setError(errorMessages[err]);
            } else {
                setError(errorMessages[err[0]]);
            }
        }
    }, [router]);

    return (
        <>
            <Head>
                <title>Sign In</title>
            </Head>
            <AuthLayout>
                <DrawerHeader/>
                <Stack spacing="2">
                    {debug && <DebugSigninForm credentials={debug}/>}
                    {server && <SigninForm credentials={server}/>}
                </Stack>

                {error && (
                    <div className="text-center mt-8">
                        <p className="text-orange-600">Error: {error}</p>
                    </div>
                )}
            </AuthLayout>
        </>
    );
}

Signin.getLayout = (page) => (
    <div className="grid grid-rows-[min-content_1fr_min-content] h-full justify-items-stretch">
        <Header/>
        {page}
        <Footer/>
    </div>
);

export default Signin;

const SigninButton = (props: ButtonProps) => {
    return (
        <Button
            variant="contained"
            startIcon={props.startIcon}
            type="submit"
            {...props}>
        </Button>
    );
};

interface SigninFormData {
    username: string;
    password: string;
}

const SigninForm = ({credentials}: { credentials: ClientSafeProvider; }) => {
    const {register, handleSubmit} = useForm<SigninFormData>();

    function signinWithCredentials(data: SigninFormData) {
        signIn(credentials.id, {
            callbackUrl: "/dashboard",
            ...data
        });
    }

    return (
        <Box component="form"
             onSubmit={handleSubmit(signinWithCredentials)}>
            <Stack spacing={2}>
                <TextField label="Username" variant="outlined" {...register("username")} />
                <TextField label="Password" variant="outlined" type="password" {...register("password")} />

                <Box sx={{display: "flex", justifyContent: "center"}}>
                    <SigninButton startIcon={<PersonIcon/>}>Continue with User</SigninButton>
                </Box>
            </Stack>
        </Box>
    );
};

interface DebugSigninFormData {
    username: string;
    role: Role;
}

const DebugSigninForm = ({credentials}: { credentials: ClientSafeProvider; }) => {
    const {register, handleSubmit} = useForm<DebugSigninFormData>({
        defaultValues: {
            role: "general",
            username: "dev"
        }
    });

    function signinWithDebugCredentials(data: DebugSigninFormData) {
        signIn(credentials.id, {
            callbackUrl: "/dashboard",
            ...data
        });
    }

    return (
        <Box component="form" sx={{
            color: "orange", m: 2, p: 2,
            border: 1, borderRadius: 1, borderWidth: 5
        }}
             onSubmit={handleSubmit(signinWithDebugCredentials)}
        >

            <Stack spacing={2}>
                <TextField label="Username" variant="outlined" {...register("username")} />

                <RoleSelect {...register("role")}></RoleSelect>

                <Box sx={{display: "flex", justifyContent: "center"}}>
                    <SigninButton startIcon={<BugReportIcon/>}>Continue with Debug User</SigninButton>
                </Box>
            </Stack>
            <Divider sx={{p: 2}}>
                <Chip sx={{color: "white", background: "orange"}} label="FOR DEBUGGING ONLY"/>
            </Divider>
        </Box>
    );
};
export const getServerSideProps: GetServerSideProps<SigninProps> = async ({locale}) => {
    const providers = await getProviders();
    return {
        props: {
            providers,
        }
    };
};