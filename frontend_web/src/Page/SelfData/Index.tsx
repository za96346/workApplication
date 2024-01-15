import { Button, DatePicker, Form, Input, Select } from 'antd'
import api from 'api/Index'
import dayjs from 'dayjs'
import { useAppSelector } from 'hook/redux'
import { usePermission } from 'hook/usePermission'
import useRoleBanchList from 'hook/useRoleBanchUserList'
import React, { useEffect } from 'react'
import Btn from 'shared/Button'
import { FuncCodeEnum, OperationCodeEnum } from 'types/system'
import ModelChangePassword from 'Page/SelfData/ModelChangePassword/Index'

const Index = (): JSX.Element => {
    const [form] = Form.useForm()
    const data = useAppSelector((v) => v?.user?.mine)
    const permission = usePermission({ funcCode: FuncCodeEnum.selfData })

    const rolBanchList = useRoleBanchList({
        funcCode: FuncCodeEnum.selfData,
        operationCode: OperationCodeEnum.edit
    })

    useEffect(() => {
        void api.user.getMine()
    }, [])

    return (
        <div>
            <ModelChangePassword />
            <Form
                disabled={!permission?.isEditable({})}
                onFinish={(v) => {
                    void api.user.updateMine(v).then(() => {
                        void api.user.getMine()
                    })
                }}
                name="validateOnly"
                className='row'
                autoComplete="off"
            >
                <Form.Item
                    className='col-md-6'
                    name="Account"
                    label="帳號"
                    initialValue={data?.Account}
                    rules={[{ required: true }]}
                >
                    <Input disabled />
                </Form.Item>
                <Form.Item
                    className='col-md-6'
                    name="UserName"
                    label="姓名"
                    initialValue={data?.UserName}
                    rules={[{ required: true }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    className='col-md-6'
                    name="EmployeeNumber"
                    label="員工編號"
                    initialValue={data?.EmployeeNumber}
                    rules={[{ required: true }]}
                >
                    <Input disabled />
                </Form.Item>
                <Form.Item
                    className='col-md-6'
                    name="BanchId"
                    label="部門"
                    initialValue={data?.BanchId}
                    rules={[{ required: true }]}
                >
                    <Select
                        options={rolBanchList.banchSelectList}
                        disabled
                    />
                </Form.Item>
                <Form.Item
                    className='col-md-6'
                    name="RoleId"
                    label="角色"
                    initialValue={data?.RoleId}
                    rules={[{ required: true }]}
                >
                    <Select
                        options={rolBanchList.roleSelectList}
                        disabled
                    />
                </Form.Item>
                <Form.Item
                    className='col-md-6'
                    name="OnWorkDay"
                    label="到職日"
                    initialValue={dayjs(data?.OnWorkDay || new Date())}
                    rules={[{ required: true }]}
                >
                    <DatePicker disabled />
                </Form.Item>

                <Form.Item className='d-flex justify-content-end'>
                    <Button
                        onClick={() => {
                            ModelChangePassword.open({
                                UserId: data?.UserId
                            })
                        }}
                    >
                        更換密碼
                    </Button>
                    <Btn.Submit disabled={!permission?.isEditable({})} text='儲存' form={form} />
                </Form.Item>
            </Form>
        </div>
    )
}
export default Index
