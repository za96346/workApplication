import apiAbs from './apiAbs'
import axios from 'axios'
import {
    BanchRuleType,
    BanchStyleType,
    BanchType,
    CompanyType,
    LoginType,
    performanceType,
    RegisterType,
    ResMessage,
    ResType,
    SelfDataType,
    shiftMonthType,
    shiftTotalType,
    UserType,
    WaitReplyType,
    WeekendSettingType,
    workTimeType,
    yearPerformanceType,
    shiftHistoryType
} from '../type'
import userAction from '../reduxer/action/userAction'
import { store } from '../reduxer/store'
import { FullMessage } from '../method/notice'
import companyAction from '../reduxer/action/companyAction'
import statusAction from '../reduxer/action/statusAction'
import { clearAll } from '../reduxer/clearAll'
import shiftEditAction from 'reduxer/action/shiftEditAction'

type createUserType = Pick<SelfDataType, 'Account' | 'Password' | 'UserName' | 'EmployeeNumber' | 'OnWorkDay' | 'Banch' | 'Permession'>
interface yearPerformanceParamsType {
    startYear: number
    endYear: number
    userName: string
}
interface shiftMonthParamsType {
    banch: number
    year: number
    month: number
}

class ApiControl extends apiAbs {
    baseUrl: string
    protected readonly route = {
        login: 'entry/login',
        getEmailCaptcha: 'entry/email/captcha',
        registe: 'entry/register',
        BanchAll: 'company/banch/all',
        getUserAll: 'user/all',
        selfData: 'user/my',
        banchStyle: 'company/banch/style',
        banchRule: 'company/banch/rule',
        companyInfo: 'company/info',
        userSingle: 'user/single',
        changePassword: 'user/changePassword',
        forgetPassword: 'user/forgetPassword',
        weekendSetting: 'company/weekend/setting',
        waitReply: 'company/wait/reply',
        workTime: 'shift/workTime',
        performance: 'pr/performance',
        performanceCopy: 'pr/performance/copy',
        yearPerformance: 'pr/performance/year',
        googleLogin: 'google/login',
        shiftMonth: 'shift/month',
        shiftTotal: 'shift/total',
        shiftHistory: 'shift/history'
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
                params: { ...params }
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
            const res = await axios.put(`${this.baseUrl}/${url}`, body, {
                headers: {
                    token
                },
                params: { ...params }
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

    protected async DELETE<T> ({
        url,
        params = {},
        successShow = true,
        FailShow = true
    }: {
        url: any
        params?: any
        successShow?: boolean
        FailShow?: boolean
    }): Promise<ResType<T>> {
        const { token } = store.getState().user
        try {
            const res = await axios.delete(`${this.baseUrl}/${url}`, {
                headers: {
                    token
                },
                params: { ...params }
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

    async login (data: LoginType): Promise<void> {
        store.dispatch(statusAction.onEntry(true))
        const res = await this.POST<string>({
            url: this.route.login,
            body: data
        })
        console.log(res)
        if (res.status) {
            store.dispatch(userAction.setToken(res.data))
            await this.getSelfData()
        }
        store.dispatch(statusAction.onEntry(false))
    }

    async register (data: RegisterType): Promise<boolean> {
        store.dispatch(statusAction.onEntry(true))
        const res = await this.PUT<ResMessage>({
            url: this.route.registe,
            body: data
        })
        store.dispatch(statusAction.onEntry(false))
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

    async deleteBanch (banchId: BanchType['Id']): Promise<ResType<null>> {
        store.dispatch(statusAction.onDeleteBanch(true))
        const res = await this.DELETE<null>({
            url: this.route.BanchAll,
            params: { banchId }
        })
        store.dispatch(statusAction.onDeleteBanch(false))
        return res
    }

    // 自己的個人資料
    async getSelfData (): Promise<void> {
        store.dispatch(statusAction.onFetchSelfData(true))
        const res = await this.GET<SelfDataType[]>({
            url: this.route.selfData,
            successShow: false
        })
        if (res.status) {
            store.dispatch(userAction.setSelfData(res.data[0]))
        }
        store.dispatch(statusAction.onFetchSelfData(false))
    }

    async UpdateSelfData (UserName: string): Promise<ResType<null>> {
        store.dispatch(statusAction.onUpdateSelfData(true))
        const res = await this.POST<null>({
            url: this.route.selfData,
            body: {
                UserName
            }
        })
        store.dispatch(statusAction.onUpdateSelfData(false))
        return res
    }

    // 部門樣式
    async getBanchStyle (banchId: number): Promise<ResType<BanchStyleType[]>> {
        store.dispatch(statusAction.onFetchBanchStyle(true))
        const res = await this.GET<BanchStyleType[]>({
            url: this.route.banchStyle,
            params: {
                banchId
            },
            successShow: false
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

    async deleteBanchStyle (StyleId: BanchStyleType['StyleId']): Promise<ResType<null>> {
        store.dispatch(statusAction.onUpdateBanchStyle(true))
        const res = await this.DELETE<null>({
            url: this.route.banchStyle,
            params: {
                StyleId
            }
        })
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
            },
            successShow: false
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

    async deleteBanchRule (RuleId: number): Promise<ResType<null>> {
        store.dispatch(statusAction.onDeleteBanchRule(true))
        const res = await this.DELETE<null>({
            url: this.route.banchRule,
            params: {
                RuleId
            }
        })
        store.dispatch(statusAction.onDeleteBanchRule(false))
        return res
    }

    // 公司資料
    async getCompanyInfo (
        props: {
            companyCode?: CompanyType['CompanyCode']
            companyId?: CompanyType['CompanyId']
        }): Promise<ResType<CompanyType[]>> {
        store.dispatch(statusAction.onFetchCompany(true))
        const res = await this.GET<CompanyType[]>({
            url: this.route.companyInfo,
            params: {
                ...props
            },
            successShow: false
        })
        if (res.status && res?.data?.length > 0) {
            store.dispatch(companyAction.setCompany(res.data[0]))
        } else {
            store.dispatch(companyAction.setCompany(null))
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
                SettlementDate: company.SettlementDate,
                TermStart: new Date().toISOString(),
                TermEnd: new Date().toISOString()
            }
        })
        store.dispatch(statusAction.onUpdateCompany(false))
        return res
    }

    // 員工 資料
    async getUserAll (v: any): Promise<ResType<UserType[]>> {
        store.dispatch(statusAction.onFetchUserAll(true))
        const res = await this.GET<UserType[]>({
            url: this.route.getUserAll,
            successShow: false,
            params: v
        })
        if (res.status) {
            store.dispatch(companyAction.setEmployee(res.data))
        }
        store.dispatch(statusAction.onFetchUserAll(false))
        // console.log(res)
        return res
    }

    async updateUser (user: UserType): Promise<ResType<null>> {
        store.dispatch(statusAction.onUpdateUser(true))
        const res = await this.POST<null>({
            url: this.route.userSingle,
            body: { ...user },
            params: {
                userId: user.UserId
            }
        })
        store.dispatch(statusAction.onUpdateUser(false))
        return res
    }

    async createUser (user: createUserType): Promise<ResType<null>> {
        store.dispatch(statusAction.onCreateUser(true))
        const res = await this.PUT<null>({
            url: this.route.userSingle,
            body: {
                ...user
            }
        })
        store.dispatch(statusAction.onCreateUser(false))
        return res
    }

    // 忘記密碼
    async forgetPassword ({
        Captcha, Password, PasswordConfirm, Account
    }): Promise<ResType<null>> {
        const res = await this.POST<null>({
            url: this.route.forgetPassword,
            body: {
                Captcha,
                NewPassword: Password,
                NewPasswordConfirm: PasswordConfirm,
                Email: Account
            }
        })
        return res
    }

    // 更換密碼
    async changePassword ({
        Captcha, Password, PasswordConfirm, OldPassword
    }): Promise<ResType<null>> {
        const res = await this.POST<null>({
            url: this.route.changePassword,
            body: {
                Captcha,
                NewPassword: Password,
                NewPasswordConfirm: PasswordConfirm,
                OldPassword
            }
        })
        return res
    }

    // 假日設定
    async getWeekendSetting (): Promise<ResType<WeekendSettingType[]>> {
        store.dispatch(statusAction.onFetchWeekendSetting(true))
        const res = await this.GET<WeekendSettingType[]>({
            url: this.route.weekendSetting,
            successShow: false
        })
        if (res.status) {
            store.dispatch(companyAction.setWeekendSetting(res.data))
        }
        store.dispatch(statusAction.onFetchWeekendSetting(false))
        return res
    }

    async createWeekendSetting (date: WeekendSettingType['Date']): Promise<ResType<null>> {
        store.dispatch(statusAction.onCreateWeekendSetting(true))
        const res = await this.PUT<null>({
            url: this.route.weekendSetting,
            body: {
                Date: date
            }
        })
        store.dispatch(statusAction.onCreateWeekendSetting(false))
        return res
    }

    async deleteWeekendSetting (weekendId: WeekendSettingType['WeekendId']): Promise<ResType<null>> {
        store.dispatch(statusAction.onDeleteWeekendSetting(true))
        const res = await this.DELETE<null>({
            url: this.route.weekendSetting,
            params: {
                weekendId
            }
        })
        store.dispatch(statusAction.onDeleteWeekendSetting(false))
        return res
    }

    // 等待公司回覆
    async getWaitReply (): Promise<ResType<WaitReplyType[]>> {
        store.dispatch(statusAction.onFetchWaitReply(true))
        const res = await this.GET<WaitReplyType[]>({
            url: this.route.waitReply,
            successShow: false
        })
        if (res.status) {
            store.dispatch(companyAction.setWaitReply(res.data))
        }
        store.dispatch(statusAction.onFetchWaitReply(false))
        return res
    }

    async updateWaitReply ({
        SpecifyTag, IsAccept, WaitId
    }: {
        SpecifyTag: WaitReplyType['SpecifyTag']
        IsAccept: WaitReplyType['IsAccept']
        WaitId: WaitReplyType['WaitId']
    }): Promise<ResType<null>> {
        store.dispatch(statusAction.onUpdateWaitReply(true))
        const res = await this.POST<null>({
            url: this.route.waitReply,
            body: {
                SpecifyTag,
                IsAccept,
                WaitId
            }
        })
        store.dispatch(statusAction.onUpdateWaitReply(false))
        return res
    }

    async createWaitReply (companyCode: CompanyType['CompanyCode']): Promise<ResType<null>> {
        store.dispatch(statusAction.onCreateWaitReply(true))
        const res = await this.PUT<null>({
            url: this.route.waitReply,
            body: {
                companyCode
            }
        })
        store.dispatch(statusAction.onCreateWaitReply(false))
        return res
    }

    async getWorkTime (
        year: workTimeType['Year'],
        month: workTimeType['Month'],
        userId: workTimeType['UserId']
    ): Promise<ResType<workTimeType[]>> {
        store.dispatch(statusAction.onFetchWorkTime(true))
        const res = await this.GET<workTimeType[]>({
            url: this.route.workTime,
            params: {
                ...(year && { year }),
                ...(month && { month }),
                ...(userId && { userId })
            },
            successShow: false
        })
        if (res.status) {
            store.dispatch(companyAction.setWorkTime(res.data))
        }
        store.dispatch(statusAction.onFetchWorkTime(false))
        return res
    }

    // 工作時數
    async createWorkTime (data: workTimeType): Promise<ResType<null>> {
        const res = await this.PUT<null>({
            url: this.route.workTime,
            body: data
        })
        return res
    }

    // 績效
    async getPerformance (v: any): Promise<ResType<performanceType[]>> {
        // console.log(v)
        store.dispatch(statusAction.onFetchPerformance(true))
        const res = await this.GET<performanceType[]>({
            url: this.route.performance,
            successShow: false,
            params: v
        })
        if (res.status) {
            store.dispatch(companyAction.setPerformance(res.data))
        }
        store.dispatch(statusAction.onFetchPerformance(false))
        return res
    }

    async updatePerformance (v: performanceType): Promise<ResType<null>> {
        store.dispatch(statusAction.onUpdatePerformance(true))
        const res = await this.POST<null>({
            url: this.route.performance,
            body: v
        })
        store.dispatch(statusAction.onUpdatePerformance(false))
        return res
    }

    async deletePerformance (performanceId: performanceType['PerformanceId']): Promise<ResType<null>> {
        store.dispatch(statusAction.onDeletePerformance(true))
        const res = await this.DELETE<null>({
            url: this.route.performance,
            params: {
                performanceId
            }
        })
        store.dispatch(statusAction.onDeletePerformance(false))
        return res
    }

    async createPerformance (v: performanceType): Promise<ResType<null>> {
        store.dispatch(statusAction.onCreatePerformance(true))
        const res = await this.PUT<null>({
            url: this.route.performance,
            body: v
        })
        store.dispatch(statusAction.onCreatePerformance(false))
        return res
    }

    // 績效複製
    async copyPerformance (v: performanceType['PerformanceId']): Promise<ResType<null>> {
        store.dispatch(statusAction.onCopyPerformance(true))
        const res = await this.PUT<null>({
            url: this.route.performanceCopy,
            body: {
                PerformanceId: v,
                IsResetGrade: true, // 重設 分數
                IsResetDirections: true // 重設績效描述
            }
        })
        store.dispatch(statusAction.onCopyPerformance(false))
        return res
    }

    // 年度績效分數
    async getYearPerformance (v: yearPerformanceParamsType): Promise<ResType<yearPerformanceType[]>> {
        // console.log(v)
        store.dispatch(statusAction.onFetchYearPerformance(true))
        const res = await this.GET<yearPerformanceType[]>({
            url: this.route.yearPerformance,
            successShow: false,
            params: v
        })
        if (res.status) {
            store.dispatch(companyAction.setYearPerformance(res.data))
        }
        store.dispatch(statusAction.onFetchYearPerformance(false))
        return res
    }

    async googleLogin (): Promise<ResType<string>> {
        const res = await this.GET<string>({
            url: this.route.googleLogin
        })
        return res
    }

    // 班表查詢
    async getShiftMonth (v: shiftMonthParamsType): Promise<ResType<shiftMonthType[]>> {
        // console.log(v)
        store.dispatch(statusAction.onFetchShiftMonth(true))
        const res = await this.GET<shiftMonthType[]>({
            url: this.route.shiftMonth,
            successShow: false,
            params: v
        })
        if (res.status) {
            store.dispatch(shiftEditAction.setShiftEdit(res as any))
        }
        store.dispatch(statusAction.onFetchShiftMonth(false))
        return res
    }

    // 班表統計
    async getShiftTotal (v: shiftMonthParamsType): Promise<ResType<shiftTotalType[]>> {
        // console.log(v)
        store.dispatch(statusAction.onFetchShiftTotal(true))
        const res = await this.GET<shiftTotalType[]>({
            url: this.route.shiftTotal,
            successShow: false,
            params: v
        })
        if (res.status) {
            store.dispatch(shiftEditAction.setShiftTotal(res.data))
        }
        store.dispatch(statusAction.onFetchShiftTotal(false))
        return res
    }

    // 班表歷程
    async getShiftHistory (v: shiftMonthParamsType): Promise<ResType<shiftHistoryType[]>> {
        store.dispatch(statusAction.onFetchShiftHistory(true))
        const res = await this.GET<shiftHistoryType[]>({
            url: this.route.shiftHistory,
            successShow: false,
            params: v
        })
        if (res.status) {
            store.dispatch(shiftEditAction.setShiftHistory(res.data))
        }
        store.dispatch(statusAction.onFetchShiftHistory(false))
        return res
    }
}
export default new ApiControl()
