import { Select } from 'antd'
import React from 'react'
import { useSelector } from 'react-redux'
import { BanchType } from '../type'

interface BanchSelectorProps {
    defaultValue: number
}

const BanchSelector = ({ defaultValue }: BanchSelectorProps): JSX.Element => {
    const { banch }: { banch: BanchType[] } = useSelector((state: any) => state.company)
    const df = banch.find((item: BanchType) => item.Id === defaultValue)
    console.log(df)
    return (
        <Select>
            {
                banch.map((item) => (
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
