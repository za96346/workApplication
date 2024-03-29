import apiAbstract from './apiAbstract'

declare namespace params {
    interface login {
        Account: string
        Password: string
    }
}

class entryApi extends apiAbstract {
    private readonly route = 'entry/login'

    async login (params: params.login): Promise<void> {
        return await this.POST<null>({
            url: this.route,
            data: params,
            checkTitle: '是否登入'
        })
            .then((v) => { window.location.href = '/' })
    }
}
export default new entryApi()
export {
    type params as entryApiParams
}
