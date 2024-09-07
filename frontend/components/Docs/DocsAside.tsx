import { ArrowLeftToLine } from "lucide-react";
import React, { ReactNode } from "react";

interface props {
    className ?: string,
    children: ReactNode
    open : boolean
    setOpen :React.Dispatch<React.SetStateAction<boolean>>
}



const Aside: React.FC<props> = ({ children, className, ...props }) => {
  return (
    <aside className={`${!props.open ? "hidden":"block bg-primary-foreground/20 backdrop-blur-xl top-[4rem] w-full"} absolute lg:relative lg:block h-full left-0 w-64 font-medium text-primary/60 text-lg p-4 space-y-4  lg:w-80 ${className}`}{...props}>
      <ArrowLeftToLine className = {`block lg:hidden`} onClick={()=>props.setOpen(false)} />
      {children}
    </aside>
  );
};

export default Aside;
