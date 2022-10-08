import api from './apiAbs'
import axios from 'axios'
import { LoginResponse, LoginType } from '../type'
import checkStatus from './checkStatus'
import userAction from '../reduxer/action/userAction'
import { store } from '../reduxer/store'

class ApiControl extends api {
    baseUrl: string
    constructor () {
        super()
        this.baseUrl = process.env.REACT_APP_API
        this.route = {
            login: 'workApp/entry/login'
        }
    }

    protected async POST<T> (url: any, body: any, callBack: any): Promise<{ data: T, status: boolean }> {
        try {
            const res = await axios.post(`${this.baseUrl}/${url}`, body)
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

    async login (data: LoginType): Promise<any> {
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
}
export default new ApiControl()
