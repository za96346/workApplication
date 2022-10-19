import React, { useEffect } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import api from '../../api/api'
import useReduceing from '../../Hook/useReducing'
import LoginForm from './LoginForm'
import RegisterForm from './RegisterForm'

const Entry = (): JSX.Element => {
    const { user } = useReduceing()
    const navigate = useNavigate()
    const { path } = useParams()
    const nav = async (): Promise<void> => {
        await api.getBanch()
        navigate('/home')
    }
    useEffect(() => {
        if (user.token) {
            nav()
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
