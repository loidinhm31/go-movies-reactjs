import "../styles/global.css";
import {SWRConfig, SWRConfiguration} from "swr";
import {SessionProvider} from "next-auth/react";
import {AppProps} from "next/app";
import {getDefaultLayout, NextPageWithLayout} from "../components/Layout/Layout";
import Head from "next/head";

type AppPropsWithLayout = AppProps & {
    Component: NextPageWithLayout;
};

const swrConfig: SWRConfiguration = {
    revalidateOnFocus: false,
    revalidateOnMount: true,
};

function MyApp({ Component, pageProps: { session, cookies, ...pageProps } }: AppPropsWithLayout) {
    const getLayout = Component.getLayout ?? getDefaultLayout;
    const page = getLayout(<Component {...pageProps} />);

    return (

        <SWRConfig value={swrConfig}>
            <SessionProvider session={session}>
                <Head>
                    <meta name="viewport" content="initial-scale=1, width=device-width" />
                </Head>
                {page}
            </SessionProvider>
        </SWRConfig>
    );
}

export default MyApp;