import React, { useState } from 'react'
import { Form, Input, Button as AntdBtn, Spin } from 'antd'
import { KeyOutlined, NumberOutlined, MailOutlined, LockOutlined } from '@ant-design/icons'
import rule from '../../method/rule'

import { Button } from '../../component/Button'
import { Modal } from '../../component/Modal'
import api from '../../api/api'
import { useNavigate } from 'react-router-dom'

const Register = (): JSX.Element => {
    const navigate = useNavigate()
    const [status, setStatus] = useState({
        modalText: '註冊成功',
        modalOpen: false,
        email: '',
        captchaBtn: false,
        confirmBtn: false
    })
    const updateModalOpen = (): void => {
        setStatus((prev) => ({ ...prev, modalOpen: !prev.modalOpen }))
    }
    const onFinish = async (v: any): Promise<void> => {
        setStatus((prev) => ({ ...prev, confirmBtn: true }))
        const res = await api.register({ ...v, Captcha: parseInt(v.Captcha, 10) })
        setStatus((prev) => ({ ...prev, confirmBtn: false }))
        if (res) {
            navigate('/entry/login')
        }
    }
    const getEmailCaptcha = async (): Promise<any> => {
        setStatus((prev) => ({ ...prev, captchaBtn: true }))
        const res = await api.getEmailCaptcha(status.email)
        console.log(res)
        setStatus((prev) => ({ ...prev, captchaBtn: false }))
    }

    return (
        <>
            <Modal
                onOk={updateModalOpen}
                onCancel={updateModalOpen}
                open={status.modalOpen}
            >
                {status.modalText}
            </Modal>
            <Form
                onFieldsChange={(e) => {
                    if (e[0].name[0] === 'Account') {
                        setStatus((prev) => ({ ...prev, email: e[0].value }))
                    }
                }}
                onFinish={onFinish}>
                <Form.Item name="Account" rules={rule.email()}>
                    <Input size="large" placeholder="請輸入電子信箱" prefix={<MailOutlined />} />
                </Form.Item>
                <div style={{ display: 'flex', alignItems: 'flex-start', justifyContent: 'space-between' }}>
                    <Form.Item name="Captcha" rules={rule.captcha()}>
                        <Input size="large" placeholder="請輸入驗證碼" prefix={<LockOutlined />} />
                    </Form.Item>
                    <AntdBtn disabled={status.captchaBtn} onClick={getEmailCaptcha} style={{ height: '40px' }}>

                        {
                            status.captchaBtn
                                ? <Spin size={'small'} style={{ fontSize: '0.3rem' }} tip={'發送中...'} />
                                : '發送信箱驗證碼'
                        }
                    </AntdBtn>
                </div>
                <Form.Item name="Password" rules={rule.password()}>
                    <Input.Password visibilityToggle size="large" placeholder="請輸入密碼" prefix={<KeyOutlined />} />
                </Form.Item>
                <Form.Item name="PasswordConfirm" rules={rule.passwordConfirm()}>
                    <Input.Password visibilityToggle size="large" placeholder="確認密碼" prefix={<KeyOutlined />} />
                </Form.Item>
                <Form.Item initialValue="" name="CompanyCode" >
                    <Input size="large" placeholder="公司編號(選填)" prefix={<NumberOutlined />} />
                </Form.Item>
                <Button onClick={() => {}} text="註冊" />
            </Form>
        </>
    )
}
export default Register
