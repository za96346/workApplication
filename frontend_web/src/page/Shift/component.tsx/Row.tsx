import useReduceing from 'Hook/useReducing'
import React, { useState } from 'react'
import { BanchStyleType } from 'type'
import { v4 as uuid } from 'uuid'

interface props {
    sendMessage: (v: string) => void
    dayArray: any[]
}

const Row = ({
    sendMessage,
    dayArray
}: props): JSX.Element => {
    const { shiftEdit, company } = useReduceing()
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
        // console.log(new Date(it.OnShiftTime))
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
    return (
        <>
            {
                shiftEdit?.EditUser?.map((i, idx) => {
                    return (
                        <tr key={uuid()}>
                            {
                                dayArray.map((item, index) => {
                                    const position = parseInt(`${idx * 31 + index}`)
                                    const bg = shiftEdit?.OnlineUser?.find((user) => user.Position === position)
                                    const findIcon = shiftEdit?.ShiftData?.find((shift) =>
                                        shift.Date === item &&
                                                    shift.UserId === i.UserId
                                    )
                                    const icon = company?.banchStyle?.find((icon) => {
                                        return icon.StyleId === findIcon?.BanchStyleId
                                    })
                                    return (
                                        index === 0
                                            ? <td key={uuid()} style={{ left: '1px' }} className={window.styles.stickyTd}>
                                                {
                                                    i.UserName
                                                }
                                            </td>
                                            : <td
                                                className={window.styles.td}
                                                key={uuid()}
                                                onClickCapture={() => onClickSendPosition(position)}
                                                style={{
                                                    backgroundColor: bg ? bg.Color : 'white'
                                                }}
                                            >
                                                {
                                                    status.clickPos === position
                                                        ? (
                                                            <div className={`${window.styles.downList} shadow-lg bg-white rounded`}>

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
        </>
    )
}
export default Row
