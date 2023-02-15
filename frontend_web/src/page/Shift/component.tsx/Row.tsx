import useReduceing from 'Hook/useReducing'
import React, { useCallback, useState } from 'react'
import { BanchStyleType } from 'type'

interface props {
    sendMessage: (v: string) => void
    dayArray: any[]
}

const Row = ({
    sendMessage,
    dayArray
}: props): JSX.Element => {
    const { shiftEdit } = useReduceing()
    const [status, setStatus] = useState({
        clickPos: -1 // date
    })

    // 寄送位置
    const onClickSendPosition = (pos: number): void => {
        setStatus((prev) => ({ ...prev, clickPos: pos }))
        sendMessage(JSON.stringify({
            Types: 1,
            Data: {
                MyPosition: pos
            }
        }))
    }

    // 寄送編輯版表
    const onClickSendShift = (userId: number, date: string, it: BanchStyleType): void => {
        console.log(date, it)
        setStatus((prev) => ({ ...prev, clickPos: -1 }))
        sendMessage(JSON.stringify({
            Types: 2,
            Data: {
                UserId: userId,
                Date: date,
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
                shiftEdit?.EditUser?.map((user, idx) => {
                    return (
                        <tr key={user.UserId}>
                            {// 日期的 欄
                                dayArray.map((day, index) => {
                                    const position = parseInt(`${idx * 31 + index}`)
                                    const bg = shiftEdit?.OnlineUser?.find((user) => user.Position === position)
                                    const findIcon = shiftEdit?.ShiftData?.find((shift) =>
                                        shift.Date === day &&
                                                    shift.UserId === user.UserId
                                    )
                                    const icon = shiftEdit.BanchStyle?.find((icon) => {
                                        return icon.StyleId === findIcon?.BanchStyleId
                                    })

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
                                            onClickCapture={() => onClickSendPosition(position)}
                                            style={{
                                                backgroundColor: bg ? bg.Color : 'white'
                                            }}
                                        >
                                            {
                                                // 找到是當前點擊的位置
                                                status.clickPos === position
                                                    ? iconOption(user.UserId, day)
                                                    : <span>{icon?.Icon || ''}</span>
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
