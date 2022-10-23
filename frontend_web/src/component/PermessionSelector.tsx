import { Select, SelectProps } from 'antd'
import React from 'react'
import statics from '../statics'

interface props extends SelectProps {
    defaultValue: number
}

const PermessionSelector = ({ ...rest }: props): JSX.Element => {
    const { defaultValue, ...other } = rest
    return (
        <>
            <Select {...other} defaultValue={statics.permession[defaultValue]}>
                {
                    Object.values(statics.permession).map((item, index) => {
                        const v = parseInt(Object.keys(statics.permession)[index])
                        return (
                            <Select.Option value={v} key={v}>{item}</Select.Option>
                        )
                    })
                }
            </Select>
        </>
    )
}
export default PermessionSelector
