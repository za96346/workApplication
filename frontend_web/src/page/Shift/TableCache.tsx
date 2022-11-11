import { v4 as uuid } from 'uuid'
import React, { useEffect, useMemo, useState } from 'react'
import { companyReducerType } from '../../reduxer/reducer/companyReducer'
import { userReducerType } from '../../reduxer/reducer/userReducer'
import { BanchStyleType } from '../../type'
import { Spin } from 'antd'
import useShiftEditSocket from '../../Hook/useShiftEdit'
import dateHandle from '../../method/dateHandle'
import { useDispatch } from 'react-redux'
import companyAction from '../../reduxer/action/companyAction'
import statics from '../../statics'

const useTableCache = (company: companyReducerType, banchId: number, user: userReducerType): any => {
    const dispatch = useDispatch()
    const { connectionStatus, sendMessage, lastJsonMessage } = useShiftEditSocket(banchId, user?.token || '')
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    const [status, setStatus] = useState({
        clickPos: -1 // date
    })
    const dayArray = useMemo(() => {
        const start = lastJsonMessage?.StartDay
        const end = lastJsonMessage?.EndDay
        const dayContainer = ['']
        for (let i = 0; i < new Date(end).getDate(); i++) {
            const res = dateHandle.addDays(start, i)
            dayContainer.push(res)
        }
        return dayContainer
    }, [lastJsonMessage?.StartDay, lastJsonMessage?.EndDay])
    const onClickSendPosition = (pos: number): void => {
        setStatus((prev) => ({ ...prev, clickPos: pos }))
        sendMessage(JSON.stringify({
            Types: 1,
            Data: {
                MyPosition: pos
            }
        }))
    }
    const onClickSendShift = (userId: number, date: string, it: BanchStyleType): void => {
        // console.log(new Date(it.OnShiftTime))
        setStatus((prev) => ({ ...prev, clickPos: -1 }))
        sendMessage(JSON.stringify({
            Types: 2,
            Data: {
                UserId: userId,
                Date: date,
                BanchStyleId: it.StyleId
                // OnShiftTime: dateHandle.transferToUtc('2022-02-02', it.OnShiftTime),
                // OffShiftTime: dateHandle.transferToUtc('2022-02-02', it.OffShiftTime)
            }
        }))
    }
    const tb = useMemo(() => {
        if (connectionStatus !== 'Connecting' &&
            connectionStatus !== 'Open') {
            return <Spin tip={'進入編輯室中...'} />
        }
        return (
            <>
                <div className={styles.shiftTable}>
                    <table >
                        <thead>
                            {
                                dayArray.map((item, index) => {
                                    return (
                                        index === 0
                                            ? <td style={{ left: '1px' }} className={styles.stickyTd}>員工</td>
                                            : <td>
                                                {item?.slice(8, 10)}<br/>
                                                {statics.days[new Date(item).getDay()]}
                                            </td>
                                    )
                                })
                            }
                        </thead>
                        {
                            lastJsonMessage?.EditUser?.map((i, idx) => {
                                return (
                                    <tr key={uuid()}>
                                        {
                                            dayArray.map((item, index) => {
                                                const position = parseInt(`${idx * 31 + index}`)
                                                const bg = lastJsonMessage?.OnlineUser?.find((user) => user.Position === position)
                                                const findIcon = lastJsonMessage?.ShiftData?.find((shift) =>
                                                    shift.Date === item &&
                                                    shift.UserId === i.UserId
                                                )
                                                const icon = company?.banchStyle?.find((icon) => {
                                                    return icon.StyleId === findIcon?.BanchStyleId
                                                })
                                                return (
                                                    index === 0
                                                        ? <td key={uuid()} style={{ left: '1px' }} className={styles.stickyTd}>
                                                            {
                                                                i.UserName
                                                            }
                                                        </td>
                                                        : <td
                                                            className={styles.td}
                                                            key={uuid()}
                                                            onClickCapture={() => onClickSendPosition(position)}
                                                            style={{
                                                                backgroundColor: bg ? bg.Color : 'white'
                                                            }}
                                                        >
                                                            {
                                                                status.clickPos === position
                                                                    ? (
                                                                        <div className={styles.downList}>

                                                                            {
                                                                                company.banchStyle.map((it) => (
                                                                                    <div
                                                                                        onClick={(e) => {
                                                                                            e.stopPropagation()
                                                                                            onClickSendShift(i.UserId, item, it)
                                                                                        }}
                                                                                        key={uuid()}
                                                                                    >
                                                                                        {it.Icon}
                                                                                    </div>
                                                                                ))
                                                                            }

                                                                        </div>
                                                                    )
                                                                    : <>{icon?.Icon || ''}</>
                                                            }
                                                        </td>
                                                )
                                            })
                                        }
                                    </tr>
                                )
                            })
                        }
                    </table>
                </div>
            </>
        )
    }, [company, lastJsonMessage, status.clickPos, connectionStatus])

    useEffect(() => {
        dispatch(companyAction.setBanchStyle(lastJsonMessage?.BanchStyle || []))
    }, [lastJsonMessage?.BanchStyle])
    return {
        tb
    }
}
export default useTableCache
