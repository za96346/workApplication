import { Button, DatePicker, Form } from "antd"
import api from "api/api"
import useReduceing from "Hook/useReducing"
import React, { useRef } from "react"

const SearchBar = () => {
    const formRef = useRef({})
    const { state } = useReduceing()
    const onSearch = () => {
        api.getShiftMonth({
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
