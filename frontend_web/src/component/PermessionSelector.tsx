import { Select } from "antd"
import React from "react"
import statics from "../statics"

interface props {
    defaultValue: number
}

const PermessionSelector = ({ defaultValue }: props): JSX.Element => {
    const df = statics.permession[defaultValue || 2]
    return (
        <>
            <Select defaultValue={df}>
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
