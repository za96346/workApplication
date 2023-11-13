import { Divider, Form, Input, Select } from 'antd'
import api from 'api/Index'
import React from 'react'
import Btn from 'shared/Button'

const Searchbar = (): JSX.Element => {
    const [form] = Form.useForm()

    return (
        <>
            <Divider />
            <Form
                onFinish={(v) => { void api.user.getEmployee() }}
                name="validateOnly"
                autoComplete="off"
                className='row'
            >
                <Form.Item
                    name="BanchId"
                    label="部門"
                    className='col-md-6'
                >
                    <Select />
                </Form.Item>
                <Form.Item
                    name="UserName"
                    label="姓名"
                    className='col-md-6'
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="EmployeeNumber"
                    label="員工編號"
                    className='col-md-6'
                >
                    <Input />
                </Form.Item>
                <Form.Item className='d-flex justify-content-end'>
                    <Btn.Submit text='搜尋' form={form} />
                </Form.Item>
            </Form>
            <Divider />
        </>
    )
}
export default Searchbar
