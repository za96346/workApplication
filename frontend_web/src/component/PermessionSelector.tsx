import { Select, SelectProps } from "antd"
import React from "react"
import statics from "../statics"

interface props extends SelectProps {
    defaultValue: number
}

const PermessionSelector = ({ ...rest }: props): JSX.Element => {
    const { defaultValue, ...other } = rest
    const df = statics.permession[defaultValue || 2]
    return (
        <>
            <Select {...other} defaultValue={df}>
                {
                    Object.values(statics.permession).map((item, index) =>
                        <Select.Option value={item} key={index}>{item}</Select.Option>
                    )
                }
            </Select>
        </>
    )
}
export default PermessionSelector
