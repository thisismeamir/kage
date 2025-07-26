import type { Metadata } from 'next'
import { IBM_Plex_Sans } from 'next/font/google'
import './globals.css'
import AppShell from "@/app/components/AppShell";


const plex = IBM_Plex_Sans({ subsets: ['latin'] })

export const metadata: Metadata = {
    title: 'Kage - Autonomous Graph Execution',
    description: 'Kernel of Autonomous Graph-based Execution',
}

export default function RootLayout({
                                       children,
                                   }: {
    children: React.ReactNode
}) {
    return (
        <html lang="en">
        <body className={plex.className}>
        <AppShell>{children}</AppShell>
        </body>
        </html>
    )
}