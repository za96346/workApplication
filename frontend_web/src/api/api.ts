import api from './apiAbs'
import axios from 'axios'
import { BanchType, LoginType, RegisterType, ResMessage, ResType } from '../type'
import userAction from '../reduxer/action/userAction'
import { store } from '../reduxer/store'
import { FullMessage } from '../method/notice'
import companyAction from '../reduxer/action/companyAction'

class ApiControl extends api {
    baseUrl: string
    protected readonly route = {
        login: 'entry/login',
        getEmailCaptcha: 'entry/email/captcha',
        registe: 'entry/register',
        getBanchAll: 'company/banch/all'
    }

    constructor () {
        super()
        this.baseUrl = process.env.REACT_APP_API
    }

    protected async GET<T> ({
        url,
        params = {},
        successShow = true,
        FailShow = true
    }: {
        url: any
        params?: any
        successShow?: boolean
        FailShow?: boolean
    }
    ): Promise<ResType<T>> {
        const { token } = store.getState().user
        try {
            const res = await axios.get(`${this.baseUrl}/${url}`, { params: { ...params, token } })
            if (successShow) {
                void FullMessage.success(res.data.message)
            }
            return {
                ...res.data,
                status: true
            }
        } catch (e) {
            if (FailShow) {
                void FullMessage.error(e.response.data.message)
            }
            return {
                ...e.response.data,
                status: true
            }
        }
    }

    protected async POST<T> ({
        url,
        body,
        params = {},
        successShow = true,
        FailShow = true
    }: {
        url: any
        body: any
        params?: any
        successShow?: boolean
        FailShow?: boolean
    }): Promise<ResType<T>> {
        const { token } = store.getState().user
        try {
            const res = await axios.post(`${this.baseUrl}/${url}`, body, {
                headers: {
                    token
                },
                ...params
            })
            if (successShow) {
                void FullMessage.success(res.data.message)
            }
            return {
                ...res.data,
                status: true
            }
        } catch (e) {
            if (FailShow) {
                void FullMessage.error(e.response.data.message)
            }
            return {
                ...e.response.data,
                status: false
            }
        }
    }

    protected async PUT<T> ({
        url,
        body,
        params = {}
    }: {
        url: any
        body: any
        params?: any
    }): Promise<ResType<T>> {
        const { token } = store.getState().user
        try {
            const res = await axios.put(`${this.baseUrl}/${url}`, body, {
                headers: {
                    token
                },
                ...params
            })
            void FullMessage.success(res.data.message)
            return {
                ...res.data,
                status: true
            }
        } catch (e) {
            void FullMessage.error(e.response.data.message)
            return {
                ...e.response.data,
                status: false
            }
        }
    }

    async login (data: LoginType): Promise<void> {
        const res = await this.POST<string>({
            url: this.route.login,
            body: data
        })
        console.log(res)
        if (res.status) {
            store.dispatch(userAction.setToken(res.data))
        }
    }

    async register (data: RegisterType): Promise<boolean> {
        const res = await this.PUT<ResMessage>({
            url: this.route.registe,
            body: data
        })
        return res.status
    }

    async getEmailCaptcha (email: string): Promise<void> {
        const res = await this.POST<ResMessage>({
            url: this.route.getEmailCaptcha,
            body: { Email: email }
        })
        console.log(res)
    }

    async getBanch (): Promise<ResType<BanchType[]>> {
        const res = await this.GET<BanchType[]>({
            url: this.route.getBanchAll,
            successShow: false
        })
        store.dispatch(companyAction.setBanch(res.data))
        console.log(res)
        return res
    }
}
export default new ApiControl()
