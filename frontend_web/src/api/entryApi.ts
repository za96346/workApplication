import apiAbstract from './apiAbstract'

declare namespace params {
    interface login {
        Account: string
        Password: string
    }
}

class entryApi extends apiAbstract {
    private readonly route = 'workApp/entry/login'

    async login (params: params.login): Promise<void> {
        return await this.POST<null>({
            url: this.route,
            data: params,
            check_text: '是否登入'
        })
            .then((v) => { window.location.href = '/' })
    }
}
export default new entryApi()
export {
    params as entryApiParams
}
