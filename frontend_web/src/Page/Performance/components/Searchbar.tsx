import { Divider, Form, Input, Select } from 'antd'
import api from 'api/Index'
import useRoleBanchList from 'hook/useRoleBanchUserList'
import React from 'react'
import Btn from 'shared/Button'
import { FuncCodeEnum, OperationCodeEnum } from 'types/system'

const Searchbar = (): JSX.Element => {
    const [form] = Form.useForm()

    const rolBanchList = useRoleBanchList({
        funcCode: FuncCodeEnum.performance,
        operationCode: OperationCodeEnum.inquire
    })

    return (
        <>
            <Divider />
            <Form
                onFinish={(v) => {
                    void api.performance.get(v)
                }}
                id="performanceManage"
                autoComplete="off"
                className='row'
            >
                <Form.Item
                    name="BanchId"
                    label="部門"
                    className='col-md-6'
                >
                    <Select allowClear options={rolBanchList.banchSelectList} />
                </Form.Item>
                <Form.Item
                    name="RoleId"
                    label="角色"
                    className='col-md-6'
                >
                    <Select allowClear options={rolBanchList.roleSelectList} />
                </Form.Item>
                <Form.Item
                    name="UserName"
                    label="姓名"
                    className='col-md-6'
                >
                    <Input name='UserName' />
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
