import { DatePicker } from 'antd'
import React, { useEffect, useState } from 'react'
import dayjs from 'dayjs'

interface props {
    onChange: (v: dayjs.Dayjs[]) => void
    type: 'year' | 'month'
}
const DateSelect = ({ onChange, type = 'month' }: props): JSX.Element => {
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
                picker={type}
            />
        </>
    )
}

// 西元年轉換成 民國年
DateSelect.yearTransferToZhtw = (v: number): number => (v ?? new Date().getFullYear()) - 1911

// 轉換 month 要補零
DateSelect.monthFormat = (v: number): string => {
    const result = v + 1
    return result < 10
        ? `0${result}`
        : `${result}`
}

// 轉換 antd time to 民國
DateSelect.getZhtwDate = (v: any) => {
    // 開始
    const StartDate = `${
        DateSelect.yearTransferToZhtw(v?.[0]?.year())
    }-${
        DateSelect.monthFormat((v?.[0]?.month() ?? new Date().getMonth()) as number)
    }`

    // 結束
    const EndDate = `${
        DateSelect.yearTransferToZhtw(v?.[1]?.year())
    }-${
        DateSelect.monthFormat((v?.[1]?.month() ?? new Date().getMonth()) as number)
    }`
    return {
        StartDate,
        EndDate
    }
}

// 轉換 antd time to 民國
DateSelect.getZhtwYear = (v: any) => {
    // 開始
    const StartYear = `${
        DateSelect.yearTransferToZhtw(v?.[0]?.year())
    }`

    // 結束
    const EndYear = `${
        DateSelect.yearTransferToZhtw(v?.[1]?.year())
    }`
    return {
        StartYear,
        EndYear
    }
}

export default DateSelect
