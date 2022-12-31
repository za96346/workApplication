import { DatePicker } from 'antd'
import React, { useEffect, useState } from 'react'
import dayjs from 'dayjs'

interface props {
    onChange: (v: dayjs.Dayjs[]) => void
}
const DateSelect = ({ onChange }: props): JSX.Element => {
    const [dates, setDates] = useState<any>(null)
    const [value, setValue] = useState<any>([dayjs(), dayjs()])

    const disabledDate = (current: any): boolean => {
        if (!dates) {
            return false
        }
        // const tooLate = dates[0] && current.diff(dates[0], 'days') > 7
        const tooEarly = dates[1]
        return !!tooEarly
    }

    const onOpenChange = (open: boolean): void => {
        if (open) {
            setDates([null, null])
        } else {
            setDates(null)
        }
    }
    useEffect(() => {
        onChange(value)
    }, [value])
    return (
        <>
            <DatePicker.RangePicker
                allowClear={false}
                allowEmpty={[false, true]}
                value={dates || value}
                disabledDate={disabledDate}
                defaultValue={[dayjs(), dayjs()]}
                onCalendarChange={(val) => setDates(val)}
                placeholder={['開始', '結束']}
                onChange={(val) => setValue(val)}
                onOpenChange={onOpenChange}
                picker='month'
            />
        </>
    )
}
export default DateSelect
