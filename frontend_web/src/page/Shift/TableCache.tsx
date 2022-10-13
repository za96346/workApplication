import { Select } from 'antd'
import { v4 as uuid } from 'uuid'
import React, { useMemo } from 'react'

const useTableCache = (): any => {
    const tb = useMemo(() => {
        return (
            <>
                <div className={styles.shiftTable}>
                    <table >
                        <thead>
                            {
                                new Array(31).fill('').map((item, index) => {
                                    return (
                                        <>
                                            {
                                                index === 0
                                                    ? <td style={{ left: '1px' }} className={styles.stickyTd}>員工</td>
                                                    : <td>
                                                        {index}<br/>
                                                        {'一'}
                                                    </td>

                                            }
                                        </>
                                    )
                                })
                            }
                        </thead>
                        {
                            new Array(10).fill('').map((i, idx) => {
                                return (
                                    <tr key={uuid()}>
                                        {
                                            new Array(31).fill('').map((item, index) => {
                                                return (
                                                    <>
                                                        {
                                                            index === 0
                                                                ? <td style={{ left: '1px' }} className={styles.stickyTd}>
                                                                        jack
                                                                </td>
                                                                : <td style={{ height: '10px', width: '10px' }}>

                                                                    <Select suffixIcon={null}>
                                                                        <Select.Option value={'*'} key={0}>*</Select.Option>
                                                                        <Select.Option value={'%'} key={1}>%</Select.Option>
                                                                        <Select.Option value={'#'} key={2}>#</Select.Option>
                                                                    </Select>

                                                                </td>
                                                        }

                                                    </>
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
    }, [])
    return {
        tb
    }
}
export default useTableCache
