import apiAbstract from './apiAbstract'

class authApi extends apiAbstract {
    private readonly route: string
    constructor () {
        super()
        this.route = '/workapp/auth'
    }

    async get (): Promise<void> {
        this.GET({
            url: this.route
        })
    }
}
export default new authApi()
