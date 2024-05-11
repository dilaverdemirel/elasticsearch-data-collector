import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { Providers } from "./providers";
import AppNavBar from "./navbar";
import 'bootstrap/dist/css/bootstrap.min.css';
import 'react-toastify/dist/ReactToastify.css';
import { ToastContainer } from "react-toastify";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Elasticsearch Data Collector",
  description: "Elasticsearch Data Collector",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={inter.className}>

        <AppNavBar />

        <Providers>
          <div className="flex z-40 w-full h-auto items-center justify-center data-[menu-open=true]:border-none sticky top-0 inset-x-0 backdrop-blur-lg data-[menu-open=true]:backdrop-blur-xl backdrop-saturate-150 bg-background/70"
            style={{ paddingTop: 20 }}>
            <div className="z-40 flex px-6 gap-4 w-full flex-row relative flex-nowrap items-center justify-between h-[var(--navbar-height)] max-w-[1024px]">
              {children}
            </div>
          </div>
          <ToastContainer />
        </Providers>

        <div style={{width:'100%', textAlign: "center", marginTop:30, marginBottom:10}}>
          <hr/>
          <i>Elasticsearch Data Collector</i>
          </div>

      </body>
    </html>
  );
}
