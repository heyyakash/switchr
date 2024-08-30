// components/Steps.js
import React from 'react';
import {
    Carousel,
    CarouselContent,
    CarouselItem,
    CarouselNext,
    CarouselPrevious,
} from "@/components/ui/carousel"
import { Card, CardContent } from '../ui/card';


const steps = [
    {
        title: "Create a Project",
        description: "Start by creating a new project to group related feature flags. Easily invite team members to collaborate.",
    },
    {
        title: "Add Feature Flags",
        description: "Define and create feature flags for different parts of your application. Toggle features on or off with a click.",
    },
    {
        title: "Use the API",
        description: "Use our flexible API to access your feature flags directly in your code, making it easy to manage features dynamically.",
    },
    {
        title: "Instant Updates",
        description: "Update your feature flags instantly across your projects, no redeployment needed.",
    },
];

const Steps = ()=> {
    return (
        <div className='grid my-[4rem] gap-6 grid-cols-1 backdrop-blur-md lg:grid-cols-2 xl:grid-cols-4'>
            {steps.map((step, index)=>{
                return(
                    <div className='p-6 h-[350px] relative cursor-pointer hover:bg-primary/10 trans  bg-secondary/30 rounded-lg flex items-center justify-center flex-col gap-3'>
                        <div className='absolute text-[20rem] text-primary/10 z-0  right-0'>{index+1}</div>
                        <span className='text-3xl '>Step <b>{index+1}.</b></span>
                        <p className='text-center'>
                            {step.description}
                        </p>
                    </div>
                )
            })}
        </div>
      )
}


export default Steps;
