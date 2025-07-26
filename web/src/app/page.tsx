'use client'
import React, { useState, useEffect } from 'react'
import {
    Grid,
    Column,
    Tile,
    DataTable,
    Table,
    TableHead,
    TableRow,
    TableHeader,
    TableBody,
    TableCell,
    Tag,
    ProgressBar,
    InlineLoading,
    Button
} from '@carbon/react'
import {Play, Stop, Restart, AddAlt, FlowData} from '@carbon/icons-react'

interface SystemStatus {
    status: string
    modules: number
    atoms: number
    memoryUsage: number
    cpuUsage: number
}

interface ModuleStatus {
    id: string
    name: string
    status: 'running' | 'stopped' | 'error'
    lastRun: string
    nextRun?: string
}

export default function Dashboard() {
    const [systemStatus, setSystemStatus] = useState<SystemStatus | null>(null)
    const [modules, setModules] = useState<ModuleStatus[]>([])
    const [loading, setLoading] = useState(true)

    useEffect(() => {
        const fetchData = async () => {
            try {
                // Simulate API calls
                await new Promise(resolve => setTimeout(resolve, 1000))

                setSystemStatus({
                    status: 'running',
                    modules: 5,
                    atoms: 23,
                    memoryUsage: 45,
                    cpuUsage: 23
                })

                setModules([
                    { id: '1', name: 'Email Auto-Responder', status: 'running', lastRun: '2025-01-15 10:30:00', nextRun: '2025-01-15 11:00:00' },
                    { id: '2', name: 'Content Pipeline', status: 'running', lastRun: '2025-01-15 10:25:00' },
                    { id: '3', name: 'System Monitor', status: 'stopped', lastRun: '2025-01-15 09:45:00' },
                    { id: '4', name: 'Git Backup', status: 'error', lastRun: '2025-01-15 09:30:00' },
                    { id: '5', name: 'Document Processor', status: 'running', lastRun: '2025-01-15 10:35:00', nextRun: '2025-01-15 14:00:00' }
                ])
            } catch (error) {
                console.error('Failed to fetch data:', error)
            } finally {
                setLoading(false)
            }
        }

        fetchData()
    }, [])

    const getStatusTag = (status: string) => {
        switch (status) {
            case 'running': return <Tag type="green">Running</Tag>
            case 'stopped': return <Tag type="gray">Stopped</Tag>
            case 'error': return <Tag type="red">Error</Tag>
            default: return <Tag>{status}</Tag>
        }
    }

    const moduleRows = modules.map((module) => ({
        id: module.id,
        name: module.name,
        status: getStatusTag(module.status),
        lastRun: module.lastRun,
        nextRun: module.nextRun || 'On-demand',
        actions: (
            <div className="flex gap-2">
                <Button kind="ghost" size="sm" renderIcon={Play} iconDescription="Start" />
                <Button kind="ghost" size="sm" renderIcon={Stop} iconDescription="Stop" />
                <Button kind="ghost" size="sm" renderIcon={Restart} iconDescription="Restart" />
            </div>
        )
    }))

    const moduleHeaders = [
        { key: 'name', header: 'Module Name' },
        { key: 'status', header: 'Status' },
        { key: 'lastRun', header: 'Last Run' },
        { key: 'nextRun', header: 'Next Run' },
        { key: 'actions', header: 'Actions' }
    ]

    if (loading) {
        return (
            <div className="p-8">
                <InlineLoading description="Loading dashboard..." />
            </div>
        )
    }

    return (
        <div>Hello World </div>
    )
}