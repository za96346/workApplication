import { Calendar } from 'antd'
import React, { ReactNode, useState } from 'react'

const WeekendSettingPage = (): JSX.Element => {
    const [weekend, setWeekend] = useState<string[]>([])
    const isWeekend = (d: moment.Moment): ReactNode => {
        const dateToString = d.format('YYYY-MM-DD')
        return weekend?.map((item) => {
            if (item === dateToString) {
                return (
                    <div key={item} style={{ color: 'red', display: 'flex', justifyContent: 'flex-end' }}>假日</div>
                )
            }
            return <></>
        })
    }
    const onSelect = (v: moment.Moment): void => {
        const dateToString = v.format('YYYY-MM-DD')
        const isExited = weekend.indexOf(dateToString)
        if (isExited === -1) {
            setWeekend((prev) => ([...prev, dateToString]))
        } else {
            setWeekend((prev) => {
                prev.splice(isExited, 1)
                return prev
            })
        }
    }
    return (
        <>
            <Calendar
                dateCellRender={isWeekend}
                // headerRender={() => {
                //     return (
                //         <div style={{ color: 'rgb(195, 88 , 34)', fontSize: '1.5rem', marginLeft: '10px' }}>
                //             平假日設定
                //         </div>
                //     )
                // }}
                onSelect={onSelect}
            />
        </>
    )
}
export default WeekendSettingPage
f