import { headers } from "next/headers"

interface responseInterface {
    response: any,
    status: number
}

export async function HTTPRequest(
    endpoint :string,
    options : any,
    method:string
):Promise<responseInterface>{
    const base_url = process.env.NEXT_PUBLIC_BASE_URL as string
    const req = await fetch(base_url + endpoint,{
        method,
        credentials: "include",
        ...options
    })
    const result = await req.json()
    return {
        response:result,
        status: req.status
    }
}