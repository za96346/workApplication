import { useAppSelector } from 'hook/redux'
import React, { useEffect } from 'react'
import { v4 } from 'uuid'

const PrintList = (): JSX.Element => {
    const performance = useAppSelector((v) => v?.performance?.all)
    useEffect(() => {
        setTimeout(() => {
            print()
        }, 1000)
    }, [])
    return (
        <div translate='no' className={'print_page'}>
            <table key={v4()} className="table table-striped">
                <thead className="thead-dark">
                    <tr>
                        <th scope="col">姓名</th>
                        <th scope="col">年月份</th>
                        <th scope="col">組別</th>
                        <th scope="col">年度目標</th>
                        <th scope="col">態度</th>
                        <th scope="col">效率</th>
                        <th scope="col">專業</th>
                        <th scope="col">遲到/早退</th>
                        <th scope="col">未依規定請假</th>
                        <th scope="col">績效描述</th>
                    </tr>
                </thead>
                <tbody>
                    {
                        performance.map((item) => {
                            return (
                                <tr
                                    key={v4()}
                                >
                                    <th>{item?.UserName}</th>
                                    <th >{item.Year}/{item.Month}</th>
                                    <th >{item?.BanchName}</th>
                                    <th >
                                        {
                                            item.Goal?.split('\n')?.map((i) =>
                                                <p className='m-0' key={v4()}>{i}</p>
                                            )
                                        }
                                    </th>
                                    <th >{item.Attitude}</th>
                                    <th >{item.Efficiency}</th>
                                    <th >{item.Professional}</th>
                                    <th >{item.BeLate}</th>
                                    <th >{item.DayOffNotOnRule}</th>
                                    <th >
                                        {
                                            item.Directions?.split('\n')?.map((i) =>
                                                <p className='m-0' key={v4()}>{i}</p>
                                            )
                                        }
                                    </th>
                                </tr>
                            )
                        })
                    }
                </tbody>
            </table>
        </div>
    )
}
export default PrintList
