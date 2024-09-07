import Layout from "@/components/Layout/Layout";
import { ThemeProvider } from "@/components/theme-provider";
import "@/styles/globals.css";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { NextComponentType, NextPageContext } from "next";
import { Toaster } from "sonner";


type ComponentType = {
  Component: NextComponentType<NextPageContext, any, any> & { getLayout?: JSX.Element }
  pageProps: any
}

const isComponentMDX = (component: any) => {
  console.log("component", component.isMDX)
  return component?.isMDX;
}

export default function App({ Component, pageProps }: ComponentType) {
  const layout = Component.getLayout
  const queryClient = new QueryClient()
  isComponentMDX(Component)
  return (
    <QueryClientProvider client={queryClient}>
      <ThemeProvider
        attribute="class"
        defaultTheme="system"
        enableSystem
        disableTransitionOnChange
      >
        
        {layout && !isComponentMDX(Component) ? (
          <Component {...pageProps} />
        ) : (
          <Layout>

            <Component {...pageProps} />
          </Layout>
        )}
        <Toaster expand={true} richColors />
      </ThemeProvider></QueryClientProvider>
  )
}
