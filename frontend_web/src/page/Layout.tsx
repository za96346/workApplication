import Menu from '../component/Menu'
import React, { useEffect } from 'react'
import { Outlet, useLocation, useNavigate } from 'react-router-dom'
import useReduceing from '../Hook/useReducing'
// import { BanchRuleType, BanchStyleType } from '../type'

// const banchStyleSimulate: BanchStyleType = {
//     Icon: "mm", // 時段圖標
//     TimeRangeName: "time", // 時段名稱
//     OnShiftTime: `${new Date()}`, // 開始上班時間
//     OffShiftTime: `${new Date()}`, // 結束上班的時間
//     BanchId: 1, // 部門id
//     StyleId: 100
// }

// const banchRuleSimulate: BanchRuleType = {
//     BanchId: 1,
//     RuleId: 100,
//     MinPeople: 1,
//     MaxPeople: 3,
//     WeekDay: 1,
//     WeekType: 2,
//     OnShiftTime: `${new Date()}`,
//     OffShiftTime: `${new Date()}`
// }

const Layout = (): JSX.Element => {
    const { pathname } = useLocation()
    const navigate = useNavigate()
    const { user } = useReduceing()

    useEffect(() => {
        if (!user.token) {
            navigate('/entry/login')
        }
    }, [user.token])

    if (pathname.includes('entry')) {
        return (
            <div className={styles.entryLayOut}>
                <Outlet />
            </div>
        )
    }
    return (
        <>
            <div className={styles.article}>
                <Menu />
                <div className={styles.rightBlock}>
                    <Outlet />
                </div>
            </div>
        </>
    )
}
export default Layout
