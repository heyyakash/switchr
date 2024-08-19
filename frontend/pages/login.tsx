import Login from "@/components/Auth/Login"
import SignUp from "@/components/Auth/Signup"
import { ModeToggle } from "@/components/Toggle"
import { Button } from "@/components/ui/button"
import { Key } from "lucide-react"
import { useState } from "react"

const Auth = () => {
    const [mode, setMode] = useState<'Login' | 'Sign Up'>("Sign Up")
    const changeMode = async () => {
        if (mode === "Login") {
            setMode("Sign Up")
        } else {
            setMode("Login")
        }
    }

    return (
        <div className="grid grid-cols-1 xl:grid-cols-2 grid-rows-1 font-inter  h-screen w-full">
            <div className="w-full h-full  hidden xl:flex p-10 justify-between flex-col dark:bg-white/10 bg-black/5 ">
                <div className="flex gap-2 items-center text-[1.5rem] font-medium ">
                    <Key size={"40px"} /> Switchr
                </div>
                <div>
                    <blockquote className="text-2xl font-medium">
                    Writing is the best way to capture fleeting thoughts and immortalize profound moments. As the ink flows, so does the essence of our experiences.
                    </blockquote>
                    <footer className="font-medium mt-3">- Unknown (ChatGPT probably)</footer>

                </div>
            </div>
            <div className="w-full h-full relative grid place-items-center ">
                <div className="absolute top-5 right-5 md:top-10 md:right-10 flex gap-3">
                    <Button onClick={() => changeMode()} className="text-semibold" variant={"outline"}>{mode}</Button>
                    <ModeToggle />
                    
                </div>
                {mode === "Sign Up" ? <Login /> : <SignUp />}
            </div>
        </div>
    )

}


export default Auth


Auth.getLayout = () => {
    return (
        <></>
    )
}