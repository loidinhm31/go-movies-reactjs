import "@/styles/global.css";
import { SWRConfig, SWRConfiguration } from "swr";
import { SessionProvider } from "next-auth/react";
import { AppProps } from "next/app";
import { getDefaultLayout, NextPageWithLayout } from "@/components/Layout/Layout";
import Head from "next/head";
import wrapper from "@/redux/reduxWrapper";
import { Provider } from "react-redux";

type AppPropsWithLayout = AppProps & {
  Component: NextPageWithLayout;
};

const swrConfig: SWRConfiguration = {
  revalidateOnFocus: false,
  revalidateOnMount: true
};

function MyApp({ Component, pageProps: { session, cookies, ...pageProps } }: AppPropsWithLayout) {
  const getLayout = Component.getLayout ?? getDefaultLayout;
  const page = getLayout(<Component {...pageProps} />);

  const { store, props } = wrapper.useWrappedStore(pageProps);

  return (
    <SWRConfig value={swrConfig}>
      <SessionProvider session={session}>

        <Provider store={store}>
          <Head>
            <meta name="viewport" content="initial-scale=1, width=device-width" />
          </Head>
          {page}
        </Provider>
      </SessionProvider>
    </SWRConfig>
  );
}

export default MyApp;
