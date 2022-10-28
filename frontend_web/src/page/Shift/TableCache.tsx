import { v4 as uuid } from 'uuid'
import React, { useMemo, useState } from 'react'
import { companyReducerType } from '../../reduxer/reducer/companyReducer'
import { userReducerType } from '../../reduxer/reducer/userReducer'
import { BanchStyleType } from '../../type'
import { Spin } from 'antd'
import useShiftEditSocket from '../../Hook/useShiftEdit'

const useTableCache = (company: companyReducerType, banchId: number, user: userReducerType): any => {
    const { connectionStatus, sendMessage, lastJsonMessage } = useShiftEditSocket(banchId, user?.token || '')
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    const [status, setStatus] = useState({
        clickPos: -1 // date
    })
    const onClickSendPosition = (pos: number): void => {
        setStatus((prev) => ({ ...prev, clickPos: pos }))
        sendMessage(JSON.stringify({
            Types: 1,
            Data: {
                MyPosition: pos
            }
        }))
    }
    const onClickSendShift = (userId: number, pos: number, it: BanchStyleType): void => {
        // console.log(new Date(it.OnShiftTime))
        setStatus((prev) => ({ ...prev, clickPos: -1 }))
        sendMessage(JSON.stringify({
            Types: 2,
            Data: {
                UserId: userId,
                Position: pos,
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
                                new Array(31).fill('').map((item, index) => {
                                    return (
                                        index === 0
                                            ? <td style={{ left: '1px' }} className={styles.stickyTd}>員工</td>
                                            : <td>
                                                {index}<br/>
                                                {'一'}
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
                                            new Array(31).fill('').map((item, index) => {
                                                const position = parseInt(`${idx * 31 + index}`)
                                                const bg = lastJsonMessage?.OnlineUser?.find((user) => user.Position === position)
                                                const findIcon = lastJsonMessage?.ShiftData?.find((shift) => shift.Position === position)
                                                const icon = company?.banchStyle?.find((icon) => {
                                                    // const on = dateHandle.dateFormatToTime(new Date(findIcon?.OnShiftTime))
                                                    // const off = dateHandle.dateFormatToTime(new Date(findIcon?.OffShiftTime))
                                                    // if (findIcon?.OnShiftTime && findIcon?.OffShiftTime) {
                                                    //     console.log(on, findIcon.OnShiftTime)
                                                    // }
                                                    // return (
                                                    //     icon?.OnShiftTime === on &&
                                                    //     icon?.OffShiftTime === off
                                                    // )
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
                                                                                            onClickSendShift(i.UserId, position, it)
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
    return {
        tb
    }
}
export default useTableCache
