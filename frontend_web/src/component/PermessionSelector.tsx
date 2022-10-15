import { Select } from "antd"
import React from "react"
import statics from "../statics"

const PermessionSelector = (): JSX.Element => {
    return (
        <>
            <Select>
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
