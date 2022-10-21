import React, { useState } from 'react'
import { Form, Input } from 'antd'
import { NumberOutlined } from '@ant-design/icons'

import { Button } from '../../component/Button'
import { Modal } from '../../component/Modal'
import api from '../../api/api'
import { useNavigate } from 'react-router-dom'
import Insert from '../../component/Insert'

const Register = (): JSX.Element => {
    const navigate = useNavigate()
    const [status, setStatus] = useState({
        modalText: '註冊成功',
        modalOpen: false,
        email: '',
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
                onFinish={onFinish}
            >
                <Insert.Email />
                <Insert.Captcha email={status.email} />
                <Insert.Pwd textNum={1}/>
                <Insert.PwdConfirm />
                <Form.Item initialValue="" name="CompanyCode" >
                    <Input size="large" placeholder="公司編號(選填)" prefix={<NumberOutlined />} />
                </Form.Item>
                <Button onClick={() => {}} text="註冊" />
            </Form>
        </>
    )
}
export default Register
