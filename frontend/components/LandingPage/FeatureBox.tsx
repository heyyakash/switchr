import React, { SetStateAction, useState } from 'react'

interface props {
    serial: number
    heading: string
    desc: string
    setMainBg: React.Dispatch<SetStateAction<string>>
}

const FeatureBox: React.FC<props> = ({ serial, heading, desc, setMainBg }) => {

    const bg_list = {
        0: "from-yellow-300 to-green-400",
        1: "from-pink-400 to-rose-400",
        2: "from-cyan-200 to-blue-200",
        3: "from-violet-200 to-pink-200",
    }
    const [headingColor, setHeadingColor] = useState("bg-primary")
    const changeMainBg = (num: keyof typeof bg_list) => {
        if(num==0){
            setHeadingColor("bg-primary")
        }else{
            setHeadingColor(bg_list[num])
        }
        setMainBg(bg_list[num])
    }
    return (
        <div onMouseEnter={() => changeMainBg(serial as keyof typeof bg_list)} onMouseLeave={() => changeMainBg(0)} className="hover:bg-primary/10 hover:backdrop-blur-xl trans cursor-pointer backdrop-blur-sm p-6 bg-secondary/30 rounded-xl">
            <div className="flex items-center gap-3 ">
                <div className={`border-2 h-10 w-10 flex items-center justify-center rounded-full`}>{serial}</div>
                <h3 className={`text-xl font-semibold bg-gradient-to-r ${headingColor} text-transparent bg-clip-text`}>{heading}</h3>
            </div>
            <div className="mt-4 text-primary/70">{desc}</div>
        </div>
    )
}

export default FeatureBox