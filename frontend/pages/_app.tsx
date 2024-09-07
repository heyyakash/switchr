import Layout from "@/components/Layout/Layout";
import { ThemeProvider } from "@/components/theme-provider";
import "@/styles/globals.css";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { NextComponentType, NextPageContext } from "next";
import type { AppProps } from "next/app";
import { useState } from "react";
import { Toaster } from "sonner";
import { MDXProvider } from '@mdx-js/react';

type ComponentType = {
   Component: NextComponentType<NextPageContext, any, any> & { getLayout?: JSX.Element }
   pageProps: any
 } 
 const components = {
  h1: (props:any) => <h1 style={{ color: 'tomato' }} {...props} />,
  // Add more custom components if needed
};


 export default function App({ Component, pageProps }: ComponentType) {
  const layout = Component.getLayout
  const queryClient = new QueryClient()
  return (
    <QueryClientProvider client={queryClient}>
      <ThemeProvider
        attribute="class"
        defaultTheme="system"
        enableSystem
        disableTransitionOnChange
      >
         <MDXProvider components={components}>
         <Component {...pageProps} />
        {/* {layout?(
          <Component {...pageProps} />
          ):(
            <Layout>
          
            <Component {...pageProps} />
            </Layout>
          )}
        <Toaster expand = {true} richColors/> */}
        </MDXProvider>
      </ThemeProvider></QueryClientProvider>
  )
}
