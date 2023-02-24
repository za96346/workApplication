import { Table } from 'antd'
import useReduceing from 'Hook/useReducing'
import dateHandle from 'method/dateHandle'
import { tabHistoryColumns } from 'page/Shift/method/columns'
import React, { useMemo } from 'react'
import Head from '../tabEdit/Head'
import Row from '../tabEdit/Row'
import SearchBar from './SearchBar'

const TabHistory = (): JSX.Element => {
    const { shiftEdit } = useReduceing()

    // 日期
    const dayArray = useMemo(() => {
        const start = shiftEdit?.StartDay
        const end = shiftEdit?.EndDay
        const dayContainer = ['']
        for (let i = 0; i < new Date(end).getDate(); i++) {
            const res = dateHandle.addDays(start, i)
            dayContainer.push(res)
        }
        return dayContainer
    }, [shiftEdit?.StartDay, shiftEdit?.EndDay])

    return (
        <>
            <SearchBar />
            <table style={{ cursor: 'pointer' }} className='mb-5 table table-bordered table-hover table-striped table-responsive-md'>
                <thead>
                    <Head dayArray={dayArray}/>
                </thead>
                <tbody>
                    <Row sendMessage={() => {}} dayArray={dayArray} />
                </tbody>
            </table>
            <Table
                columns={tabHistoryColumns}
                dataSource={shiftEdit.total}
            />
        </>
    )
}
export default TabHistory
