`use client`
import React from 'react';
import { Header, HeaderName, HeaderGlobalBar, HeaderGlobalAction } from '@carbon/react';
import { Settings, User } from '@carbon/icons-react';

export default function AppShell({
                                     children,
                                 }: {
    children: React.ReactNode
}) {
    return (
        <div>
            {children}
        </div>
    )
}