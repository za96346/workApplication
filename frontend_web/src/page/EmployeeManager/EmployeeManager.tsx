import { Form, Input, Select, Popconfirm, Table, Typography, Button } from 'antd'
import React, { ReactNode, useEffect, useState } from 'react'
import dateHandle from '../../method/dateHandle'
import EditableCell from './EditableCell'
import { BanchType, EmpManagerCellType, UserType } from '../../type'
import { SearchOutlined } from '@ant-design/icons'
import api from '../../api/api'
import BanchSelector from '../../component/BanchSelector'
import PermessionSelector from '../../component/PermessionSelector'
import statics from '../../statics'
import StatusSelector from '../../component/StatusSelector'
import useReduceing from '../../Hook/useReducing'

const items = (value: UserType[], banch: BanchType[]): any => {
    return value?.map((i, index) => {
        const b = banch?.find((item) => item.Id === i.Banch)
        return {
            EmployeeNumber: i.EmployeeNumber,
            key: i.UserId.toString(),
            UserName: i.UserName,
            OnWorkDay: dateHandle.formatDate(new Date(i.OnWorkDay)),
            WorkState: i.WorkState === 'on' ? '在職' : '離職',
            Banch: b?.BanchName,
            BanchId: i.Banch,
            Permession: statics.permession[i.Permession],
            PermessionId: i.Permession
        }
    })
}

const EmployeeManager = (): JSX.Element => {
    const [form] = Form.useForm()
    const [data, setData] = useState([])
    const [editingKey, setEditingKey] = useState('')
    const { company, loading } = useReduceing()

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
            console.log(row)
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
            dataIndex: 'EmployeeNumber',
            editable: true,
            width: '10%',
            inputCop: (record: EmpManagerCellType): ReactNode => {
                return (
                    <Input disabled placeholder={`${record.EmployeeNumber}`} />
                )
            }
        },
        {
            title: '組別',
            dataIndex: 'Banch',
            editable: true,
            inputCop: (record: EmpManagerCellType): ReactNode => {
                return (
                    <BanchSelector defaultValue={record.BanchId} />
                )
            }
        },
        {
            title: '姓名',
            dataIndex: 'UserName',
            width: '10%',
            editable: true,
            inputCop: (record: EmpManagerCellType): ReactNode => {
                return (
                    <Input disabled placeholder={`${record.UserName}`} />
                )
            }
        },
        {
            title: '入職日',
            dataIndex: 'OnWorkDay',
            editable: true,
            inputCop: (record: EmpManagerCellType): ReactNode => {
                return (
                    <input type='date'/>
                )
            }
        },
        {
            title: '權限',
            dataIndex: 'Permession',
            editable: true,
            inputCop: (record: EmpManagerCellType): ReactNode => {
                return (
                    <PermessionSelector defaultValue={record.PermessionId} />
                )
            }
        },
        {
            title: '狀態',
            dataIndex: 'WorkState',
            editable: true,
            inputCop: (record: EmpManagerCellType): ReactNode => {
                return (
                    <StatusSelector defaultValue={record.WorkState} />
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
                    <Input disabled placeholder={`${record.EmployeeNumber}`} />
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

    useEffect(() => {
        void api.getUserAll()
    }, [])

    useEffect(() => {
        setData(items(company.employee, company.banch))
    }, [company.employee, company.banch])

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
                    loading={loading.onFetchUserAll}
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
