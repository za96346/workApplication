import React from 'react'
import { Button, DatePicker, Form, Input, Spin } from 'antd'
import { FullMessage } from '../../method/notice'
import { useSelector } from 'react-redux'
import { RootState } from '../../reduxer/store'
import { SelfDataType } from '../../type'
import moment from 'moment'
import BanchSelector from '../../component/BanchSelector'
import PermessionSelector from '../../component/PermessionSelector'
import StatusSelector from '../../component/StatusSelector'

const personalForm = (): JSX.Element => {
    const { onFetchSelfData } = useSelector((state: RootState) => state.status)
    const { selfData }: { selfData: SelfDataType } = useSelector((state: RootState) => state.user)
    if (onFetchSelfData) {
        return <Spin size='large' />
    }
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
                    <Input defaultValue={selfData?.UserName || ''} />
                </Form.Item>
                <Form.Item
                    label="公司編號"
                    name="CompanyCode"
                    rules={[{ required: true, message: 'Please input your username!' }]}
                >
                    <Input defaultValue={selfData?.CompanyCode || ''} />
                </Form.Item>
                <Form.Item
                    label="帳號"
                    name="Account"
                    rules={[{ required: true, message: 'Please input your username!' }]}
                >
                    <Input defaultValue={selfData?.Account || ''} />
                </Form.Item>
                <Form.Item
                    label="密碼"
                    name="Password"
                    rules={[{ required: true, message: 'Please input your username!' }]}
                >
                    <Input.Password defaultValue={selfData?.Password || ''} visibilityToggle />
                </Form.Item>
                <Form.Item
                    label="到職日"
                    name="OnWorkDay"
                    rules={[{ required: true, message: 'Please input your username!' }]}
                >
                    <DatePicker defaultValue={moment(selfData?.OnWorkDay || '2001-07-01')} />
                </Form.Item>
                <Form.Item
                    label="部門"
                    name="Banch"
                    rules={[{ required: true, message: 'Please input your username!' }]}
                >
                    <BanchSelector defaultValue={selfData?.Banch || 0} />
                </Form.Item>
                <Form.Item
                    label="權限"
                    name="Permession"
                    rules={[{ required: true, message: 'Please input your username!' }]}
                >
                    <PermessionSelector defaultValue={selfData?.Permession} />
                </Form.Item>
                <Form.Item
                    label="狀態"
                    name="WorkState"
                    rules={[{ required: true, message: 'Please input your username!' }]}
                >
                    <StatusSelector defaultValue={selfData?.WorkState} />
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
