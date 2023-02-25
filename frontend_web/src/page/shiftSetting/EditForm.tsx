import { EditOutlined, PictureOutlined } from '@ant-design/icons'
import { Button, Form, Input, TimePicker } from 'antd'
import React from 'react'
import useReduceing from 'Hook/useReducing'
import rule from 'method/rule'
import dayjs from 'dayjs'

interface props {
    onFinish: (v: any) => Promise<void>
    types: 0 | 1
    currentId?: number
    btnText?: string
}
const defaultTimer = '00:00:00'

const EditForm = ({ onFinish, currentId = -10000, btnText = '新增' }: props): JSX.Element => {
    const { company } = useReduceing()
    const banchStyle = company.banchStyle?.find((item) => item.StyleId === currentId)
    return (
        <>
            <Form style={{ marginTop: '20px' }} scrollToFirstError onFinish={onFinish}>
                <a href='https://tw.piliapp.com/symbol/' style={{ fontSize: '0.5rem' }} target={'_blank'} rel="noreferrer" >圖標參考連結</a>
                <Form.Item rules={rule.banchStyleIcon()} initialValue={banchStyle?.Icon || ''} label="排班圖標" name="Icon">
                    <Input
                        placeholder='新增班別圖標'
                        prefix={<PictureOutlined />}
                    />
                </Form.Item>
                <Form.Item rules={rule.banchStyleTimeRangeName()} initialValue={banchStyle?.TimeRangeName || ''} label="班別名稱" name="TimeRangeName">
                    <Input
                        placeholder='新增班別名稱'
                        prefix={<EditOutlined />}
                    />
                </Form.Item>
                <Form.Item rules={rule.timePicker()} initialValue={dayjs(banchStyle?.OnShiftTime || defaultTimer, 'HH:mm:ss')} label="上班時段" name="OnShiftTime">
                    <TimePicker
                        inputReadOnly
                        allowClear={false}
                        minuteStep={5}
                        secondStep={60}
                        style={{ width: '100%' }}
                        placeholder='新增上班時段'
                    />
                </Form.Item>
                <Form.Item rules={rule.timePicker()} initialValue={dayjs(banchStyle?.OffShiftTime || defaultTimer, 'HH:mm:ss')} label="下班時段" name="OffShiftTime">
                    <TimePicker
                        inputReadOnly
                        allowClear={false}
                        minuteStep={5}
                        secondStep={60}
                        style={{ width: '100%' }}
                        placeholder='新增下班時段'
                    />
                </Form.Item>
                <Form.Item rules={rule.timePicker()} initialValue={dayjs(banchStyle?.RestTime || defaultTimer, 'HH:mm:ss')} label="休息時間" name="RestTime">
                    <TimePicker
                        inputReadOnly
                        allowClear={false}
                        minuteStep={5}
                        secondStep={60}
                        style={{ width: '100%' }}
                        placeholder="新增休息時間"
                    />
                </Form.Item>
                <Form.Item>
                    <div style={{ display: 'flex', justifyContent: 'flex-end' }}>
                        <Button htmlType='submit'>{btnText}</Button>
                    </div>
                </Form.Item>
            </Form>
        </>
    )
}
export default EditForm
