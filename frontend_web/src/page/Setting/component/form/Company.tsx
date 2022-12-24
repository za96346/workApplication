import React, { useEffect } from 'react'

import { Button, Form, Input, Select, Spin } from 'antd'
import api from '../../../../api/api'
import useReduceing from '../../../../Hook/useReducing'
const CompanyForm = (): JSX.Element => {
    const { company, loading, user } = useReduceing()
    const onFinish = async (v: any): Promise<void> => {
        console.log(v)
        const res = await api.updateCompanyInfo({ ...v, CompanyId: company.info.CompanyId })
        if (res.status) {
            await api.getCompanyInfo({ companyCode: user.selfData.CompanyCode })
        }
    }
    useEffect(() => {
        api.getCompanyInfo({ companyCode: user.selfData.CompanyCode })
    }, [])
    if (loading.onFetchCompany) {
        return <Spin />
    }
    return (
        <>
            <Form
                name="basic"
                initialValues={{ remember: true }}
                onFinish={onFinish}
                // onFinishFailed={onFinishFailed}
                autoComplete="off"
            >
                <Form.Item
                    label="公司碼"
                    name="CompanyCode"
                    initialValue={company.info?.CompanyCode || ''}
                    // rules={[{ required: true, message: 'Please input your username!' }]}
                >
                    <Input disabled />
                </Form.Item>
                <Form.Item
                    label="公司名稱"
                    name="CompanyName"
                    initialValue={company.info?.CompanyName || ''}
                    // rules={[{ required: true, message: 'Please input your username!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    label="結算日"
                    name="SettlementDate"
                    initialValue={company.info?.SettlementDate || ''}
                    // rules={[{ required: true, message: 'Please input your username!' }]}
                >
                    <Select>
                        {
                            new Array(31).fill('').map((item, index) => {
                                return (
                                    <Select.Option
                                        value={index + 1}
                                        key={index + 1}
                                    >
                                        {index + 1}
                                    </Select.Option>
                                )
                            })
                        }

                    </Select>
                </Form.Item>
                <Form.Item
                    label="公司地址"
                    name="CompanyLocation"
                    initialValue={company.info?.CompanyLocation || ''}
                    // rules={[{ required: true, message: 'Please input your username!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    label="公司電話"
                    name="CompanyPhoneNumber"
                    initialValue={company.info?.CompanyPhoneNumber || ''}
                    // rules={[{ required: true, message: 'Please input your username!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    style={{ marginTop: '130px' }}
                >
                    <Button
                        style={{ width: '100%' }}
                        htmlType="submit"
                    >修改</Button>
                </Form.Item>
            </Form>
        </>
    )
}
export default CompanyForm
