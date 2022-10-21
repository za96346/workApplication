import React from 'react'
import { Button, Form } from 'antd'
// import { useNavigate } from 'react-router-dom'
import { LoginType } from '../../type'
import api from '../../api/api'
import Insert from '../../component/Insert'

const Login = (): JSX.Element => {
    // const navigate = useNavigate()
    const onFinish = async (v: LoginType): Promise<void> => {
        await api.login(v)
        console.log(v)
    }
    return (
        <>

            <Form onFinish={onFinish}>
                <Insert.Email />
                <Insert.Pwd textNum={1} />
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
