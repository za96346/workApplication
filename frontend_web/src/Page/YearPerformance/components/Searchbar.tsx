import { Divider, Form, Input, Select } from 'antd'
import api from 'api/Index'
import useRoleBanchList from 'hook/useRoleBanchUserList'
import React, { useEffect } from 'react'
import Btn from 'shared/Button'
import { FuncCodeEnum, OperationCodeEnum } from 'types/system'
import DateSelect from 'shared/AntdOverwrite/DateSelectRangePicker'
import YearPerformanceSession from '../methods/yearPerformanceSession'

const Searchbar = (): JSX.Element => {
    const [form] = Form.useForm()

    const rolBanchList = useRoleBanchList({
        funcCode: FuncCodeEnum.yearPerformance,
        operationCode: OperationCodeEnum.inquire
    })

    const onSearch = async (v: any): Promise<void> => {
        const currentParams = {
            ...v,
            ...DateSelect.getZhtwYear(v?.range)
        }
        void api.performance.getYear(currentParams)

        YearPerformanceSession.Instance.set((prev) => ({
            ...prev,
            currentParams
        }))
    }

    useEffect(() => {
        void onSearch(form.getFieldsValue())
    }, [])

    return (
        <>
            <Divider />
            <Form
                onFinish={(v) => {
                    void onSearch(v)
                }}
                id="performanceManage"
                autoComplete="off"
                className='row'
            >
                <Form.Item label='範圍搜尋' name='range' className='col-md-4'>
                    <DateSelect type='year' onChange={(v) => { }} />
                </Form.Item>
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
