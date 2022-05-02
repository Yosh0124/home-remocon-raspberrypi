import 'bootstrap/dist/css/bootstrap.min.css';
import '../styles/globals.css'
import type { AppProps } from 'next/app'
import Head from 'next/head';

function MyApp({ Component, pageProps }: AppProps) {
  return (
  <>
    <Head>
      <meta name="theme-color" content="#39f" />
      <link rel="manifest" href="/manifest.webmanifest" />
    </Head>
    <Component {...pageProps} />
  </>
  )
}

export default MyApp
