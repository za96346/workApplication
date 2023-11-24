import { Divider, Form, Input } from 'antd'
import api from 'api/Index'
import React from 'react'
import Btn from 'shared/Button'

const Searchbar = (): JSX.Element => {
    const [form] = Form.useForm()

    return (
        <>
            <Divider />
            <Form
                onFinish={(v) => { void api.companyBanch.get(v) }}
                name="validateOnly"
                autoComplete="off"
                className='row'
            >
                <Form.Item
                    name="BanchName"
                    label="部門"
                    className='col-md-6'
                >
                    <Input />
                </Form.Item>
                <Form.Item className='d-flex justify-content-end'>
                    <Btn.Reset />
                    <Btn.Submit text='搜尋' form={form} />
                </Form.Item>
            </Form>
            <Divider />
        </>
    )
}
export default Searchbar
