import React, { ReactNode, useState } from 'react'
import Nav from '../LandingPage/Nav'
import Aside from './DocsAside'
import Link from 'next/link';
import { useRouter } from 'next/router';
import { HamburgerMenuIcon } from '@radix-ui/react-icons';


const links = [
  { name: "Introduction", href: "/docs/introduction" },
  { name: "Getting Started", href: "/docs/gettingstarted" },
  { name: "Create a new Project", href: "/docs/newproject" },
  { name: "Create new Feature Flags", href: "/docs/newflag" },
  { name: "REST API", href: "/docs/restapi" },
];

interface props {
  children: ReactNode
}
const DocsLayout = (props: any) => {
  const router = useRouter()
  const path = router.pathname
  const [open,setOpen] = useState(false)
  return (
    <div className='relative'>
      <Nav />

      <div className='max-w-[1200px]  flex flex-col lg:flex-row relative py-[6rem] z-[1] w-full mx-auto'>
        <HamburgerMenuIcon onClick={()=>setOpen(true)} className='mx-4 lg:hidden' />
        <Aside open = {open} setOpen = {setOpen}>
          <nav className='flex flex-col gap-2'>
            {links.map((x, i) => {
              return (
                <Link onClick={()=>setOpen(false)} href={x.href} key={i} className={`${path && path === x.href ? "font-bold text-primary" : ""}`}>{x.name}</Link>
              )
            })}
          </nav>
        </Aside>
        <div className='prose w-full p-4 dark:prose-invert mt-4'>
          {props.children}
        </div>

      </div>
    </div>
  )
}

DocsLayout.isMDX = true

export default DocsLayout

