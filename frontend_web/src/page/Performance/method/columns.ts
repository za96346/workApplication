import type { ColumnsType } from 'antd/es/table'
import { performanceType } from 'type'

export const columns: ColumnsType<performanceType> = [
    {
        key: 'UserName',
        title: '姓名',
        dataIndex: 'UserName',
        width: '80px'
    },
    {
        key: 'Year',
        title: '年份',
        dataIndex: 'Year',
        width: '60px'
    },
    {
        key: 'Month',
        title: '月份',
        dataIndex: 'Month',
        width: '60px'
    },
    {
        key: 'BanchName',
        title: '組別',
        dataIndex: 'BanchName',
        width: '90px',
        responsive: ['md']
    },
    {
        key: 'Goal',
        title: '年度目標',
        dataIndex: 'Goal',
        responsive: ['md']
    },
    {
        key: 'Attitude',
        title: '態度',
        dataIndex: 'Attitude',
        width: '50px'
    },
    {
        key: 'Efficiency',
        title: '效率',
        dataIndex: 'Efficiency',
        width: '50px'
    },
    {
        key: 'Professional',
        title: '專業',
        dataIndex: 'Professional',
        width: '50px'
    },
    {
        key: 'BeLate',
        title: '遲到/早退',
        dataIndex: 'BeLate',
        width: '50px',
        responsive: ['md']
    },
    {
        key: 'DayOffNotOnRule',
        title: '未依規定請假',
        dataIndex: 'DayOffNotOnRule',
        width: '50px',
        responsive: ['md']
    },
    {
        key: 'Directions',
        title: '績效描述',
        dataIndex: 'Directions',
        responsive: ['md']
    },
    {
        key: 'action',
        title: '',
        dataIndex: 'action',
        width: '100px'
    }
]
