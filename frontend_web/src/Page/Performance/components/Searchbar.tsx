import { Divider, Form, Input, Select } from 'antd'
import api from 'api/Index'
import useRoleBanchList from 'hook/useRoleBanchUserList'
import React, { useEffect } from 'react'
import Btn from 'shared/Button'
import { FuncCodeEnum, OperationCodeEnum } from 'types/system'
import DateSelect from 'shared/AntdOverwrite/DateSelectRangePicker'
import dayjs from 'dayjs'
import PerformanceSession from '../methods/performanceSession'

interface propsType {
    userId: number
    year: number
}

const Searchbar = ({
    userId,
    year
}: propsType): JSX.Element => {
    const [form] = Form.useForm()

    const rolBanchList = useRoleBanchList({
        funcCode: FuncCodeEnum.performance,
        operationCode: OperationCodeEnum.inquire
    })

    const onSearch = (v: any): void => {
        const currentParams = {
            ...v,
            ...DateSelect.getZhtwDate(v?.range),
            ...(year ? { StartDate: `${year}-01` } : {}),
            ...(year ? { EndDate: `${year}-12` } : {}),
            ...(userId ? { UserId: userId } : {})
        }

        void api.performance.get(currentParams)

        PerformanceSession.Instance.set((prev) => ({
            ...prev,
            currentParams
        }))
    }

    useEffect(() => {
        onSearch(form.getFieldsValue())
    }, [])

    return (
        <>
            <Divider />
            <Form
                onFinish={(v) => {
                    onSearch(v)
                }}
                id="performanceManage"
                autoComplete="off"
                className='row'
            >
                <Form.Item
                    label='範圍搜尋'
                    name='range'
                    className='col-md-4'
                >
                    <DateSelect
                        {...(year
                            ? {
                                defaultValue: [
                                    dayjs(`${year + 1911}-01`),
                                    dayjs(`${year + 1911}-12`)
                                ]
                            }
                            : {}
                        )}
                        type='month'
                        onChange={(v) => { }}
                    />
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
