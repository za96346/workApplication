import { companyReducerType } from 'reducer/reducer/companyReducer'
import apiAbstract from './apiAbstract'

declare namespace params {
    interface update {
        CompanyCode: string
        CompanyName: string
        CompanyLocation: string
        CompanyPhoneNumber: string
        BossId: number
    }
}

class companyApi extends apiAbstract {
    private readonly route = 'workApp/company/'

    async getMine (): Promise<companyReducerType['mine']> {
        return await this.GET<companyReducerType['mine']>({
            url: this.route
        }).then((res) => {
            this.store.dispatch(this.action.company.setMine(res))
            return res
        })
    }

    async update (v: params.update): Promise<void> {
        return await this.POST<null>({
            url: this.route,
            data: v,
            check_title: this.checkTitle.confirmUpdate
        })
    }
}
export default new companyApi()
export {
    type params as companyApiParams
}
