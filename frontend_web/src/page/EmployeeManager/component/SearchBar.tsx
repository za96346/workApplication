/* eslint-disable array-callback-return */
import { SearchOutlined } from '@ant-design/icons'
import { Button, Divider, Form, Input } from 'antd'
import React, { useRef, useEffect } from 'react'
import api from 'api/api'

import BanchSelector from 'Share/BanchSelector'
import StatusSelector from 'Share/StatusSelector'
import useReduceing from 'Hook/useReducing'

const formInit = {
    name: '',
    workState: 'on',
    banch: null
}
const SearchBar = ({ reSearching }: { reSearching: boolean }): JSX.Element => {
    const form = useRef({ ...formInit })
    const { user } = useReduceing()
    const onSearch = (): void => {
        api.getUserAll(form.current)
    }
    useEffect(() => {
        onSearch()
    }, [reSearching])
    return (
        <>
            <div className={window.styles.empManagerFilter}>
                {/* <button className='btn btn-secondary d-flex align-items-center m-1'>
                    新增員工
                </button> */}
                <Divider/>
                <Form onFinish={onSearch} className='row'>
                    <Form.Item className='col-md-4' label={'姓名'}>
                        <Input
                            onChange={(e) => { form.current = { ...form.current, name: e.target.value } }}
                            style={{ width: '100%' }}
                            prefix={<SearchOutlined/>}
                            placeholder='姓名'
                        />
                    </Form.Item>
                    {
                        user.selfData.Permession === 100 && (
                            <Form.Item className='col-md-4' label={'組別'}>
                                <BanchSelector onChange={(e) => { form.current = { ...form.current, banch: e } }} defaultValue={0}/>
                            </Form.Item>
                        )
                    }

                    <Form.Item className='col-md-4' label={'狀態'}>
                        <StatusSelector onChange={(e) => { form.current = { ...form.current, workState: e } }} defaultValue={'on'}/>
                    </Form.Item>
                    <div className='w-100 d-flex justify-content-end'>
                        <Button htmlType='submit' onClick={onSearch} icon={<SearchOutlined />}>
                        搜尋 / reload
                        </Button>
                    </div>

                </Form>
            </div>
        </>
    )
}
export default SearchBar
