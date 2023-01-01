/* eslint-disable react/display-name */
import { KeyOutlined, LockOutlined, MailOutlined } from '@ant-design/icons'
import { Button, Form, Input, Spin } from 'antd'
import React, { useState } from 'react'
import api from '../api/api'
import rule from '../method/rule'
interface EmailProps {
    email: string
}

const Insert = (): JSX.Element => {
    return (
        <>
        </>
    )
}

Insert.Pwd = ({ textNum }: { textNum: 1 | 2 }): JSX.Element => {
    return (
        <>
            <Form.Item name="Password" rules={rule.password()}>
                <Input.Password
                    visibilityToggle size="large"
                    placeholder={textNum === 1 ? '請輸入密碼' : '請輸入新密碼'}
                    prefix={<KeyOutlined />}
                />
            </Form.Item>
        </>
    )
}
Insert.OldPwd = (): JSX.Element => {
    return (
        <>
            <Form.Item name="OldPassword" rules={rule.password()}>
                <Input.Password
                    visibilityToggle size="large"
                    placeholder={'請輸入舊密碼'}
                    prefix={<KeyOutlined />}
                />
            </Form.Item>
        </>
    )
}
Insert.PwdConfirm = (): JSX.Element => {
    return (
        <Form.Item name="PasswordConfirm" rules={rule.passwordConfirm()}>
            <Input.Password visibilityToggle size="large" placeholder="確認密碼" prefix={<KeyOutlined />} />
        </Form.Item>
    )
}

Insert.Email = (): JSX.Element => {
    return (
        <Form.Item name="Account" rules={rule.email()}>
            <Input size="large" placeholder="請輸入 電子信箱 / 員工編號" prefix={<MailOutlined />} />
        </Form.Item>
    )
}

Insert.Captcha = ({ email }: EmailProps): JSX.Element => {
    const [status, setStatus] = useState({
        captchaBtn: false
    })
    const getEmailCaptcha = async (): Promise<void> => {
        setStatus((prev) => ({ ...prev, captchaBtn: true }))
        await api.getEmailCaptcha(email)
        setStatus((prev) => ({ ...prev, captchaBtn: false }))
    }
    return (
        <>
            <div style={{ display: 'flex', alignItems: 'flex-start', justifyContent: 'space-between' }}>
                <Form.Item name="Captcha" rules={rule.captcha()}>
                    <Input size="large" placeholder="請輸入驗證碼" prefix={<LockOutlined />} />
                </Form.Item>
                <Button disabled={status.captchaBtn} onClick={getEmailCaptcha} style={{ height: '40px' }}>

                    {
                        status.captchaBtn
                            ? <Spin size={'small'} style={{ fontSize: '0.3rem' }} tip={'發送中...'} />
                            : '發送信箱驗證碼'
                    }
                </Button>
            </div>
        </>
    )
}
export default Insert
