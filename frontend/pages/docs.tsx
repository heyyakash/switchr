import Aside from '@/components/Docs/DocsAside'
import Nav from '@/components/LandingPage/Nav'
import GettingStarted from "./test.md"
import rehypeHighlight from 'rehype-highlight'
import Link from 'next/link'
import React, { ReactNode } from 'react'
import ReactMarkdown from 'react-markdown';
import { useMDXComponents } from '@/mdx-components'

const links = [
    { name: "Introduction", href: "/docs/introduction" },
    { name: "Getting Started", href: "/docs/getting-started" },
    { name: "New Project", href: "/docs/components" },
    { name: "New Flags", href: "/docs/customization" },
    { name: "FAQs", href: "/docs/faqs" },
  ];

interface props {
    children : ReactNode
}

const Docs: React.FC<props> = ({children}) => {
  return (
    <>
        <Nav />
        <div className='max-w-[1200px] flex  relative py-[6rem] w-full mx-auto'>
            <Aside>
                <nav className='flex flex-col gap-2'>
                    {links.map((x,i)=>{
                        return(
                            <Link href = {x.href} key = {i}>{x.name}</Link> 
                        )
                    })}
                </nav>
            </Aside>
            <div className='no-tailwind '>
            </div>

        </div> 
    </>
  )
}

export default Docs


Docs.getLayout = () => {
    return <></>
  }


              {/* <h2 className='text-3xl font-semibold text-center pt-[4rem]'>
                Accessing feature flags using REST API
            </h2>
            <div className='text-primary/70 mt-12 text-xl px-4 flex flex-col gap-3'>
                <p>The following procedures can be followed once a project is created and some flags are added in it</p>
                <p><span className='font-bold'>Step 1</span>. Login and open the project whose keys you wish to access and the click on <Badge>API</Badge> button on the project page.</p>
                <p><span className='font-bold'>Step 2</span>. Once a dialog box appears click on <Badge>Generate Key</Badge> button, this will generate key to access the flag.</p>
                <p><span className='font-bold'>Step 3</span>. Once the key is generated, it is valid for 120 days and shall be sent in the header called <span className='font-extrabold text-primary'>token</span> in a GET Request to the following URL</p>
                <div className='p-3 w-full overflow-auto bg-secondary/50 rounded-lg'>
                    {process.env.NEXT_PUBLIC_BASE_URL}/api/get/&lt;FLAG_NAME&gt;
                </div>
                <b>Example Request</b>
                <div className='p-3 w-full overflow-auto bg-secondary/50 rounded-lg'>
                curl -X GET &apos;{process.env.NEXT_PUBLIC_BASE_URL}/api/get/FLAG_NAME&apos;
                --header &apos;token: &lt;YOUR_TOKEN / KEY&gt;&apos;;
                </div>
            </div> */}