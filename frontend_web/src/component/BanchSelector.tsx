import { Select } from 'antd'
import React from 'react'
import useReduceing from '../Hook/useReducing'
import { BanchType } from '../type'

interface BanchSelectorProps {
    defaultValue: number
}

const BanchSelector = ({ defaultValue }: BanchSelectorProps): JSX.Element => {
    const { company } = useReduceing()
    const df = company.banch?.find((item: BanchType) => item.Id === defaultValue)
    console.log(df, defaultValue)
    return (
        <Select defaultValue={df?.Id}>
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
