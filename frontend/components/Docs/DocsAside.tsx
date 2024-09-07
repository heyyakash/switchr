import React, { ReactNode } from "react";

interface props {
    className ?: string,
    children: ReactNode
}



const Aside: React.FC<props> = ({ children, className, ...props }) => {
  return (
    <aside className={` h-full left-0 w-64 font-medium text-primary/60 text-lg p-4 space-y-4  lg:w-80 ${className}`}{...props}>
      {children}
    </aside>
  );
};

export default Aside;
