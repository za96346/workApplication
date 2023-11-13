import apiAbstract from './apiAbstract'

declare namespace params {
}

class userApi extends apiAbstract {
    private readonly route = 'workApp/user/'

    async getMine (): Promise<void> {
        return await this.GET<null>({
            url: this.route + 'my'
        }).then((res) => {
            this.store.dispatch(this.action.user.setMine(res))
        })
    }

    async get (): Promise<void> {
        return await this.GET<null>({
            url: this.route
        }).then((res) => {
            this.store.dispatch(this.action.user.setEmployee(res))
        })
    }
}
export default new userApi()
export {
    params as userApiParams
}
