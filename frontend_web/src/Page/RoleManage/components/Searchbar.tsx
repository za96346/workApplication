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
                onFinish={(v) => { void api.role.get() }}
                id="roleManage"
                autoComplete="off"
                className='row'
            >
                <Form.Item
                    name="RoleName"
                    label="角色名稱"
                    initialValue={''}
                    className='col-md-6'
                >
                    <Input name='RoleName' />
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
