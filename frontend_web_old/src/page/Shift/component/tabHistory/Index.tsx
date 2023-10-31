import { Button, Table } from 'antd'
import useReduceing from 'Hook/useReducing'
import { tabHistoryColumns } from 'page/Shift/method/columns'
import React from 'react'
import ShiftTable from '../ShiftTable'
import SearchBar from './SearchBar'

const TabHistory = (): JSX.Element => {
    const { shiftEdit } = useReduceing()

    return (
        <>
            <Button
                onClick={() => {
                    window.open(
                        '/print/shift',
                        '班表',
                        'height=800,width=800'
                    )
                }}
            >
                列印 班表
            </Button>
            <SearchBar />
            <ShiftTable />
            <Table
                columns={tabHistoryColumns}
                dataSource={shiftEdit.total}
            />
        </>
    )
}
export default TabHistory
