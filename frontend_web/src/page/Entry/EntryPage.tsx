import React, { useEffect } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import api from '../../api/api'
import useReduceing from '../../Hook/useReducing'
import ForgetPwdForm from './ForgetPwdForm'
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
                {
                    path !== 'forgetPwd' && (
                        <div>
                            <div className={path === 'login' ? styles.tabActive : ''} onClick={() => navigate('/entry/login')}>登入</div>
                            <div className={path === 'register' ? styles.tabActive : ''} onClick={() => navigate('/entry/register')}>註冊</div>
                        </div>
                    )
                }
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
                {
                    path !== 'forgetPwd' && (
                        <div onClick={() => navigate('/entry/forgetPwd')} className={styles.forgetPwd}>忘記密碼?</div>
                    )
                }
                {
                    path === 'forgetPwd' && (
                        <>
                            <ForgetPwdForm />
                            <div onClick={() => navigate('/entry/login')} className={styles.forgetPwd}>返回登入畫面</div>
                        </>
                    )
                }

            </div>
        </>
    )
}

export default Entry
