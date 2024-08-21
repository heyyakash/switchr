import { headers } from "next/headers"

export interface responseInterface {
    response: {
        message:any,
        success:boolean
    },
    status: number
}

export interface CustomError extends Error {
    status?: number;
    message: string;
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
    if(req.status!==200){
        const error : CustomError = new Error(result?.message || "Unkown Error")
        error.status = req.status
        throw error
    }
    return {
        response:result,
        status: req.status
    }
}