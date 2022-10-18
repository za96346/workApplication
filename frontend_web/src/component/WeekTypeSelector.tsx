import { Select } from "antd"
import React from "react"

const WeekTypeSelector = (): JSX.Element => {
    return (
        <>
            <Select defaultValue={0}>
                <Select.Option value={0} key={0}>平日</Select.Option>
                <Select.Option value={1} key={1}>假日</Select.Option>
            </Select>
        </>
    )
}
export default WeekTypeSelector
