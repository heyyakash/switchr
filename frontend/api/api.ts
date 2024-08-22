import Router from "next/router"


export interface responseInterface {
    response: {
        message:any,
        success:boolean
    },
    status: number
}

export async function HTTPRequest(
    endpoint :string,
    options : any,
    method:string
):Promise<responseInterface | null>{
    try{
        const base_url = process.env.NEXT_PUBLIC_BASE_URL as string
        const req = await fetch(base_url + endpoint,{
            method,
            credentials: "include",
            ...options
        })
        console.log(req.status)
        if (req.status === 302){
            console.log("Should've redirected")
           Router.push('/login')
           return null
        }
        const result = await req.json()
        return {
            response:result,
            status: req.status
        }
    }catch(err){
        console.log(err)
        return null
    }
}