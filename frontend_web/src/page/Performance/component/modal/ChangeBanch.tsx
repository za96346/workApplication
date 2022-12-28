import { Form, Modal } from 'antd'
import React, { useRef, useEffect } from 'react'
import BanchSelector from 'Share/BanchSelector'
import { performanceType } from 'type'
import useReduceing from 'Hook/useReducing'
import DescribeValue from '../DescribeValue'

interface props {
    open: boolean
    value?: performanceType | null
    onClose: () => void
    onSave: ((v: performanceType) => void)
}

const ChangeBanch = ({ open, value, onClose, onSave }: props): JSX.Element => {
    const form = useRef<performanceType>(value)
    const { company } = useReduceing()
    const findBanchName = (): string => company.banch.find((item) => `${item.Id}` === `${form.current?.BanchId}`)?.BanchName || ''
    useEffect(() => {
        form.current = value
    }, [value])
    return (
        <Modal
            title="更換部門"
            destroyOnClose
            open={open}
            onCancel={onClose}
            onOk={() => { onSave({ ...value, ...form.current, BanchName: findBanchName() }) }}
            okText="儲存"
            cancelText="取消"
        >
            <DescribeValue value={value}/>
            <Form onValuesChange={(v, allV) => { form.current = allV; console.log(allV) }}>
                <Form.Item initialValue={value.BanchId || ''} name="BanchId" label="選擇組別">
                    <BanchSelector defaultValue={value.BanchId} />
                </Form.Item>
            </Form>

        </Modal>
    )
}
export default ChangeBanch
