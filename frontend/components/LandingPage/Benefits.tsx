import { ChartNoAxesCombined, Flag, Handshake, SlidersVertical, Zap } from 'lucide-react'
import React from 'react'

const Benefits = () => {
    return (
        <div className='grid grid-cols-1 place-items-center lg:grid-cols-3 mt-10 gap-6'>

            <div className='bg-secondary/30 p-6 rounded-lg flex flex-col  gap-3 hover:bg-primary/10 trans cursor-pointer'>
                <Zap size={60} className="text-yellow-300  mb-4" />
                <h2 className='text-2xl font-semibold'>
                    Speed Up Development
                </h2>
                <p className='text-primary/70'>Deploy faster by decoupling code deployment from feature release</p>
            </div>
            <div className='bg-secondary/30 p-6 rounded-lg flex flex-col  gap-3 hover:bg-primary/10 trans cursor-pointer'>
                <Handshake  size={60} className="text-blue-300 mb-4" />
                <h2 className='text-2xl font-semibold'>
                Improve Collaboration
                </h2>
                <p className='text-primary/70'>Deploy faster by decoupling code deployment from feature release</p>
            </div>
            <div className='bg-secondary/30 p-6 rounded-lg flex flex-col  gap-3 hover:bg-primary/10 trans cursor-pointer'>
            <SlidersVertical size={60} className="text-slate-300  mb-4" />
                <h2 className='text-2xl font-semibold'>
                Flexibility and Control
                </h2>
                <p className='text-primary/70'>Toggle features on or off instantly to test, manage, or roll back releases</p>
            </div>
            <div className='bg-secondary/30 p-6 rounded-lg flex flex-col  gap-3 hover:bg-primary/10 trans cursor-pointer'>
            <ChartNoAxesCombined size={60} className="text-green-300" />
                <h2 className='text-2xl font-semibold'>
                Scale Effortlessly
                </h2>
                <p className='text-primary/70'>Switchr's Redis-backed architecture ensures fast and reliable performance, no matter the scale</p>
            </div>
            <div className='bg-secondary/30 p-6 rounded-lg flex flex-col  gap-3 hover:bg-primary/10 trans cursor-pointer'>
                <Flag size={60} className="text-orange-300" />
                <h2 className='text-2xl font-semibold'>
                    Easy Flag Management
                </h2>
                <p className='text-primary/70'>Effortlessly create, update, and delete feature flags with full CRUD support, directly from the dashboard.</p>
            </div>

        </div>
    )
}

export default Benefits