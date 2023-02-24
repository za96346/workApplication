/* eslint-disable @typescript-eslint/restrict-plus-operands */
import { Button, DatePicker, Form } from 'antd'
import api from 'api/api'
import useReduceing from 'Hook/useReducing'
import React, { useRef } from 'react'

const SearchBar = (): JSX.Element => {
    const formRef = useRef<any>({})
    const { state } = useReduceing()
    const onSearch = (): void => {
        api.getShiftMonth({
            year: formRef.current?.range?.$y,
            month: formRef.current?.range?.$M + 1,
            banch: state.banchId
        })
        api.getShiftTotal({
            year: formRef.current?.range?.$y,
            month: formRef.current?.range?.$M + 1,
            banch: state.banchId
        })
    }
    return (
        <>
            <Form onValuesChange={(v, allV) => { formRef.current = allV }}>
                <Form.Item label='年月份' name='range' className='col-md-4'>
                    <DatePicker picker='month'/>
                </Form.Item>
            </Form>
            <Button onClick={onSearch}>搜尋</Button>
        </>
    )
}
export default SearchBar
