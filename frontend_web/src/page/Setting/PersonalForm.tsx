import React from 'react'
import { Button, DatePicker, Form, Input, Select } from 'antd'
import { FullMessage } from '../../method/notice'

const personalForm = (): JSX.Element => {
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
                    label="姓名"
                    name="Name"
                    rules={[{ required: true, message: 'Please input your username!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    label="公司編號"
                    name="CompanyCode"
                    rules={[{ required: true, message: 'Please input your username!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    label="帳號"
                    name="Account"
                    rules={[{ required: true, message: 'Please input your username!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    label="密碼"
                    name="Password"
                    rules={[{ required: true, message: 'Please input your username!' }]}
                >
                    <Input.Password visibilityToggle />
                </Form.Item>
                <Form.Item
                    label="到職日"
                    name="OnWorkDay"
                    rules={[{ required: true, message: 'Please input your username!' }]}
                >
                    <DatePicker />
                </Form.Item>
                <Form.Item
                    label="部門"
                    name="Banch"
                    rules={[{ required: true, message: 'Please input your username!' }]}
                >
                    <Select />
                </Form.Item>
                <Form.Item
                    label="權限"
                    name="Permession"
                    rules={[{ required: true, message: 'Please input your username!' }]}
                >
                    <Select />
                </Form.Item>
                <Form.Item
                    label="狀態"
                    name="WorkState"
                    rules={[{ required: true, message: 'Please input your username!' }]}
                >
                    <Select />
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
export default personalForm
