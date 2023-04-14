import useReduceing from 'Hook/useReducing'
import React, { useCallback, useState } from 'react'
import statics from 'statics'
import { BanchStyleType } from 'type'

interface props {
    sendMessage: (v: string) => void
    dayArray: any[]
}

const Row = ({
    sendMessage = () => {},
    dayArray = []
}: props): JSX.Element => {
    const { shiftEdit } = useReduceing()
    const [status, setStatus] = useState({
        clickPos: -1 // date
    })

    // 寄送位置
    const onClickSendPosition = (pos: number): void => {
        if (shiftEdit?.State?.disabledTable) return // 擋掉不能編輯的
        setStatus((prev) => ({ ...prev, clickPos: pos }))
        sendMessage(JSON.stringify({
            Types: statics.shiftSocketEvent.position,
            Data: {
                MyPosition: pos
            }
        }))
    }

    // 寄送編輯版表
    const onClickSendShift = (userId: number, date: string, it: BanchStyleType): void => {
        if (shiftEdit?.State?.disabledTable) return // 擋掉不能編輯的
        setStatus((prev) => ({ ...prev, clickPos: -1 }))
        sendMessage(JSON.stringify({
            Types: statics.shiftSocketEvent.shift,
            Data: {
                UserId: userId,
                Date: date,
                Icon: it.Icon,
                BanchStyleId: it.StyleId,
                RestTime: it.RestTime,
                OnShiftTime: it.OnShiftTime,
                OffShiftTime: it.OffShiftTime
            }
        }))
    }

    // 班表圖標
    const iconOption = useCallback((userId: number, day: string) => {
        return <div className={`${window.styles.downList} shadow-lg bg-white rounded`}>

            {
                shiftEdit.BanchStyle?.map((it) => (
                    <div
                        onClick={(e) => {
                            e.stopPropagation()
                            onClickSendShift(userId, day, it)
                        }}
                        key={it.StyleId}
                    >
                        {it.Icon}
                    </div>
                ))
            }

        </div>
    }, [shiftEdit.BanchStyle])

    return (
        <>
            {// 人的 列
                [
                    ...shiftEdit?.EditUser,
                    { UserId: -999, UserName: '總時數' }
                ]?.map((user, idx) => {
                    return (
                        <tr style={{ cursor: shiftEdit?.State?.disabledTable ? 'not-allowed' : 'pointer' }} key={user.UserId}>
                            {// 日期的 欄
                                dayArray.map((day, index) => {
                                    const position = parseInt(`${idx * 31 + index}`)
                                    const bg = shiftEdit?.OnlineUser?.find((user) => user.Position === position)
                                    const findShift = shiftEdit?.ShiftData?.find((shift) =>
                                        (shift.Date === day || shift?.OnShiftTime?.substring(0, 10) === day) &&
                                            shift.UserId === user.UserId
                                    )

                                    const key = `${day}-${user.UserId}`
                                    // 渲染 使用者名字
                                    if (index === 0) {
                                        return (
                                            <td key={key} style={{ left: '1px' }} className={window.styles.stickyTd}>
                                                {
                                                    user.UserName
                                                }
                                            </td>
                                        )
                                    }
                                    // 渲染編輯區
                                    return (
                                        <td
                                            className={window.styles.td}
                                            key={key}
                                            onClickCapture={() => {
                                                // 這邊要擋掉總時數 欄列
                                                if (user.UserId !== -999 && day !== '總時數') {
                                                    onClickSendPosition(position)
                                                }
                                            }}
                                            style={{
                                                backgroundColor: bg ? bg.Color : 'white'
                                            }}
                                        >
                                            {
                                                day === '總時數'
                                                    ? (
                                                        <>
                                                            {shiftEdit?.RowsShiftTotal?.[user.UserId] || 0}
                                                        </>
                                                    )
                                                    : user.UserId === -999
                                                        ? (
                                                            <>
                                                                {shiftEdit?.ColumnsShiftTotal?.[day] || 0}
                                                            </>
                                                        )
                                                        : status.clickPos === position // 找到是當前點擊的位置
                                                            ? iconOption(user.UserId, day)
                                                            : <span>{findShift?.Icon || ''}</span>
                                            }

                                        </td>
                                    )
                                })
                            }
                        </tr>
                    )
                })
            }
        </>
    )
}
export default Row
