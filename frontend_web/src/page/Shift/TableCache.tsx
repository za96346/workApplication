import { Select } from 'antd'
import { v4 as uuid } from 'uuid'
import React, { useMemo } from 'react'
import { companyReducerType } from '../../reduxer/reducer/companyReducer'

const useTableCache = (company: companyReducerType): any => {
    const tb = useMemo(() => {
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
                            new Array(10).fill('').map((i, idx) => {
                                return (
                                    <tr key={uuid()}>
                                        {
                                            new Array(31).fill('').map((item, index) => {
                                                return (
                                                    index === 0
                                                        ? <td style={{ left: '1px' }} className={styles.stickyTd}>
                                                                        jack
                                                        </td>
                                                        : <td style={{ height: '10px', width: '10px' }}>

                                                            <Select suffixIcon={null}>
                                                                {
                                                                    company.banchStyle?.map((item) => {
                                                                        return (
                                                                            <Select.Option value={item.Icon} key={item.StyleId}>{item.Icon}</Select.Option>
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
    }, [company])
    return {
        tb
    }
}
export default useTableCache
