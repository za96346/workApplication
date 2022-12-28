
interface init {
    dataIndex: string
    title: string
    key: string
}
export const column: init[] = [
    {
        dataIndex: 'EmployeeNumber',
        title: '員工編號',
        key: 'EmployeeNumber'
    },
    {
        dataIndex: 'UserName',
        title: '姓名',
        key: 'UserName'
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
