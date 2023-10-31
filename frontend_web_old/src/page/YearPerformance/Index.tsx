import { Table } from 'antd'
import useReduceing from 'Hook/useReducing'
import React from 'react'
import SearchBar from './component/SearchBar'

const columns = [
    {
        dataIndex: 'UserName',
        key: 1,
        title: '姓名'
    },
    {
        dataIndex: 'Year',
        key: 2,
        title: '年度'
    },
    {
        dataIndex: 'Avg',
        key: 3,
        title: '平均分數'
    }
]

const Index = (): JSX.Element => {
    const { company } = useReduceing()
    return (
        <>
            <SearchBar />
            <Table
                dataSource={company.yearPerformance}
                columns={columns}
            />
        </>
    )
}
export default Index
