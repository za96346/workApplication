import { Form, Input } from 'antd'
import api from 'api/Index'

import { useAppSelector } from 'hook/redux'
import usePermission from 'hook/usePermission'
import React, { useEffect } from 'react'
import Btn from 'shared/Button'
import { funcCode } from 'types/system'

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
                disabled={!permission?.isEditable}
                onFinish={(v) => { }}
                name="validateOnly"
                className='row'
                autoComplete="off"
            >
                <Form.Item
                    className='col-md-6'
                    name="CompanyName"
                    label="公司名稱"
                    rules={[{ required: true }]}
                >
                    <Input defaultValue={data?.CompanyName} />
                </Form.Item>
                <Form.Item
                    className='col-md-6'
                    name="CompanyPhone"
                    label="公司電話"
                    rules={[{ required: true }]}
                >
                    <Input defaultValue={data?.companyPhoneNumber} />
                </Form.Item>
                <Form.Item
                    className='col-md-6'
                    name="CompanyCode"
                    label="公司代碼"
                    rules={[{ required: true }]}
                >
                    <Input defaultValue={data?.CompanyCode} />
                </Form.Item>
                <Form.Item
                    className='col-md-6'
                    name="CompanyLocation"
                    label="公司位置"
                    rules={[{ required: true }]}
                >
                    <Input defaultValue={data?.CompanyLocation} />
                </Form.Item>
                <Form.Item
                    className='col-md-6'
                    name="CompanyId"
                    label="公司id"
                    rules={[{ required: true }]}
                >
                    <Input defaultValue={data?.CompanyId} />
                </Form.Item>
                <Form.Item
                    className='col-md-6'
                    name="BossId"
                    label="boss id"
                    rules={[{ required: true }]}
                >
                    <Input defaultValue={data?.BossId} />
                </Form.Item>

                <Form.Item className='d-flex justify-content-end'>
                    <Btn.Submit text='儲存' form={form} />
                </Form.Item>
            </Form>
        </div>
    )
}
export default Index
