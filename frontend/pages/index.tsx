import FeatureBox from "@/components/LandingPage/FeatureBox"
import { ModeToggle } from "@/components/Toggle"
import { Button } from "@/components/ui/button"
import { Key } from "lucide-react"
import { useState } from "react"

const Home = () =>  {
  const [mainBg, setMainBg] = useState("from-yellow-300 to-green-400")

  return (
    <>
    <nav className="w-full h-[70px] fixed top-0 backdrop-blur-sm  border-b-[1px]">
      <div className="max-w-[1200px] w-full py-4 mx-auto flex items-center justify-between">
        <div className="flex gap-2 items-center">
          <Key size={27}/>
          <p className="text-xl font-semibold">Switchr</p>
        </div>
        <div className="flex items-center gap-3">
          <Button variant={"secondary"}>API Docs</Button>
          <Button>Login</Button>
          <ModeToggle />
        </div>
      </div>
    </nav>
    <section className="w-full py-[6rem] bg-[url('/pattern.png')]">
      <div className=" max-w-[1200px] w-full min-h-[500px] mx-auto flex flex-col items-center justify-center gap-[3rem]">
        <h1 className={`bg-gradient-to-r ${mainBg} bg-clip-text text-transparent p-4 text-6xl font-extrabold text-center leading-[3.5rem] transition-all duration-500`}>
        Supercharge Your Development <br /> with Dynamic Feature Flags
        </h1>
        {/* <p className="w-[600px] text-lg text-center text-primary/80">Switchr is a powerful and dynamic feature flag management tool that lets you control your features in real-time. Seamlessly manage, toggle, and test features without redeploying your code.</p> */}
        <div className="w-full mt-2 grid grid-cols-3 gap-3">
          <FeatureBox serial={1} heading="Redis-Backed Performance" desc="Our Redis-powered infrastructure ensures that your feature flags are stored and accessed with lightning-fast performance, even at scale." setMainBg={setMainBg} />
          <FeatureBox serial={2} heading="Team Collaboration" desc="Manage your development teams effortlessly. Add or remove members, assign roles, and control permissions within each project." setMainBg={setMainBg} />
          <FeatureBox serial={3} heading="API-First Design" desc="Switchr's robust API allows you to programmatically access your feature flags. Easily integrate it into any application stack." setMainBg={setMainBg} />
        </div>
        <Button className="bg-green-400" variant={"default"}>Enter App</Button>
      </div>

    </section> 

    <section className="w-full py-[3rem]">

    <div className=" max-w-[1200px] w-full mx-auto">
        <h2 className="text-3xl font-semibold">Why Choose Switchr?</h2>
      </div>
    </section>
    </>
  )  
}

export default Home

Home.getLayout = () => {
  return <></>
}