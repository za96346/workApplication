import type { ColumnsType } from 'antd/es/table'

const columns: ColumnsType<any> = [
    {
        key: 'UserName',
        title: '姓名',
        dataIndex: 'UserName',
        fixed: 'left'
    },
    {
        key: 'Year',
        title: '年份',
        dataIndex: 'Year',
        fixed: 'left'
    },
    {
        key: 'Score',
        title: '分數',
        dataIndex: 'Score'
    }
]
export default columns
