import React, { useState } from 'react'
import { Button, Form } from 'antd'
import Insert from '../../component/Insert'
import api from '../../api/api'
import { useNavigate } from 'react-router-dom'

const ForgetPwdForm = (): JSX.Element => {
    const navigate = useNavigate()
    const [status, setStatus] = useState({
        email: ''
    })
    const onFinish = async (v: any): Promise<void> => {
        console.log(v)
        const res = await api.forgetPassword({ ...v, Captcha: parseInt(v.Captcha) })
        if (res.status) {
            navigate('/entry/login')
        }
    }
    return (
        <>
            <Form
                onFieldsChange={(e) => {
                    if (e[0].name[0] === 'Account') {
                        setStatus((prev) => ({ ...prev, email: e[0].value }))
                    }
                }}
                onFinish={onFinish}
            >
                <Insert.Email />
                <Insert.Captcha email={status.email} />
                <Insert.Pwd textNum={2} />
                <Insert.PwdConfirm/>
                <Button htmlType='submit' block>更改密碼</Button>
            </Form>
        </>
    )
}
export default ForgetPwdForm
