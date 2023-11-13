import systemAction from 'reducer/action/systemAction'
import apiAbstract from './apiAbstract'
import systemTypes from 'types/system'

class systemApi extends apiAbstract {
    private readonly route = 'workApp/system/auth'

    async auth (): Promise<systemTypes.auth> {
        return await this.GET<systemTypes.auth>({
            url: this.route
        }).then((res) => {
            this.store.dispatch(systemAction.setAuth(res))
            return res
        })
    }
}
export default new systemApi()
