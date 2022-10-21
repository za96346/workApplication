import React, { useState } from 'react'
import { Button, DatePicker, Form, Input, Modal, Spin } from 'antd'
import { FullMessage } from '../../method/notice'
import moment from 'moment'
import BanchSelector from '../../component/BanchSelector'
import PermessionSelector from '../../component/PermessionSelector'
import useReduceing from '../../Hook/useReducing'
import { Button as MainBtn } from '../../component/Button'
import Insert from '../../component/Insert'
import api from '../../api/api'

const statusInit = {
    changePwdBtn: false,
    onChangePwd: false
}

const personalForm = (): JSX.Element => {
    const { loading, user } = useReduceing()
    const [status, setStatus] = useState({
        changePwdBtn: false,
        onChangePwd: false
    })

    const onFinish = async (v: any, types: 1 | 2): Promise<void> => {
        console.log(v)
        if (types === 1) {
            const res = await api.UpdateSelfData(v.UserName, v.OnWorkDay)
            console.log(res)
        } else if (types === 2) {
            const res = await api.changePassword({ ...v, Captcha: parseInt(v.Captcha) })
            console.log(res)
        }
    }
    if (loading.onFetchSelfData) {
        return <Spin size='large' />
    }
    return (
        <>
            <Modal
                closeIcon={<></>}
                footer={null}
                open={status.changePwdBtn}
            >
                <Form onFinish={async (v) => await onFinish(v, 2)} style={{ marginTop: '20px' }}>
                    <Insert.Captcha email={user.selfData.Account} />
                    <Insert.OldPwd />
                    <Insert.Pwd textNum={2} />
                    <Insert.PwdConfirm />
                    <div style={{ display: 'flex', justifyContent: 'space-between' }}>
                        <Button onClick={() => setStatus(statusInit)}>取消</Button>
                        <Button htmlType='submit'>送出</Button>
                    </div>
                </Form>
            </Modal>
            <Form
                name="basic"
                initialValues={{ remember: true }}
                onFinish={async (v) => await onFinish(v, 1)}
                autoComplete="off"
            >
                <Form.Item
                    label="姓名"
                    name="UserName"
                    initialValue={user.selfData?.UserName || ''}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    label="公司編號"
                    name="CompanyCode"
                    initialValue={user.selfData?.CompanyCode || ''}
                >
                    <Input disabled />
                </Form.Item>
                <Form.Item
                    label="帳號"
                    name="Account"
                    initialValue={user.selfData?.Account || ''}
                >
                    <Input disabled />
                </Form.Item>
                <Form.Item
                    label="到職日"
                    name="OnWorkDay"
                    initialValue={moment(user.selfData?.OnWorkDay || '2001-07-01')}
                >
                    <DatePicker />
                </Form.Item>
                <Form.Item
                    label="部門"
                    name="Banch"
                >
                    <BanchSelector disabled defaultValue={user.selfData?.Banch || 0} />
                </Form.Item>
                <Form.Item
                    label="權限"
                    name="Permession"
                >
                    <PermessionSelector disabled defaultValue={user.selfData?.Permession} />
                </Form.Item>
                <Form.Item
                    style={{ marginTop: '130px' }}
                    name="username"
                >
                    <Button
                        style={{ width: '100%', height: '50px' }}
                        htmlType="submit"
                        onClick={() => FullMessage.success('修改成功 ')}
                    >
                        修改
                    </Button>
                </Form.Item>
            </Form>
            <MainBtn
                text='更改密碼'
                onClick={() => setStatus((prev) => ({ ...prev, changePwdBtn: true }))}
            />
        </>
    )
}
export default personalForm
