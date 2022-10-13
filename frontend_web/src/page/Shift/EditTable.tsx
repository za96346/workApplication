import { Result } from 'antd'
import React from 'react'
import useTableCache from './TableCache'

interface EditTableProps {
    currentTabs: number
}
const EditTable = ({ currentTabs }: EditTableProps): JSX.Element => {
    const { tb } = useTableCache()
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
                            <div>19:00 - 21:00: <span>▲</span></div>
                            <div>08:00 - 21:00: <span>☼</span></div>

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
