import { Select } from "antd"
import React from "react"
import statics from "../statics"

const WeekDaySelector = (): JSX.Element => {
    return (
        <>
            <Select>
                {
                    Object.values(statics.weekDay).map((item, index) => {
                        const keys = Object.keys(statics.weekDay)[index]
                        return (
                            <Select.Option value={keys} key={keys}>{item}</Select.Option>
                        )
                    })
                }
            </Select>
        </>
    )
}
export default WeekDaySelector
