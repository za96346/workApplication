
interface init {
    dataIndex: string
    title: string
    key: string
    fixed?: any
    width?: any
}
export const column: init[] = [
    {
        dataIndex: 'UserName',
        title: '姓名',
        key: 'UserName',
        fixed: 'left',
        width: '80px'
    },
    {
        dataIndex: 'EmployeeNumber',
        title: '員工編號',
        key: 'EmployeeNumber'
    },
    {
        dataIndex: 'BanchName',
        title: '部門',
        key: 'BanchName'
    },
    {
        dataIndex: 'Permession',
        title: '權限',
        key: 'Permession'
    },
    {
        dataIndex: 'OnWorkDay',
        title: '到職日',
        key: 'OnWorkDay'
    },
    {
        dataIndex: 'action',
        title: '',
        key: 'action'
    }
]
