import { Tabs } from 'antd'
import React, { useEffect, useState } from 'react'
import { useSelector } from 'react-redux'
import { useLocation, useNavigate } from 'react-router-dom'
import LoginForm from './LoginForm'
import RegisterForm from './RegisterForm'

const Entry = (): JSX.Element => {
    const { token } = useSelector((state: any) => state.user)
    const navigate = useNavigate()

    const [status, setStatus] = useState({
        currentTab: '登入'
    })

    const { pathname } = useLocation()
    console.log(pathname, status)
    useEffect(() => {
        if (token) {
            navigate('/shift')
        }
    }, [token])
    return (
        <>
            <div className={styles.entryBlock}>
                <Tabs onChange={(v) => setStatus((prev) => ({ ...prev, currentTab: v }))}>
                    <Tabs.TabPane tab="登入" key={1}>
                        <LoginForm />
                    </Tabs.TabPane>
                    <Tabs.TabPane tab="註冊" key={2}>
                        <RegisterForm />
                    </Tabs.TabPane>
                </Tabs>
            </div>
        </>
    )
}

export default Entry
