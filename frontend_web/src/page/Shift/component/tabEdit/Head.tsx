import useReduceing from 'Hook/useReducing'
import React from 'react'
import statics from 'statics'

interface props {
    dayArray: any[]
}

const Head = ({
    dayArray = []
}: props): JSX.Element => {
    const { shiftEdit } = useReduceing()
    return (
        <tr>
            {
                dayArray.map((item, index) => {
                    const findWeekend = shiftEdit?.WeekendSetting?.find((weekend) => weekend?.Date === item)
                    return (
                        index === 0
                            ? <td key={item} style={{ left: '1px' }} className={window.styles.stickyTd}>員工</td>
                            : <td key={item} style={{ backgroundColor: findWeekend ? 'rgba(255, 0, 0, 0.4)' : 'white' }}>
                                {
                                    item === '總時數'
                                        ? '總時數'
                                        : (<>
                                            {item?.slice(8, 10)}<br/>
                                            {statics.days[new Date(item).getDay()]}
                                        </>)
                                }

                            </td>
                    )
                })
            }
        </tr>
    )
}
export default Head
