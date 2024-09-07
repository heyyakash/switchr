import React, { ReactNode } from 'react'
import Nav from '../LandingPage/Nav'
import Aside from './DocsAside'
import Link from 'next/link';
import { useRouter } from 'next/router';


const links = [
    { name: "Introduction", href: "/docs/introduction" },
    { name: "Getting Started", href: "/docs/gettingstarted" },
    { name: "New Project", href: "/docs/components" },
    { name: "New Flags", href: "/docs/customization" },
    { name: "FAQs", href: "/docs/faqs" },
  ];

  interface props {
    children : ReactNode
}
  const DocsLayout= (props:any) => {
    const router = useRouter()
    const path = router.pathname
    console.log(path)
  return (
    <div className='relative'>
        <Nav />
        <div className='max-w-[1200px]  flex  relative py-[6rem] z-[1] w-full mx-auto'>
            <Aside>
                <nav className='flex flex-col gap-2'>
                    {links.map((x,i)=>{
                        return(
                            <Link href = {x.href} key = {i} className={`${path && path === x.href ? "font-bold text-white":""}`}>{x.name}</Link> 
                        )
                    })}
                </nav>
            </Aside>
            <div className='prose dark:prose-invert mt-4'>
                {props.children}
            </div>

        </div> 
    </div>
  )
}

DocsLayout.isMDX = true

export default DocsLayout

