import { Divider } from 'antd'
import React, { useEffect } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import api from 'api/api'
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
        navigate('/setting/1000')
    }
    useEffect(() => {
        if (user.token) {
            nav()
        }
    }, [user.token])
    return (
        <>
            <div className={window.styles.entryBlock}>
                {
                    path !== 'forgetPwd' && (
                        <div>
                            <div className={path === 'login' ? window.styles.tabActive : ''} onClick={() => navigate('/entry/login')}>登入</div>
                            <div className={path === 'register' ? window.styles.tabActive : ''} onClick={() => navigate('/entry/register')}>註冊</div>
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
                        <div onClick={() => navigate('/entry/forgetPwd')} className={window.styles.forgetPwd}>忘記密碼?</div>
                    )
                }
                {
                    path === 'forgetPwd' && (
                        <>
                            <ForgetPwdForm />
                            <div onClick={() => navigate('/entry/login')} className={window.styles.forgetPwd}>返回登入畫面</div>
                        </>
                    )
                }
                <Divider />
                <a
                    className='btn btn-secondary'
                    onClick={async () => {
                        const res = await api.googleLogin()
                        console.log(res.data)
                        window.location.assign(res.data)
                    }}
                >
                    google 登入
                </a>
            </div>
        </>
    )
}

export default Entry
