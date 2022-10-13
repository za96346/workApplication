import { Result, Select } from 'antd'
import React from 'react'

const EditTable = (): JSX.Element => {
    return (
        <>
            {
                false && (
                    <Result
                        status="404"
                        title="404"
                        subTitle="找不到資料"
                    // extra={<Button type="primary">Back Home</Button>}
                    />
                )
            }
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
                                <tr key={idx}>
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
                                                                <div>
                                                                    <Select suffixIcon={null}>
                                                                        <Select.Option value={'*'} key={0}>*</Select.Option>
                                                                        <Select.Option value={'%'} key={1}>%</Select.Option>
                                                                        <Select.Option value={'#'} key={2}>#</Select.Option>
                                                                    </Select>
                                                                </div>
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
}
export default EditTable
