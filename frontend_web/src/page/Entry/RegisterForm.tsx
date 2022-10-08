import React, { useState } from 'react'
import { Form, Input } from 'antd'
import { KeyOutlined, NumberOutlined, MailOutlined, LockOutlined } from '@ant-design/icons'
import rule from '../../method/rule'

import { Button } from '../../component/Button'
import { Modal } from '../../component/Modal'

const Register = (): JSX.Element => {
    const [status, setStatus] = useState({
        modalText: '註冊成功',
        modalOpen: false
    })
    const updateModalOpen = (): void => {
        setStatus((prev) => ({ ...prev, modalOpen: !prev.modalOpen }))
    }
    const onClick = (): void => {
        updateModalOpen()
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
            <Form onFinish={onClick}>
                <Form.Item name="Account" rules={rule.email()}>
                    <Input size="large" placeholder="請輸入電子信箱" prefix={<MailOutlined />} />
                </Form.Item>
                <Form.Item name="captcha" rules={rule.captcha()}>
                    <Input size="large" placeholder="請輸入驗證碼" prefix={<LockOutlined />} />
                </Form.Item>
                <Form.Item name="Password" rules={rule.password()}>
                    <Input.Password visibilityToggle size="large" placeholder="請輸入密碼" prefix={<KeyOutlined />} />
                </Form.Item>
                <Form.Item name="PasswordConfirm" rules={rule.passwordConfirm()}>
                    <Input.Password visibilityToggle size="large" placeholder="確認密碼" prefix={<KeyOutlined />} />
                </Form.Item>
                <Form.Item name="CompanyCode" >
                    <Input size="large" placeholder="公司編號(選填)" prefix={<NumberOutlined />} />
                </Form.Item>
                <Button onClick={onClick} text="註冊" />
            </Form>
        </>
    )
}
export default Register
