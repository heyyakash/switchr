import DocsLayout from "@/components/Docs/DocsLayout";
import Layout from "@/components/Layout/Layout";
import { ThemeProvider } from "@/components/theme-provider";
import "@/styles/globals.css";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { NextComponentType, NextPageContext } from "next";
import Router, { useRouter } from "next/router";
import { Toaster } from "sonner";


type ComponentType = {
  Component: NextComponentType<NextPageContext, any, any> & { getLayout?: JSX.Element }
  pageProps: any
}

const isComponentMDX = (component: any) => {
  return component === "MDXContent"
}

export default function App({ Component, pageProps }: ComponentType) {
  let layout = Component.getLayout
  const router = useRouter()
  const isDocsPage = router.pathname.startsWith('/docs');
  if (isDocsPage){
    layout= <></>
  }
  const queryClient = new QueryClient()
  
  return (
    <QueryClientProvider client={queryClient}>
      <ThemeProvider
        attribute="class"
        defaultTheme="system"
        enableSystem
        disableTransitionOnChange
      >
        {layout? isDocsPage ?(
          <DocsLayout>
          <Component {...pageProps} />
          </DocsLayout>
        ):(
          <Component {...pageProps} />
        )
         : (
          <Layout>

            <Component {...pageProps} />
          </Layout>
        )}
        <Toaster expand={true} richColors />
      </ThemeProvider></QueryClientProvider>
  )
}
