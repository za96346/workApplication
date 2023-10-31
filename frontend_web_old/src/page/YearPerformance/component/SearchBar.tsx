import { Button, Form, Input } from 'antd'
import api from 'api/api'
import DateSelect from 'page/Performance/component/DateSelect'
import React, { useEffect } from 'react'

const SearchBar = (): JSX.Element => {
    const onSearch = (v: any): void => {
        const start = v.timeRange[0]?.$y - 1911
        const end = v.timeRange[1]?.$y - 1911
        console.log(v)
        api.getYearPerformance({
            startYear: start,
            endYear: end,
            userName: v.userName
        })
    }
    useEffect(() => {
        api.getYearPerformance({
            startYear: new Date().getFullYear() - 1911,
            endYear: new Date().getFullYear() - 1911,
            userName: ''
        })
    }, [])
    return (
        <div
            className={window.styles.performanceBlock}
        >
            <Form
                className='row'
                onFinish={onSearch}
            >
                <Form.Item className='col-md-4' label="查詢範圍" name='timeRange'>
                    <DateSelect type='year' onChange={(v) => { console.log(v) }} />
                </Form.Item>
                {/* <Form.Item className='col-md-4' label="結束年份" name='endYear'>
                    <Input/>
                </Form.Item> */}
                <Form.Item className='col-md-4' label="姓名" name='userName'>
                    <Input/>
                </Form.Item>
                <div className='w-100 d-flex justify-content-end'>
                    <Button className='' htmlType='submit'>搜尋 / reload</Button>
                </div>
            </Form>
        </div>
    )
}
export default SearchBar
