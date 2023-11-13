import apiAbstract from './apiAbstract'
import { companyBanchReducerType } from 'reducer/reducer/companyBanchReducer'

declare namespace params {
    interface update {
        BanchName: string

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

    async add (v: params.update): Promise<null> {
        return await this.PUT<null>({
            url: this.route,
            data: v,
            check_text: '確認新增？'
        })
    }
}
export default new companyBanchApi()
export {
    type params as companyBanchApiParams
}
