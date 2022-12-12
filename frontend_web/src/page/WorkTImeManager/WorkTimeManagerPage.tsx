import React, { useEffect, useState } from 'react'
import api from '../../api/api'
import useReduceing from '../../Hook/useReducing'
import Input from '../../component/Input'
import { Button } from '../../component/Button'
import { Table } from 'antd'
import { workTimeType } from '../../type'
import Modal from './Modal'
import Btn from '../../component/Btn'

const columns = [
    {
        key: 'user',
        dataIndex: 'user',
        title: '員工姓名'
    },
    {
        key: 'year',
        dataIndex: 'year',
        title: '年份'
    },
    {
        key: 'month',
        dataIndex: 'month',
        title: '月份'
    },
    {
        key: 'workHours',
        dataIndex: 'workHours',
        title: '總工作時數'
    },
    {
        key: 'timeOff',
        dataIndex: 'timeOff',
        title: '總休息時數'
    },
    {
        key: 'usePaidVocation',
        dataIndex: 'usePaidVocation',
        title: '以使用特休天數'
    },
    {
        key: 'action',
        dataIndex: 'action',
        title: '編輯'
    }
]

const dataSource = (arr: workTimeType[]): any[] => {
    return arr.map((item) => {
        return {
            user: item.UserId,
            year: item.Year,
            month: item.Month,
            workHours: item.WorkHours,
            timeOff: item.TimeOff,
            usePaidVocation: item.UsePaidVocation,
            action: (
                <>
                    <Btn.Edit onClick={() => {}}/>
                    <Btn.Delete onClick={() => {}}/>
                </>
            )
        }
    })
}

const WorkTimeManagerPage = (): JSX.Element => {
    const [open, setOpen] = useState(false)
    const { company } = useReduceing()
    const onClose = (): void => {
        setOpen(false)
    }
    useEffect(() => {
        api.getWorkTime(2022, 5, null)
    }, [])
    console.log(company.workTime)
    return (
        <>
            <Modal
                onCancel={onClose}
                onOk={onClose}
                open={open}
            />
            <div className={styles.workTimeBlock}>
                <Button onClick={() => { setOpen(true) }} className="btn btn-secondary" text="新增"/>
                <div className="row">
                    <Input className="col-md-4" title='年份' />
                    <Input className="col-md-3" title='月份' />
                    <Input className="col-md-3" title='年份' />
                    <Button text="搜尋" className="col-md-2 btn btn-secondary" onClick={() => {}} />
                </div>
                <Table
                    dataSource={dataSource(company.workTime)}
                    columns={columns}
                />

            </div>
        </>
    )
}
export default WorkTimeManagerPage
