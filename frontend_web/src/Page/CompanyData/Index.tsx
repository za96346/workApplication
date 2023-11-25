import { Form, Input } from 'antd'
import api from 'api/Index'

import { useAppSelector } from 'hook/redux'
import { usePermission } from 'hook/usePermission'
import React, { useEffect } from 'react'
import Btn from 'shared/Button'
import { funcCode } from 'types/system'
import UserSelector from 'shared/UserSelector/Index'

const Index = (): JSX.Element => {
    const [form] = Form.useForm()
    const data = useAppSelector((v) => v?.company?.mine)
    const permission = usePermission({ funcCode: funcCode.companyData })

    useEffect(() => {
        void api.company.getMine()
    }, [])

    return (
        <div>
            <Form
                disabled={!permission?.isEditable({})}
                onFinish={(v) => {
                    void api.company.update(v)
                        .then(() => {
                            void api.company.getMine()
                        })
                }}
                name="validateOnly"
                className='row'
                autoComplete="off"
                form={form}
            >
                <Form.Item
                    className='col-md-6'
                    name="CompanyName"
                    label="公司名稱"
                    initialValue={data?.CompanyName}
                    rules={[{ required: true }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    className='col-md-6'
                    name="CompanyPhoneNumber"
                    label="公司電話"
                    initialValue={data?.companyPhoneNumber}
                    rules={[{ required: true }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    className='col-md-6'
                    name="CompanyCode"
                    label="公司代碼"
                    initialValue={data?.CompanyCode}
                    rules={[{ required: true }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    className='col-md-6'
                    name="CompanyLocation"
                    label="公司位置"
                    initialValue={data?.CompanyLocation}
                    rules={[{ required: true }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    className='col-md-6'
                    label="負責人"
                    name="BossId"
                    rules={[{ required: true }]}
                >
                    <UserSelector
                        onChange={(v) => {
                            form.setFieldValue('BossId', v?.[0]?.UserId || null)
                        }}
                        type='radio'
                        subComponents='tag'
                        defaultValue={[data?.BossId]}
                    />
                </Form.Item>

                <Form.Item className='d-flex justify-content-end'>
                    <Btn.Submit
                        disabled={!permission?.isEditable({})}
                        text='儲存'
                        form={form}
                    />
                </Form.Item>
            </Form>
        </div>
    )
}
export default Index
