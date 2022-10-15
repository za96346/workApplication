import { Select } from 'antd'
import React from 'react'

interface props {
    defaultValue: 'on' | 'off' | string
}

const StatusSelector = ({ defaultValue }: props): JSX.Element => {
    return (
        <Select defaultValue={defaultValue}>
            <Select.Option value={'on'} key={'on'}>在職</Select.Option>
            <Select.Option value={'off'} key={'off'}>離職</Select.Option>
        </Select>
    )
}

export default StatusSelector
