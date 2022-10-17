import React, { useEffect } from 'react'
import { useSelector } from 'react-redux'
import { useNavigate, useParams } from 'react-router-dom'
import { userReducerType } from '../../reduxer/reducer/userReducer'
import { RootState } from '../../reduxer/store'
import LoginForm from './LoginForm'
import RegisterForm from './RegisterForm'

const Entry = (): JSX.Element => {
    const user: userReducerType = useSelector((state: RootState) => state.user)
    const navigate = useNavigate()
    const { path } = useParams()
    useEffect(() => {
        if (user.token) {
            navigate('/shift')
        }
    }, [user.token])
    return (
        <>
            <div className={styles.entryBlock}>
                <div>
                    <div className={path === 'login' ? styles.tabActive : ''} onClick={() => navigate('/entry/login')}>登入</div>
                    <div className={path === 'register' ? styles.tabActive : ''} onClick={() => navigate('/entry/register')}>註冊</div>
                </div>
                {
                    path === 'login' && (
                        <LoginForm />
                    )
                }
                {
                    path === 'register' && (
                        <RegisterForm />
                    )
                }
            </div>
        </>
    )
}

export default Entry
