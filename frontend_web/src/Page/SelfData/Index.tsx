import { DatePicker, Form, Input } from 'antd'
import api from 'api/Index'
import dayjs from 'dayjs'
import { useAppSelector } from 'hook/redux'
import usePermission from 'hook/usePermission'
import React, { useEffect } from 'react'
import Btn from 'shared/Button'
import { funcCode } from 'types/system'

const Index = (): JSX.Element => {
    const [form] = Form.useForm()
    const data = useAppSelector((v) => v?.user?.mine)
    const permission = usePermission({ funcCode: funcCode.selfData })

    useEffect(() => {
        void api.user.getMine()
    }, [])

    return (
        <div>
            <Form
                disabled={!permission?.isEditable}
                onFinish={(v) => { }}
                name="validateOnly"
                className='row'
                autoComplete="off"
            >
                <Form.Item
                    className='col-md-6'
                    name="Account"
                    label="帳號"
                    rules={[{ required: true }]}
                >
                    <Input defaultValue={data?.Account} />
                </Form.Item>
                <Form.Item
                    className='col-md-6'
                    name="UserName"
                    label="姓名"
                    rules={[{ required: true }]}
                >
                    <Input defaultValue={data?.UserName} />
                </Form.Item>
                <Form.Item
                    className='col-md-6'
                    name="EmployeeNumber"
                    label="員工編號"
                    rules={[{ required: true }]}
                >
                    <Input defaultValue={data?.EmployeeNumber} />
                </Form.Item>
                <Form.Item
                    className='col-md-6'
                    name="Banch"
                    label="部門"
                    rules={[{ required: true }]}
                >
                    <Input defaultValue={data?.BanchId} />
                </Form.Item>
                <Form.Item
                    className='col-md-6'
                    name="UserId"
                    label="userId"
                    rules={[{ required: true }]}
                >
                    <Input defaultValue={data?.UserId} />
                </Form.Item>
                <Form.Item
                    className='col-md-6'
                    name="RoleId"
                    label="roleId"
                    rules={[{ required: true }]}
                >
                    <Input defaultValue={data?.RoleId} />
                </Form.Item>
                <Form.Item
                    className='col-md-6'
                    name="OnWorkDay"
                    label="到職日"
                    rules={[{ required: true }]}
                >
                    <DatePicker defaultValue={dayjs(data?.OnWorkDay)} />
                </Form.Item>

                <Form.Item className='d-flex justify-content-end'>
                    <Btn.Submit text='儲存' form={form} />
                </Form.Item>
            </Form>
        </div>
    )
}
export default Index
