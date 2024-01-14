import React from 'react'
import { Form, FormInstance, Switch as AntdSwitch, SwitchProps } from 'antd'

interface props {
    formInstance: FormInstance<any>
    antdSwitchProps?: SwitchProps
    onValue?: any
    offValue?: any
    defaultValue: any
    fieldName: string
    label: string
}

const Switch = ({
    formInstance,
    antdSwitchProps,
    defaultValue,
    onValue = 'Y',
    offValue = 'N',
    fieldName,
    label
}: props): JSX.Element => {
    return (
        <>
            <Form.Item
                name={fieldName}
                className='d-none'
                rules={[{ required: true }]}
                initialValue={defaultValue}
            />
            <Form.Item
                name={`${fieldName}Status`}
                label={label}
                initialValue={defaultValue === onValue}
            >
                <AntdSwitch
                    defaultChecked={defaultValue === onValue}
                    onChange={(v) => {
                        formInstance.setFieldValue(
                            fieldName,
                            v ? onValue : offValue
                        )
                    }}
                    {...antdSwitchProps}
                />
            </Form.Item>
        </>
    )
}
export default Switch
