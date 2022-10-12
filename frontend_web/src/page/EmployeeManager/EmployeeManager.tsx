import { Form, Input, Select, Popconfirm, Table, Typography, Button } from 'antd'
import React, { ReactNode, useState } from 'react'
import dateHandle from '../../method/dateHandle'
import EditableCell from './EditableCell'
import { EmpManagerCellType } from '../../type'
import { SearchOutlined } from '@ant-design/icons'

const originData: EmpManagerCellType[] = []
for (let i = 0; i < 100; i++) {
    originData.push({
        empIdx: i,
        key: i.toString(),
        name: `Edrward ${i}`,
        onWorkDay: dateHandle.formatDate(new Date()),
        workState: '在職',
        banch: '保育組',
        permession: '一般職員'
    })
}

const EmployeeManager = (): JSX.Element => {
    const [form] = Form.useForm()
    const [data, setData] = useState(originData)
    const [editingKey, setEditingKey] = useState('')

    const isEditing = (record: EmpManagerCellType): any => record.key === editingKey

    const edit = (record: Partial<EmpManagerCellType> & { key: React.Key }): void => {
        form.setFieldsValue({ name: '', ...record })
        setEditingKey(record.key)
    }

    const cancel = (): void => {
        setEditingKey('')
    }

    const save = async (key: React.Key): Promise<void> => {
        try {
            const row = (await form.validateFields()) as EmpManagerCellType

            const newData = [...data]
            const index = newData.findIndex(item => key === item.key)
            if (index > -1) {
                const item = newData[index]
                newData.splice(index, 1, {
                    ...item,
                    ...row
                })
                setData(newData)
                setEditingKey('')
            } else {
                newData.push(row)
                setData(newData)
                setEditingKey('')
            }
        } catch (errInfo) {
            console.log('Validate Failed:', errInfo)
        }
    }

    const columns = [
        {
            title: '員工編號',
            dataIndex: 'empIdx',
            editable: true,
            width: '10%',
            inputCop: (record: EmpManagerCellType): ReactNode => {
                return (
                    <Input disabled placeholder={`${record.empIdx}`} />
                )
            }
        },
        {
            title: '組別',
            dataIndex: 'banch',
            editable: true,
            inputCop: (record: EmpManagerCellType): ReactNode => {
                return (
                    <Select defaultValue={record.banch} >
                        <Select.Option value={'保育組'} key={0}>保育組</Select.Option>
                        <Select.Option value={'公關組'} key={1}>公關組</Select.Option>
                    </Select>
                )
            }
        },
        {
            title: '姓名',
            dataIndex: 'name',
            width: '10%',
            editable: true,
            inputCop: (record: EmpManagerCellType): ReactNode => {
                return (
                    <Input placeholder={`${record.empIdx}`} />
                )
            }
        },
        {
            title: '入職日',
            dataIndex: 'onWorkDay',
            editable: true,
            inputCop: (record: EmpManagerCellType): ReactNode => {
                return (
                    <input type='date'/>
                )
            }
        },
        {
            title: '權限',
            dataIndex: 'permession',
            editable: true,
            inputCop: (record: EmpManagerCellType): ReactNode => {
                return (
                    <Select defaultValue={record.permession}>
                        <Select.Option value={0} key={0}>管理員</Select.Option>
                        <Select.Option value={1} key={1}>主管</Select.Option>
                        <Select.Option value={2} key={1}>一般職員</Select.Option>
                    </Select>
                )
            }
        },
        {
            title: '狀態',
            dataIndex: 'workState',
            editable: true,
            inputCop: (record: EmpManagerCellType): ReactNode => {
                return (
                    <Select defaultValue={record.workState}>
                        <Select.Option value={'在職'} key={0}>在職</Select.Option>
                        <Select.Option value={'離職'} key={1}>離職</Select.Option>
                    </Select>
                )
            }
        },
        {
            title: '',
            dataIndex: 'operation',
            render: (_: any, record: EmpManagerCellType) => {
                const editable = isEditing(record)
                return editable
                    ? (
                        <span>
                            <Typography.Link onClick={async () => await save(record.key)} style={{ marginRight: 8 }}>
                            儲存
                            </Typography.Link>
                            <Popconfirm title="確定取消嗎？" onConfirm={cancel}>
                                <a>取消</a>
                            </Popconfirm>
                        </span>
                    )
                    : (
                        <Typography.Link disabled={editingKey !== ''} onClick={() => edit(record)}>
                            編輯
                        </Typography.Link>
                    )
            },
            inputCop: (record: EmpManagerCellType): ReactNode => {
                return (
                    <Input disabled placeholder={`${record.empIdx}`} />
                )
            }
        }
    ]

    const mergedColumns = columns.map(col => {
        if (!col.editable) {
            return col
        }
        return {
            ...col,
            onCell: (record: EmpManagerCellType) => ({
                record,
                inputType: col.inputCop(record),
                dataIndex: col.dataIndex,
                title: col.title,
                editing: isEditing(record)
            })
        }
    })

    return (
        <Form form={form} component={false}>
            <div className={styles.empMangerFilter}>
                <div>xx組</div>
                <div>
                    <Input style={{ width: '150px', marginRight: '0.4rem' }} prefix={<SearchOutlined />} placeholder={'請輸入姓名'} />
                    <Select defaultValue={'保育組'} >
                        <Select.Option value={'保育組'} key={0}>保育組</Select.Option>
                        <Select.Option value={'公關組'} key={1}>公關組</Select.Option>
                    </Select>
                    <Select defaultValue={'在職'}>
                        <Select.Option value={'在職'} key={0}>在職</Select.Option>
                        <Select.Option value={'離職'} key={1}>離職</Select.Option>
                    </Select>
                    <Button>搜尋</Button>
                </div>
            </div>
            <div style={{ maxHeight: '80%', overflow: 'scroll' }}>
                <Table
                    components={{
                        body: {
                            cell: EditableCell
                        }
                    }}
                    sticky
                    showHeader
                    bordered
                    dataSource={data}
                    columns={mergedColumns}
                    rowClassName="editable-row"
                    pagination={false}
                />
            </div>
        </Form>
    )
}

export default EmployeeManager
