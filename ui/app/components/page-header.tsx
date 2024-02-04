import { Divider } from "@nextui-org/react";

export default function PageHeader({children}) {
    return (
        <div style={{ width: "100%", marginBottom:20, paddingLeft:5 }}>
            <h1 className="relative inline-flex items-center tap-highlight-transparent outline-none data-[focus-visible=true]:z-10 
      data-[focus-visible=true]:outline-2 data-[focus-visible=true]:outline-focus data-[focus-visible=true]:outline-offset-2 
      text-maximum text-foreground no-underline hover:opacity-80 active:opacity-disabled transition-opacity"
                style={{ fontSize: "var(--nextui-font-size-medium)" }}>
                {children}
                </h1>
            
        </div>
    )
}