import { Table } from 'antd'
import React, { useEffect, useMemo } from 'react'
import columns from './methods/column'
import modalType from 'static'
import api from 'api/Index'

const Index = (): JSX.Element => {
    const dataSource = useMemo(() => {
        return []
    }, [])

    useEffect(() => {
        api.
    }, [])
    return (
        <>
            <Table
                dataSource={dataSource}
                columns={columns}
            />
        </>
    )
}

export default Index
