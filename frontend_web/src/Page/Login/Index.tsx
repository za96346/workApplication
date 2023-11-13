import { Form, Input } from 'antd'
import api from 'api/Index'
import React from 'react'
import Btn from 'shared/Button'

const Index = (): JSX.Element => {
    const [form] = Form.useForm()

    return (
        <div>
            <Form
                onFinish={(v) => { api.entry.login(v) }}
                name="validateOnly"
                autoComplete="off"
            >
                <Form.Item name="Account" label="帳號" rules={[{ required: true }]}>
                    <Input />
                </Form.Item>
                <Form.Item name="Password" label="密碼" rules={[{ required: true }]}>
                    <Input />
                </Form.Item>
                <Form.Item className='d-flex justify-content-center'>
                    <Btn.Submit text='登入' form={form} />
                </Form.Item>
            </Form>
        </div>
    )
}
export default Index
