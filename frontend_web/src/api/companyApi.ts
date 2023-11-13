import apiAbstract from './apiAbstract'

declare namespace params {
}

class companyApi extends apiAbstract {
    private readonly route = 'workApp/company/'

    async getMine (): Promise<void> {
        return await this.GET<null>({
            url: this.route
        }).then((res) => {
            this.store.dispatch(this.action.company.setMine(res))
        })
    }
}
export default new companyApi()
export {
    params as companyApiParams
}
