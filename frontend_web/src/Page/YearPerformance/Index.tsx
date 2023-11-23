import { Table } from 'antd'
import React, { useEffect } from 'react'
import columns from './methods/column'
import api from 'api/Index'
import Searchbar from './components/Searchbar'
import { useAppSelector } from 'hook/redux'

const Index = (): JSX.Element => {
    const performance = useAppSelector((v) => v?.performance?.year)

    useEffect(() => {
        void api.performance.getYear()
    }, [])
    return (
        <>
            <Searchbar/>
            <Table
                dataSource={performance}
                columns={columns}
                sticky={{ offsetHeader: -20 }}
                style={{
                    fontSize: '0.5rem'
                    // width: 'fit-content'
                }}
                size='small'
            />
        </>
    )
}

export default Index
