import { Select, Spin } from 'antd'
import { v4 as uuid } from 'uuid'
import React, { useMemo, useState } from 'react'
import { companyReducerType } from '../../reduxer/reducer/companyReducer'
import { useWebsocket } from '../../Hook/useWebsocket'
import { userReducerType } from '../../reduxer/reducer/userReducer'

const useTableCache = (company: companyReducerType, banchId: number, user: userReducerType): any => {
    const { connectionStatus, sendMessage, lastJsonMessage } = useWebsocket({
        options: {
            onClose: (event: any) => {
                console.log(event)
            },
            Headers: {
                banchId,
                token: user?.token || ''
            }
        }
    })
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    const [status, setStatus] = useState({
        hoveredPos: -1 // date
    })
    const onMouseEnter = (pos: number): void => {
        setStatus((prev) => ({ ...prev, hoveredPos: pos }))
        sendMessage("fwef")
    }
    const onMouseLeave = (): void => {
        // setStatus(() => ())
    }
    console.log(lastJsonMessage)
    const tb = useMemo(() => {
        if (connectionStatus !== 'Connecting') {
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
                            company.employee.filter((item) => item.Banch === banchId).map((i, idx) => {
                                return (
                                    <tr key={uuid()}>
                                        {
                                            new Array(31).fill('').map((item, index) => {
                                                return (
                                                    index === 0
                                                        ? <td style={{ left: '1px' }} className={styles.stickyTd}>
                                                            {
                                                                i.UserName
                                                            }
                                                        </td>
                                                        : <td
                                                            onMouseLeave={onMouseLeave}
                                                            onMouseEnter={() => onMouseEnter(parseInt(`${idx}${index}`))}
                                                            style={{ height: '10px', width: '10px' }}
                                                        >
                                                            <Select suffixIcon={null}>
                                                                {
                                                                    company.banchStyle?.map((item) => {
                                                                        return (
                                                                            <Select.Option value={item.Icon} key={item.StyleId}>
                                                                                {item.Icon}
                                                                            </Select.Option>
                                                                        )
                                                                    })
                                                                }
                                                            </Select>

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
    }, [company, connectionStatus])
    return {
        tb
    }
}
export default useTableCache
