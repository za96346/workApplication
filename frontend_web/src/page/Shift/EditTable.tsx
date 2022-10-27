import { Result, Skeleton, Spin } from 'antd'
import React, { useEffect } from 'react'
import api from '../../api/api'
import useReduceing from '../../Hook/useReducing'
import useTableCache from './TableCache'

interface EditTableProps {
    currentTabs: number
    banchId: number
}
const EditTable = ({ banchId, currentTabs }: EditTableProps): JSX.Element => {
    const { loading, user, company } = useReduceing()
    const { tb } = useTableCache(company, banchId, user)
    useEffect(() => {
        api.getBanchStyle(banchId)
        api.getUserAll()
    }, [banchId])
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

                    </>
                )
            }
        </>
    )
}
export default EditTable
