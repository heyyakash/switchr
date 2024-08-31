import { HTTPRequest } from "@/api/api"
import Login from "@/components/Auth/Login"
import SignUp from "@/components/Auth/Signup"
import Loading from "@/components/Loading/Loading"
import { ModeToggle } from "@/components/Toggle"
import { Button } from "@/components/ui/button"
import { useQuery } from "@tanstack/react-query"
import { Key } from "lucide-react"
import Link from "next/link"
import { useRouter } from "next/router"
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

    const {data, isLoading} = useQuery({ queryKey: ["user"], queryFn: async () => { return (await HTTPRequest("/user", {}, "GET",false)) }})
    const router = useRouter()


    if (isLoading) return <Loading />
    if(data?.status==200){
        router.push("/dashboard")
    }
    return (
        <div className="grid grid-cols-1 xl:grid-cols-2 grid-rows-1 font-inter  h-screen w-full">
            <div className="w-full h-full  hidden xl:flex p-10 justify-between flex-col bg-[url('/pattern.png')] bg-black/5 ">
                <Link href = "/" className="flex gap-2 items-center text-[1.5rem] font-medium ">
                    <Key size={"40px"} /> Switchr
                </Link>
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