import React, { type JSX } from 'react'
import LineChart from './components/LineChart'

import { dumbData } from './constants/dumbData'

const Dashboard: React.FC = (): JSX.Element => {
    return (
        <section>
            <h1>Dashboard Page</h1>
            <LineChart />
        </section>
    )
}

export default Dashboard