import DocsLayout from '@/components/Docs/DocsLayout'
import Nav from '@/components/LandingPage/Nav'
import { Badge } from '@/components/ui/badge'
import { GetServerSideProps } from 'next'
import React from 'react'

const Docs = () => {
  return (
    <>
    <DocsLayout>
      # Hi
    </DocsLayout>
    </>
  )
}

export default Docs
// export const getServerSideProps: GetServerSideProps = async (context) => {
//     return {
//       redirect: {
//         destination: '/docs/introduction',
//         permanent: true, // Set to true if this is a permanent redirect
//       },
//     };
//   };

Docs.getLayout = () => {
    return <></>
  }