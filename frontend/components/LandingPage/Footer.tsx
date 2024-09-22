import { GitHubLogoIcon } from '@radix-ui/react-icons'
import { Copyright, Github, GithubIcon, Heart, Key } from 'lucide-react'
import React from 'react'
import { Badge } from '../ui/badge'

const Footer = () => {
    return (
        <div className='flex items-center justify-center gap-4 flex-col'>
            <div className="flex gap-2 items-center">
                <Key size={27} />
                <p className="text-xl font-semibold">Switchr</p>
                <a href="https://github.com/heyyakash/switchr" target='_blank'><Badge>Github</Badge></a>
                <a href="https://docs.google.com/forms/d/e/1FAIpQLSf4Q8vVbrItyhW1iZ0sXQSKGbjKKRvi4G8F-iNLpsRkn3TQ-A/viewform?usp=sf_link" target='_blank'><Badge variant={"outline"}>Feedback</Badge></a>

            </div>
           
            <p className='text-lg flex gap-2 items-center text-primary/70'> <Copyright /> Switchr, Based in India</p>
            <p className='flex items-center gap-2'>Made with  <Heart className='text-rose-500' strokeWidth={2} />By <a className='underline' href="https://github.com/heyyakash" target='_blank'>Akash Sharma</a></p>
        </div>
    )
}

export default Footer