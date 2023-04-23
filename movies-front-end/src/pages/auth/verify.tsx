import {GetServerSideProps} from "next";
import Head from "next/head";
import {getCsrfToken, getProviders} from "next-auth/react";

export default function Verify() {
    return (
        <>
            <Head>
                <title>Sign Up - ShiftFlix</title>
                <meta name="Sign Up" content="Sign up to access ShiftFlix"/>
            </Head>
            <div>
                <div>
                    <h1>A sign-in link has been sent to your email address (likely going to
                        spam).
                    </h1>
                </div>
            </div>
        </>
    );
}

export const getServerSideProps: GetServerSideProps = async () => {
    const csrfToken = await getCsrfToken();
    const providers = await getProviders();
    return {
        props: {
            csrfToken,
            providers,
        },
    };
};
