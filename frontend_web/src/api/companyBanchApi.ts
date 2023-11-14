import apiAbstract from './apiAbstract'
import { companyBanchReducerType } from 'reducer/reducer/companyBanchReducer'

declare namespace params {
    interface add {
        BanchName: string
    }

    interface deleted {
        BanchId: number
    }

    interface edit {
        BanchName: string
        BanchId: number
    }
}

class companyBanchApi extends apiAbstract {
    private readonly route = 'workApp/banch/'

    async get (): Promise<companyBanchReducerType['all']> {
        return await this.GET<companyBanchReducerType['all']>({
            url: this.route
        }).then((res) => {
            this.store.dispatch(this.action.companyBanch.setAll(res))
            return res
        })
    }

    async add (v: params.add): Promise<null> {
        return await this.PUT<null>({
            url: this.route,
            data: v,
            check_text: '確認新增？'
        })
    }

    async delete (v: params.deleted): Promise<null> {
        return await this.DELETE<null>({
            url: this.route,
            data: v,
            check_text: '確認刪除？'
        })
    }

    async edit (v: params.edit): Promise<null> {
        return await this.POST<null>({
            url: this.route,
            data: v,
            check_text: '確認儲存？'
        })
    }

    async getSelector (): Promise<companyBanchReducerType['selector']> {
        return await this.GET<companyBanchReducerType['selector']>({
            url: this.route + 'selector'
        }).then((res) => {
            this.store.dispatch(this.action.companyBanch.setSelector(res))
            return res
        })
    }
}
export default new companyBanchApi()
export {
    type params as companyBanchApiParams
}
