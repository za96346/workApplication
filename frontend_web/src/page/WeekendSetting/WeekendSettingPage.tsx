import { Calendar } from 'antd'
import { v4 as uuid } from 'uuid'
import React, { ReactNode, useEffect } from 'react'
import api from '../../api/api'
import useReduceing from '../../Hook/useReducing'
import dateHandle from '../../method/dateHandle'

const WeekendSettingPage = (): JSX.Element => {
    const { company } = useReduceing()
    // const [weekend, setWeekend] = useState<Array<WeekendSettingType['Date']>>([])
    const isWeekend = (d: moment.Moment): ReactNode => {
        const dateToString = d.format('YYYY-MM-DD')
        return company.weekendSetting?.map((item) => {
            const trans = dateHandle.formatDate(new Date(item.Date))
            if (trans === dateToString) {
                return (
                    <div
                        key={uuid()}
                        style={{
                            color: 'red',
                            display: 'flex',
                            justifyContent: 'flex-end'
                        }}>
                        假日
                    </div>
                )
            }
            return <></>
        })
    }
    const onSelect = async (v: moment.Moment): Promise<void> => {
        const dateToString = v.format('YYYY-MM-DD')
        const isExited = company.weekendSetting?.filter((item) => {
            const trans = dateHandle.formatDate(new Date(item.Date))
            return trans === dateToString
        })
        if (isExited?.length === 0) {
            // 新增
            await api.createWeekendSetting(dateToString)
        } else {
            // 刪除
            await api.deleteWeekendSetting(isExited[0].WeekendId)
        }
        await api.getWeekendSetting()
    }
    useEffect(() => {
        api.getWeekendSetting()
    }, [])
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
