import React from 'react'

import { Button, Form, Input } from 'antd'
import { FullMessage } from '../../method/notice'

const CompanyForm = (): JSX.Element => {
    return (
        <>
            <Form
                name="basic"
                initialValues={{ remember: true }}
                // onFinish={onFinish}
                // onFinishFailed={onFinishFailed}
                autoComplete="off"
            >
                <Form.Item
                    label="公司碼"
                    name="CompanyCode"
                    // rules={[{ required: true, message: 'Please input your username!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    label="公司名稱"
                    name="CompanyName"
                    // rules={[{ required: true, message: 'Please input your username!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    label="公司地址"
                    name="CompanyLocation"
                    // rules={[{ required: true, message: 'Please input your username!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    label="公司電話"
                    name="CompanyPhone"
                    // rules={[{ required: true, message: 'Please input your username!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    style={{ marginTop: '130px' }}
                    name="username"
                >
                    <Button
                        style={{ width: '100%' }}
                        htmlType="submit"
                        onClick={() => FullMessage.success('修改成功 ')}>修改</Button>
                </Form.Item>
            </Form>
        </>
    )
}
export default CompanyForm
