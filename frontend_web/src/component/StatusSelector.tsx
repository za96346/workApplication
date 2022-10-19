import { Select, SelectProps } from 'antd'
import React from 'react'

interface props extends SelectProps {
    defaultValue: 'on' | 'off' | string
}

const StatusSelector = ({ ...rest }: props): JSX.Element => {
    const { defaultValue, ...other } = rest
    return (
        <Select {...other} defaultValue={defaultValue}>
            <Select.Option value={'on'} key={'on'}>在職</Select.Option>
            <Select.Option value={'off'} key={'off'}>離職</Select.Option>
        </Select>
    )
}

export default StatusSelector
