import { BanchType, UserType } from 'type'
import React from 'react'
import Btn from 'Share/Btn'

export const columns: any = [
    {
        title: '部門名稱',
        dataIndex: 'BanchName',
        key: '1',
        align: 'center'
    },
    {
        title: '人數',
        dataIndex: 'count',
        key: '2',
        align: 'center'
    },
    {
        title: '',
        dataIndex: 'action',
        key: '3',
        align: 'center'
    }
]
export const dataSource = (
    v: BanchType[],
    onEdit: Function,
    emp: UserType[],
    onDelete: Function
): any =>
    v.map((item) => {
        const res = emp?.filter((i) => i.Banch === item.Id)?.length || 0
        return {
            ...item,
            action: <>
                <Btn.Edit onClick={() => { onEdit(item) }} />
                <Btn.Delete onClick={() => { onDelete(item) }} />
            </>,
            count: res
        }
    })
