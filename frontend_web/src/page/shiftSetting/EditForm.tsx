import { EditOutlined, PictureOutlined } from '@ant-design/icons'
import { Button, Form, Input, InputNumber, TimePicker } from 'antd'
import moment from 'moment'
import React from 'react'
import WeekDaySelector from '../../component/WeekDaySelector'
import WeekTypeSelector from '../../component/WeekTypeSelector'
import useReduceing from '../../Hook/useReducing'
import rule from '../../method/rule'

interface props {
    onFinish: (v: any) => Promise<void>
    types: 0 | 1
    currentId?: number
    btnText?: string
}
const defaultTimer = '00:00:00'

const EditForm = ({ onFinish, types, currentId = -10000, btnText = '新增' }: props): JSX.Element => {
    const { company } = useReduceing()
    const banchStyle = company.banchStyle?.find((item) => item.StyleId === currentId)
    const banchRule = company.banchRule.find((item) => item.RuleId === currentId)
    return (
        <>
            {/* types 0 ,, banch style form */}
            {
                types === 0 && (
                    <Form style={{ marginTop: '20px' }} scrollToFirstError onFinish={onFinish}>
                        <Form.Item rules={rule.banchStyleIcon()} initialValue={banchStyle?.Icon || ''} label="排班圖標" name="Icon">
                            <Input placeholder='新增班別圖標' prefix={<PictureOutlined />} />
                        </Form.Item>
                        <Form.Item rules={rule.banchStyleTimeRangeName()} initialValue={banchStyle?.TimeRangeName || ''} label="班別名稱" name="TimeRangeName">
                            <Input placeholder='新增班別名稱' prefix={<EditOutlined />} />
                        </Form.Item>
                        <Form.Item rules={rule.timePicker()} initialValue={moment(banchStyle?.OnShiftTime || defaultTimer, 'HH:mm:ss')} label="上班時段" name="OnShiftTime">
                            <TimePicker style={{ width: '100%' }} placeholder='新增上班時段' />
                        </Form.Item>
                        <Form.Item rules={rule.timePicker()} initialValue={moment(banchStyle?.OffShiftTime || defaultTimer, 'HH:mm:ss')} label="下班時段" name="OffShiftTime">
                            <TimePicker style={{ width: '100%' }} placeholder='新增下班時段' />
                        </Form.Item>
                        <Form.Item rules={rule.timePicker()} initialValue={moment(banchStyle?.RestTime || defaultTimer, 'HH:mm:ss')} label="休息時間" name="RestTime">
                            <TimePicker style={{ width: '100%' }} placeholder="新增休息時間" />
                        </Form.Item>
                        <Form.Item>
                            <div style={{ display: 'flex', justifyContent: 'flex-end' }}>
                                <Button htmlType='submit'>{btnText}</Button>
                            </div>
                        </Form.Item>
                    </Form>
                )
            }

            {/* types 0 ,, banch rule form */}
            {
                types === 1 && (
                    <>
                        <Form style={{ marginTop: '20px' }} onFinish={onFinish}>
                            <Form.Item label="平假日" initialValue={banchRule?.WeekType || 1} name="WeekType">
                                <WeekTypeSelector defaultValue={(banchRule?.WeekType || 1) as 1 | 2} />
                            </Form.Item>
                            <Form.Item label="星期" initialValue={banchRule?.WeekDay || 1} name='WeekDay'>
                                <WeekDaySelector defaultValue={(banchRule?.WeekDay || 1)} />
                            </Form.Item>
                            <div style={{ display: 'flex', justifyContent: 'space-around' }}>
                                <Form.Item label="最少員工數" initialValue={banchRule?.MinPeople || 1} name="MinPeople">
                                    <InputNumber min={0} />
                                </Form.Item>
                                <Form.Item label="最多員工數" initialValue={banchRule?.MaxPeople || 1} name="MaxPeople">
                                    <InputNumber min={0} />
                                </Form.Item>
                            </div>
                            <Form.Item rules={rule.timePicker()} initialValue={moment(banchRule?.OnShiftTime || defaultTimer, 'HH:mm:ss')} label="上班時段" name="OnShiftTime">
                                <TimePicker style={{ width: '100%' }} placeholder='新增上班時段' />
                            </Form.Item>
                            <Form.Item rules={rule.timePicker()} initialValue={moment(banchRule?.OffShiftTime || defaultTimer, 'HH:mm:ss')} label="下班時段" name="OffShiftTime">
                                <TimePicker style={{ width: '100%' }} placeholder='新增下班時段' />
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
        </>
    )
}
export default EditForm
