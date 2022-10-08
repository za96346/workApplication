import React from 'react'
import { Button, Form, Input } from 'antd'
import { KeyOutlined, UserOutlined } from '@ant-design/icons'
// import { useNavigate } from 'react-router-dom'
import rule from '../../method/rule'
import api from '../../api/api'
import { LoginType } from '../../type'

const Login = (): JSX.Element => {
    // const navigate = useNavigate()
    const onFinish = async (v: LoginType): Promise<void> => {
        await api.login(v)
        console.log(v)
    }
    return (
        <>

            <Form onFinish={onFinish}>
                <Form.Item name="Account" rules={rule.email()}>
                    <Input size="large" placeholder="請輸入電子信箱" prefix={<UserOutlined />} />
                </Form.Item>
                <Form.Item name="Password" rules={rule.password()}>
                    <Input.Password visibilityToggle size="large" placeholder="請輸入密碼" prefix={<KeyOutlined />} />
                </Form.Item>
                <Form.Item>
                    <Button
                        block
                        htmlType="submit">
                        登入
                    </Button>
                </Form.Item>

            </Form>
        </>
    )
}

export default Login
