import { Select, SelectProps } from "antd"
import React from "react"
import statics from "../statics"

interface props extends SelectProps {
    defaultValue: number
}
const WeekDaySelector = ({ ...rest }: props): JSX.Element => {
    const { defaultValue, ...other } = rest
    const df = statics?.weekDay[defaultValue] || statics.weekDay[1]
    return (
        <>
            <Select {...other} defaultValue={df}>
                {
                    Object.values(statics.weekDay).map((item, index) => {
                        const keys = parseInt(Object.keys(statics.weekDay)[index])
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
