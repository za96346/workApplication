import { Select, SelectProps } from "antd"
import React from "react"
import statics from "../statics"

interface props extends SelectProps {
    defaultValue: 1 | 2
}

const WeekTypeSelector = ({ ...rest }: props): JSX.Element => {
    const { defaultValue, ...other } = rest
    const df = statics?.weekType[defaultValue]
    return (
        <>
            <Select {...other} defaultValue={df}>
                <Select.Option value={1} key={1}>平日</Select.Option>
                <Select.Option value={2} key={2}>假日</Select.Option>
            </Select>
        </>
    )
}
export default WeekTypeSelector
