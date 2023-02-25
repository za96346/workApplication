import React, { useRef, useEffect } from 'react'
import { Modal, Form, Input, InputNumber, DatePicker, Select } from 'antd'
import { BanchType, performanceType } from 'type'
import useReduceing from 'Hook/useReducing'
import DescribeValue from '../DescribeValue'
import dayjs from 'dayjs'
import api from 'api/api'

interface props {
    open: boolean
    value: performanceType | null
    onClose: () => void
    onSave: ((v: performanceType) => void)
    type: '1' | '4' // 1 編輯,  4 新增
    banchId: BanchType['Id']
}
const Edit = ({
    open,
    value,
    onClose,
    onSave,
    type,
    banchId
}: props): JSX.Element => {
    const form = useRef<performanceType>(value)
    const { user, company } = useReduceing()
    const disabled = user.selfData?.Permession === 2 ||
        (user.selfData?.Permession === 1 &&
        value?.UserId === user.selfData?.UserId)
    useEffect(() => {
        api.getUserAll({
            workState: 'on',
            name: '',
            banch: banchId
        })
    }, [])
    return (
        <Modal
            title={type === '1' ? '編輯績效' : '新增績效'}
            destroyOnClose
            open={open}
            onCancel={onClose}
            onOk={() => { onSave({ ...value, ...form.current }) }}
            okText="儲存"
            cancelText="取消"
        >
            {
                type === '1' && (
                    <DescribeValue value={value}/>
                )
            }
            <Form
                onValuesChange={(v, allV) => {
                    form.current = {
                        ...allV,
                        Year: allV.Year?.$y - 1911
                    }
                }}
                className="row mt-4"
            >
                {
                    type === '4' && (
                        <>
                            <Form.Item name="Year" initialValue={dayjs()} className="col-6" label="年度">
                                <DatePicker picker='year' placeholder="選擇年度"/>
                            </Form.Item>
                            <Form.Item name="Month" initialValue={value?.Month} className="col-6" label="月份">
                                <Select>
                                    {
                                        new Array(12).fill('').map((item, index) => (
                                            <Select.Option value={index + 1} key={index + 1}>{index + 1} 月</Select.Option>
                                        ))
                                    }
                                </Select>
                            </Form.Item>
                            <Form.Item name="UserId" className="col-md-6" label="姓名">
                                <Select>
                                    {
                                        company.employee?.map((item, index) => (
                                            <Select.Option key={index} value={item.UserId}>
                                                {item.UserName}
                                            </Select.Option>
                                        ))
                                    }
                                </Select>
                            </Form.Item>

                        </>
                    )
                }
                <Form.Item name="Goal" initialValue={value?.Goal || ''} className="" label="年度目標">
                    <Input.TextArea autoSize placeholder="輸入年度目標"/>
                </Form.Item>
                <Form.Item name="Attitude" initialValue={value?.Attitude || 0} className="col-4" label="態度">
                    <InputNumber disabled={disabled} min={0} max={10} placeholder="態度"/>
                </Form.Item>
                <Form.Item name="Efficiency" initialValue={value?.Efficiency || 0} className="col-4" label="效率">
                    <InputNumber disabled={disabled} min={0} max={10} placeholder="效率"/>
                </Form.Item>
                <Form.Item name="Professional" initialValue={value?.Professional || 0} className="col-4" label="專業">
                    <InputNumber disabled={disabled} min={0} max={10} placeholder="專業"/>
                </Form.Item>
                <Form.Item name="BeLate" initialValue={value?.BeLate || 0} className="col-6" label="遲到/早退">
                    <InputNumber disabled={disabled} min={0} max={100} placeholder="次數"/>
                </Form.Item>
                <Form.Item name="DayOffNotOnRule" initialValue={value?.DayOffNotOnRule || 0} className="col-6" label="未依規定請假">
                    <InputNumber disabled={disabled} min={0} max={100} placeholder="次數"/>
                </Form.Item>
                <Form.Item name="Directions" initialValue={value?.Directions || ''} className="" label="績效描述">
                    <Input.TextArea disabled={disabled} autoSize placeholder="績效描述"/>
                </Form.Item>
            </Form>
        </Modal>
    )
}
export default Edit
