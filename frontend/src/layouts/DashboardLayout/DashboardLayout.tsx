import { Grid } from '@mantine/core'
import React, { type JSX } from 'react'
import { Outlet } from 'react-router'
import Sidebar from '../../components/ui/Sidebar'
import Header from '../../components/ui/Header'
import Footer from '../../components/ui/Footer'
const DashboardLayout: React.FC = (): JSX.Element => {
    return (
        <main>
            <Header />
            <Grid>
                <Grid.Col span={1}>
                    <Sidebar />
                </Grid.Col>
                <Grid.Col span={11}>
                    <Outlet />
                </Grid.Col>
            </Grid>
            <Footer />
        </main>
    );
};

export default DashboardLayout