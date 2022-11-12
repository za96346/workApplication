import { Result, Skeleton, Spin } from 'antd'
import { v4 as uuid } from 'uuid'
import React, { useEffect } from 'react'
import useReduceing from '../../Hook/useReducing'
import useTableCache from './TableCache'
import { useNavigate } from 'react-router-dom'

interface EditTableProps {
    currentTabs: number
    banchId: number
}
const EditTable = ({ banchId, currentTabs }: EditTableProps): JSX.Element => {
    const navigate = useNavigate()
    const { loading, user, company } = useReduceing()
    const { tb, lonelyShift, close } = useTableCache(company, banchId, user)
    useEffect(() => {
        if (currentTabs !== 0) {
            close()
        }
    }, [currentTabs])
    useEffect(() => {
        return () => close()
    }, [navigate])
    if (loading.onFetchBanchStyle || loading.onFetchUserAll) {
        return (
            <>
                <Spin />
                <Skeleton active />
            </>
        )
    }
    return (
        <>
            {
                currentTabs === 1 && (
                    <>search bar</>
                )
            }
            {
                currentTabs === 1 && (
                    <Result
                        status="404"
                        title="404"
                        subTitle="找不到資料"
                    // extra={<Button type="primary">Back Home</Button>}
                    />
                )
            }
            {
                currentTabs === 0 && (
                    <>
                        <div className={styles.shiftSignBlock}>
                            {
                                company.banchStyle?.map((item) => {
                                    return (
                                        <div key={item.StyleId}>
                                            <div>{item.TimeRangeName}</div>
                                            {item.OnShiftTime} - {item.OffShiftTime}:
                                            <span>{item.Icon}</span>
                                        </div>
                                    )
                                })
                            }
                        </div>
                        {
                            tb
                        }
                        <div >
                            {
                                lonelyShift?.map((item) => {
                                    return (
                                        <span key={uuid()}>
                                            上班: {item.OnShiftTime}<br/>
                                            下班: {item.OffShiftTime}<br/>
                                            休息: {item.RestTime}
                                        </span>
                                    )
                                })
                            }
                        </div>

                    </>
                )
            }
        </>
    )
}
export default EditTable
