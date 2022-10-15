import api from './apiAbs'
import axios from 'axios'
import { BanchType, LoginType, RegisterType, ResMessage, ResType, SelfDataType, UserType } from '../type'
import userAction from '../reduxer/action/userAction'
import { store } from '../reduxer/store'
import { FullMessage } from '../method/notice'
import companyAction from '../reduxer/action/companyAction'
import statusAction from '../reduxer/action/statusAction'
import { clearAll } from '../reduxer/clearAll'

class ApiControl extends api {
    baseUrl: string
    protected readonly route = {
        login: 'entry/login',
        getEmailCaptcha: 'entry/email/captcha',
        registe: 'entry/register',
        getBanchAll: 'company/banch/all',
        getUserAll: 'user/all',
        getSelfData: 'user/my'
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
            if (e.response.status >= 510) {
                clearAll()
            }
            if (FailShow) {
                void FullMessage.error(e.response.data.message)
            }
            return {
                ...e.response.data,
                status: false
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
            if (e.response.status >= 510) {
                clearAll()
            }
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
            if (e.response.status >= 510) {
                clearAll()
            }
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

    async getEmailCaptcha (email: string): Promise<ResMessage> {
        const res = await this.POST<ResMessage>({
            url: this.route.getEmailCaptcha,
            body: { Email: email }
        })
        return res
    }

    async getBanch (): Promise<ResType<BanchType[]>> {
        store.dispatch(statusAction.onFetchBanch(true))
        const res = await this.GET<BanchType[]>({
            url: this.route.getBanchAll,
            successShow: false
        })
        if (res.status) {
            store.dispatch(companyAction.setBanch(res.data))
        }
        store.dispatch(statusAction.onFetchBanch(false))
        return res
    }

    async getUserAll (): Promise<ResType<UserType[]>> {
        store.dispatch(statusAction.onFetchUserAll(true))
        const res = await this.GET<UserType[]>({
            url: this.route.getUserAll,
            successShow: false
        })
        if (res.status) {
            store.dispatch(companyAction.setEmployee(res.data))
        }
        store.dispatch(statusAction.onFetchUserAll(false))
        // console.log(res)
        return res
    }

    async getSelfData (): Promise<ResType<SelfDataType>> {
        store.dispatch(statusAction.onFetchSelfData(true))
        const res = await this.GET<SelfDataType>({
            url: this.route.getSelfData
        })
        console.log(res)
        if (res.status) {
            store.dispatch(userAction.setSelfData(res.data[0]))
        }
        store.dispatch(statusAction.onFetchSelfData(false))
        return res
    }
}
export default new ApiControl()
