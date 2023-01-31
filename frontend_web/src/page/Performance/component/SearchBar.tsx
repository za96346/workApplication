import { Button, Divider, Form, Input } from 'antd'
import React, { useEffect, useRef } from 'react'
import api from 'api/api'
import dayjs from 'dayjs'
import StatusSelector from 'Share/StatusSelector'
import DateSelect from '../component/DateSelect'
import useReduceing from 'Hook/useReducing'
import { SearchOutlined } from '@ant-design/icons'

interface props {
    reFetchRef: React.MutableRefObject<() => Promise<void>>
}
const SearchBar = ({
    reFetchRef
}: props): JSX.Element => {
    const { state } = useReduceing()
    const formData = useRef({
        name: '',
        range: [dayjs(), dayjs()],
        workState: 'on'
    })

    // 送出收尋
    const onSearch = async (): Promise<void> => {
        // console.log(formData.current.range)
        await api.getPerformance({
            banchId: state.banchId,
            name: formData.current.name,
            startYear: formData.current.range[0]?.year() - 1911,
            startMonth: formData.current.range[0]?.month() + 1,
            endYear: formData.current.range[1]?.year() - 1911,
            endMonth: formData.current.range[1]?.month() + 1,
            workState: formData.current.workState || 'on'
        })
    }
    reFetchRef.current = onSearch
    useEffect(() => {
        onSearch()
    }, [state.banchId])
    return (
        <>

            <Divider />
            <Form onValuesChange={(v, allV) => { formData.current = allV }} className='row'>
                <Form.Item label='範圍搜尋' name='range' className='col-md-4'>
                    <DateSelect type='month' onChange={(v) => { }} />
                </Form.Item>
                <Form.Item label='姓名' name='name' className='col-md-4'>
                    <Input placeholder='輸入姓名'/>
                </Form.Item>
                <Form.Item label='狀態' name='workState' className='col-md-4'>
                    <StatusSelector defaultValue={'on'}/>
                </Form.Item>
                <div className='d-flex w-100 justify-content-end'>
                    <Button htmlType='submit' onClick={onSearch} icon={<SearchOutlined />}>
                        搜尋/reload
                    </Button>
                </div>
            </Form>
            <Divider />
            <Button onClick={() => { window.open('/printWord', '績效評核', 'height=800,width=800') }}>
                    列印表單
            </Button>
            <Button onClick={() => { window.open('/printList', '績效評核', 'height=800,width=800') }}>
                    列印清單
            </Button>
        </>
    )
}
export default SearchBar
