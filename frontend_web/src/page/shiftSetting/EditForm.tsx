import { EditOutlined, PictureOutlined } from '@ant-design/icons'
import { Button, Form, Input, TimePicker } from 'antd'
import React from 'react'
import rule from '../../method/rule'

interface props {
    onFinish: (v: any) => Promise<void>
    types: 0 | 1
    btnText?: string
}

const EditForm = ({ onFinish, types, btnText = '新增' }: props): JSX.Element => {
    return (
        <>
            {/* types 0 ,, banch style form */}
            {
                types === 0 && (
                    <Form style={{ marginTop: '20px' }} scrollToFirstError onFinish={onFinish}>
                        <Form.Item rules={rule.banchStyleIcon()} initialValue="" label="排班圖標" name="Icon">
                            <Input placeholder='新增班別圖標' prefix={<PictureOutlined />} />
                        </Form.Item>
                        <Form.Item rules={rule.banchStyleTimeRangeName()} initialValue="" label="班別名稱" name="TimeRangeName">
                            <Input placeholder='新增班別名稱' prefix={<EditOutlined />} />
                        </Form.Item>
                        <Form.Item rules={rule.timePicker()} initialValue="" label="上班時段" name="OnShiftTime">
                            <TimePicker style={{ width: '100%' }} placeholder='新增上班時段' />
                        </Form.Item>
                        <Form.Item rules={rule.timePicker()} initialValue="" label="下班時段" name="OffShiftTime">
                            <TimePicker style={{ width: '100%' }} placeholder='新增下班時段' />
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
                    <></>
                )
            }
        </>
    )
}
export default EditForm
