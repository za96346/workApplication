import { userReducerType } from 'reducer/reducer/userReducer'
import apiAbstract from './apiAbstract'

declare namespace params {
}

class userApi extends apiAbstract {
    private readonly route = 'workApp/user/'

    async getMine (): Promise<userReducerType['mine']> {
        return await this.GET<userReducerType['mine']>({
            url: this.route + 'my'
        }).then((res) => {
            this.store.dispatch(this.action.user.setMine(res))
            return res
        })
    }

    async getEmployee (): Promise<userReducerType['employee']> {
        return await this.GET<userReducerType['employee']>({
            url: this.route
        }).then((res) => {
            this.store.dispatch(this.action.user.setEmployee(res))
            return res
        })
    }
}
export default new userApi()
export {
    type params as userApiParams
}
