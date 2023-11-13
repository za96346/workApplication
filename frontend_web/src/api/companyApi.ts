import { companyReducerType } from 'reducer/reducer/companyReducer'
import apiAbstract from './apiAbstract'

declare namespace params {
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
}
export default new companyApi()
export {
    type params as companyApiParams
}
