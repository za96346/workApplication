import { Select, SelectProps } from 'antd'
import React from 'react'
import useReduceing from '../Hook/useReducing'
import { BanchType } from '../type'

interface BanchSelectorProps extends SelectProps {
    defaultValue: number
}

const BanchSelector = ({ ...rest }: BanchSelectorProps): JSX.Element => {
    const { defaultValue, style, ...other } = rest
    const { company } = useReduceing()
    const df = company.banch?.find((item: BanchType) => item.Id === defaultValue)
    return (
        <Select allowClear style={{ ...style, minWidth: '100px' }} {...other} defaultValue={df?.Id}>
            {
                company.banch.map((item) => (
                    <Select.Option
                        key={item.Id}
                        value={item.Id}
                    >
                        {
                            item.BanchName
                        }
                    </Select.Option>
                ))
            }
        </Select>
    )
}
export default BanchSelector
