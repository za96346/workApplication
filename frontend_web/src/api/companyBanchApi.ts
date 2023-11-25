import companyBanchTypes from 'types/companyBanch'
import apiAbstract from './apiAbstract'

declare namespace params {
    interface add {
        BanchName: string
    }

    interface deleted {
        BanchId: number
    }

    interface update {
        BanchName: string
        BanchId: number
    }

    interface get {
        BanchName: string
    }
}

class companyBanchApi extends apiAbstract {
    private readonly route = 'workApp/banch/'

    async get (v?: params.get): Promise<companyBanchTypes.reducerType['all']> {
        return await this.GET<companyBanchTypes.reducerType['all']>({
            url: this.route,
            data: v
        }).then((res) => {
            this.store.dispatch(this.action.companyBanch.setAll(res))
            return res
        })
    }

    async add (v: params.add): Promise<null> {
        return await this.PUT<null>({
            url: this.route,
            data: v,
            checkTitle: this.checkTitle.confirmAdd
        })
    }

    async delete (v: params.deleted): Promise<null> {
        return await this.DELETE<null>({
            url: this.route,
            data: v,
            checkTitle: this.checkTitle.confirmDelete
        })
    }

    async update (v: params.update): Promise<null> {
        return await this.POST<null>({
            url: this.route,
            data: v,
            checkTitle: this.checkTitle.confirmUpdate
        })
    }

    async getSelector (): Promise<companyBanchTypes.reducerType['selector']> {
        return await this.GET<companyBanchTypes.reducerType['selector']>({
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
