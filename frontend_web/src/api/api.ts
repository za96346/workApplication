import api from './apiAbs'
import axios from 'axios'
import { BanchRuleType, BanchStyleType, BanchType, CompanyType, LoginType, RegisterType, ResMessage, ResType, SelfDataType, UserType } from '../type'
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
        BanchAll: 'company/banch/all',
        getUserAll: 'user/all',
        getSelfData: 'user/my',
        banchStyle: 'company/banch/style',
        banchRule: 'company/banch/rule',
        companyInfo: 'company/info',
        userSingle: 'user/single'
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

    // 部門
    async getBanch (): Promise<ResType<BanchType[]>> {
        store.dispatch(statusAction.onFetchBanch(true))
        const res = await this.GET<BanchType[]>({
            url: this.route.BanchAll,
            successShow: false
        })
        if (res.status) {
            store.dispatch(companyAction.setBanch(res.data))
        }
        store.dispatch(statusAction.onFetchBanch(false))
        return res
    }

    async UpdateBanch (BanchName: BanchType['BanchName'], Id: BanchType['Id']): Promise<ResType<null>> {
        store.dispatch(statusAction.onUpdateBanch(true))
        const res = await this.POST<null>({
            url: this.route.BanchAll,
            body: {
                BanchName,
                Id
            }
        })
        store.dispatch(statusAction.onUpdateBanch(false))
        return res
    }

    async createBanch (BanchName: BanchType['BanchName']): Promise<ResType<null>> {
        store.dispatch(statusAction.onCreateBanch(true))
        const res = await this.PUT<null>({
            url: this.route.BanchAll,
            body: {
                BanchName
            }
        })
        store.dispatch(statusAction.onCreateBanch(false))
        return res
    }

    // 所有員工
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

    async getSelfData (): Promise<void> {
        store.dispatch(statusAction.onFetchSelfData(true))
        const res = await this.GET<SelfDataType[]>({
            url: this.route.getSelfData
        })
        if (res.status) {
            store.dispatch(userAction.setSelfData(res.data[0]))
        }
        store.dispatch(statusAction.onFetchSelfData(false))
    }

    // 部門樣式
    async getBanchStyle (banchId: number): Promise<ResType<BanchStyleType[]>> {
        store.dispatch(statusAction.onFetchBanchStyle(true))
        const res = await this.GET<BanchStyleType[]>({
            url: this.route.banchStyle,
            params: {
                banchId
            }
        })
        if (res.status) {
            store.dispatch(companyAction.setBanchStyle(res.data))
        }
        store.dispatch(statusAction.onFetchBanchStyle(false))
        return res
    }

    async createBanchStyle (banchStyle: BanchStyleType): Promise<ResType<null>> {
        store.dispatch(statusAction.onCreateBanchStyle(true))
        const res = await this.PUT<null>({
            url: this.route.banchStyle,
            body: {
                Icon: banchStyle.Icon,
                RestTime: banchStyle.RestTime,
                TimeRangeName: banchStyle.TimeRangeName,
                OnShiftTime: banchStyle.OnShiftTime,
                OffShiftTime: banchStyle.OffShiftTime,
                BanchId: banchStyle.BanchId
            }
        })
        console.log(res)
        store.dispatch(statusAction.onCreateBanchStyle(false))
        return res
    }

    async updateBanchStyle (banchStyle: BanchStyleType): Promise<ResType<null>> {
        store.dispatch(statusAction.onUpdateBanchStyle(true))
        const res = await this.POST<null>({
            url: this.route.banchStyle,
            body: {
                Icon: banchStyle.Icon,
                RestTime: banchStyle.RestTime,
                TimeRangeName: banchStyle.TimeRangeName,
                OnShiftTime: banchStyle.OnShiftTime,
                OffShiftTime: banchStyle.OffShiftTime,
                StyleId: banchStyle.StyleId
            }
        })
        console.log(res)
        store.dispatch(statusAction.onUpdateBanchStyle(false))
        return res
    }

    // 部門規則
    async getBanchRule (banchId: number): Promise<ResType<BanchRuleType[]>> {
        store.dispatch(statusAction.onFetchBanchRule(true))
        const res = await this.GET<BanchRuleType[]>({
            url: this.route.banchRule,
            params: {
                banchId
            }
        })
        if (res.status) {
            store.dispatch(companyAction.setBanchRule(res.data))
        }

        store.dispatch(statusAction.onFetchBanchRule(false))
        return res
    }

    async createBanchRule (banchRule: BanchRuleType): Promise<ResType<null>> {
        store.dispatch(statusAction.onCreateBanchRule(true))
        const res = await this.PUT<null>({
            url: this.route.banchRule,
            body: {
                BanchId: banchRule.BanchId,
                MinPeople: banchRule.MinPeople,
                MaxPeople: banchRule.MaxPeople,
                WeekDay: banchRule.WeekDay,
                WeekType: banchRule.WeekType,
                OnShiftTime: banchRule.OnShiftTime,
                OffShiftTime: banchRule.OffShiftTime
            }
        })
        console.log(res)
        store.dispatch(statusAction.onCreateBanchRule(false))
        return res
    }

    async updateBanchRule (banchRule: BanchRuleType): Promise<ResType<null>> {
        store.dispatch(statusAction.onUpdateBanchRule(true))
        const res = await this.POST<null>({
            url: this.route.banchRule,
            body: {
                MinPeople: banchRule.MinPeople,
                MaxPeople: banchRule.MaxPeople,
                WeekDay: banchRule.WeekDay,
                WeekType: banchRule.WeekType,
                OnShiftTime: banchRule.OnShiftTime,
                OffShiftTime: banchRule.OffShiftTime,
                RuleId: banchRule.RuleId
            }
        })
        console.log(res)
        store.dispatch(statusAction.onUpdateBanchRule(false))
        return res
    }

    // 公司資料
    async getCompanyInfo (): Promise<ResType<CompanyType>> {
        store.dispatch(statusAction.onFetchCompany(true))
        const res = await this.GET<CompanyType>({
            url: this.route.companyInfo
        })
        if (res.status && res?.data) {
            store.dispatch(companyAction.setCompany(res.data))
        }
        store.dispatch(statusAction.onFetchCompany(false))
        return res
    }

    async updateCompanyInfo (company: CompanyType): Promise<ResType<null>> {
        store.dispatch(statusAction.onUpdateCompany(true))
        const res = await this.POST<null>({
            url: this.route.companyInfo,
            body: {
                CompanyId: company.CompanyId,
                CompanyName: company.CompanyName,
                CompanyLocation: company.CompanyLocation,
                CompanyPhoneNumber: company.CompanyPhoneNumber,
                TermStart: new Date().toISOString(),
                TermEnd: new Date().toISOString()
            }
        })
        store.dispatch(statusAction.onUpdateCompany(false))
        return res
    }

    // 使用者資料
    async updateUser (user: UserType): Promise<ResType<null>> {
        const res = await this.POST<null>({
            url: this.route.userSingle,
            body: {
                ...user,
                CompanyCode: user.CompanyCode,
                EmployeeNumber: user.EmployeeNumber,
                OnWorkDay: user.OnWorkDay,
                Banch: user.Banch,
                Permession: user.Permession,
                WorkState: user.WorkState
            },
            params: {
                userId: user.UserId
            }
        })
        return res
    }
}
export default new ApiControl()
