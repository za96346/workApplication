import { Table } from 'antd'
import React, { useEffect, useMemo } from 'react'
import columns from './methods/column'
// import modalType from 'static'
import api from 'api/Index'
import Searchbar from './components/Searchbar'
import { useAppSelector } from 'hook/redux'
import { v4 } from 'uuid'
import Btn from 'shared/Button'
import Dropdown from 'shared/Dropdown'

const Index = (): JSX.Element => {
    const employee = useAppSelector((v) => v?.user?.employee)
    const dataSource = useMemo(() => {
        return employee?.map((item) => ({
            ...item,
            key: v4(),
            action: <Dropdown menu={[]}/>
        }))
    }, [employee])

    useEffect(() => {
        void api.user.getEmployee()
    }, [])
    return (
        <>
            <Btn.Add/>
            <Searchbar/>
            <Table
                dataSource={dataSource}
                columns={columns}
            />
        </>
    )
}

export default Index
