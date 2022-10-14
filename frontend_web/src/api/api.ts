import api from './apiAbs'
import axios from 'axios'
import { LoginResponse, LoginType, RegisterType, ResMessage } from '../type'
import userAction from '../reduxer/action/userAction'
import { store } from '../reduxer/store'
import { FullMessage } from '../method/notice'

class ApiControl extends api {
    baseUrl: string
    protected readonly route = {
        login: 'entry/login',
        getEmailCaptcha: 'entry/email/captcha',
        registe: 'entry/register'
    }

    constructor () {
        super()
        this.baseUrl = process.env.REACT_APP_API
    }

    protected async GET<T> (url: any, params: any): Promise<{ data: T, status: boolean }> {
        try {
            const res = await axios.get(`${this.baseUrl}/${url}`, { ...params })
            void FullMessage.success(res.data.message)
            return {
                data: res.data,
                status: true
            }
        } catch (e) {
            void FullMessage.error(e.response.data.message)
            return {
                data: e.response.data,
                status: true
            }
        }
    }

    protected async POST<T> (url: any, body: any, params?: any): Promise<{ data: T, status: boolean }> {
        try {
            const res = await axios.post(`${this.baseUrl}/${url}`, body, {
                ...params
            })
            void FullMessage.success(res.data.message)
            return {
                data: res.data,
                status: true
            }
        } catch (e) {
            void FullMessage.error(e.response.data.message)
            return {
                data: e.response.data,
                status: false
            }
        }
    }

    protected async PUT<T> (url: any, body: any, params?: any): Promise<{ data: T, status: boolean }> {
        try {
            const res = await axios.put(`${this.baseUrl}/${url}`, body, {
                ...params
            })
            void FullMessage.success(res.data.message)
            return {
                data: res.data,
                status: true
            }
        } catch (e) {
            void FullMessage.error(e.response.data.message)
            return {
                data: e.response.data,
                status: false
            }
        }
    }

    async login (data: LoginType): Promise<void> {
        const res = await this.POST<LoginResponse>(
            this.route.login,
            data
        )
        console.log(res)
        if (res.status) {
            store.dispatch(userAction.setToken(res.data.token))
        }
    }

    async register (data: RegisterType): Promise<boolean> {
        const res = await this.PUT<ResMessage>(
            this.route.registe,
            data
        )
        return res.status
    }

    async getEmailCaptcha (email: string): Promise<void> {
        const res = await this.POST<ResMessage>(
            this.route.getEmailCaptcha,
            { Email: email }
        )
        console.log(res)
    }
}
export default new ApiControl()
