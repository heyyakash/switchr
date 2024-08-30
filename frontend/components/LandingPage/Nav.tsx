import { Key} from 'lucide-react'
import React from 'react'
import { Button } from '../ui/button'
import { ModeToggle } from '../Toggle'
import Link from 'next/link'

const Nav = () => {
    return (
        <nav className="w-full px-4 lg:px-0 h-[70px] fixed top-0 backdrop-blur-sm  border-b-[1px]">
            <div className="max-w-[1200px] w-full py-4 mx-auto flex items-center justify-between">
                <Link href = "/" className="flex gap-2 items-center">
                    <Key size={27} />
                    <p className="text-xl font-semibold">Switchr</p>
                </Link>
                <div className="flex items-center gap-3">
                    <Link href = "/docs"><Button variant={"secondary"}>API Docs</Button></Link>
                    <Link href="/login"><Button>Login</Button></Link>
                    <ModeToggle />
                </div>
            </div>
        </nav>
    )
}

export default Nav