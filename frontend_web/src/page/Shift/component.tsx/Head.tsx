import useReduceing from 'Hook/useReducing'
import React from 'react'
import statics from 'statics'

interface props {
    dayArray: any[]
}

const Head = ({
    dayArray
}: props): JSX.Element => {
    const { shiftEdit } = useReduceing()
    return (
        <thead>
            {
                dayArray.map((item, index) => {
                    const findWeekend = shiftEdit?.WeekendSetting?.find((weekend) => weekend?.Date === item)
                    return (
                        index === 0
                            ? <td style={{ left: '1px' }} className={window.styles.stickyTd}>員工</td>
                            : <td style={{ backgroundColor: findWeekend ? 'rgba(255, 0, 0, 0.4)' : 'white' }}>
                                {item?.slice(8, 10)}<br/>
                                {statics.days[new Date(item).getDay()]}
                            </td>
                    )
                })
            }
        </thead>
    )
}
export default Head
