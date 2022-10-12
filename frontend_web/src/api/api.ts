import api from './apiAbs'
import axios from 'axios'
import { LoginResponse, LoginType, RegisterType, ResMessage } from '../type'
import checkStatus from './checkStatus'
import userAction from '../reduxer/action/userAction'
import { store } from '../reduxer/store'

class ApiControl extends api {
    baseUrl: string
    protected readonly route = {
        login: 'entry/login',
        getEmailCaptcha: 'entry/email/captcha',
        registe: 'entry/register'
    }

    constructor () {
        super()
        this.baseUrl = 'https:///workhard.app/workApp'
    }

    protected async GET<T> (url: any, params: any, callBack: any): Promise<{ data: T, status: boolean }> {
        try {
            const res = await axios.get(`${this.baseUrl}/${url}`, { ...params })
            callBack(res.status)
            return {
                data: res.data,
                status: true
            }
        } catch (e) {
            callBack(e.response.status)
            return {
                data: e.response.data,
                status: true
            }
        }
    }

    protected async POST<T> (url: any, body: any, callBack: any, params?: any): Promise<{ data: T, status: boolean }> {
        try {
            const res = await axios.post(`${this.baseUrl}/${url}`, body, {
                ...params
            })
            callBack(res.status)
            return {
                data: res.data,
                status: true
            }
        } catch (e) {
            callBack(e.response.status)
            return {
                data: e.response.data,
                status: false
            }
        }
    }

    protected async PUT<T> (url: any, body: any, callBack: any, params?: any): Promise<{ data: T, status: boolean }> {
        try {
            const res = await axios.put(`${this.baseUrl}/${url}`, body, {
                ...params
            })
            callBack(res.status)
            return {
                data: res.data,
                status: true
            }
        } catch (e) {
            callBack(e.response.status)
            return {
                data: e.response.data,
                status: false
            }
        }
    }

    async login (data: LoginType): Promise<void> {
        const res = await this.POST<LoginResponse>(
            this.route.login,
            data,
            async (e: number) => await checkStatus.Login(e)
        )
        console.log(res)
        if (res.status) {
            store.dispatch(userAction.setToken(res.data.token))
        }
    }

    async register (data: RegisterType): Promise<boolean> {
        const res = await this.PUT<ResMessage>(
            this.route.registe,
            data,
            async (e: number) => await checkStatus.Register(e)
        )
        return res.status
    }

    async getEmailCaptcha (email: string): Promise<void> {
        const res = await this.POST<ResMessage>(
            this.route.getEmailCaptcha,
            { Email: email },
            async (e: number) => await checkStatus.GetEmailCaptcha(e)
        )
        console.log(res)
    }
}
export default new ApiControl()
