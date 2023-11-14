import systemAction from 'reducer/action/systemAction'
import apiAbstract from './apiAbstract'
import systemTypes from 'types/system'

declare namespace params {
}

class systemApi extends apiAbstract {
    private readonly route = 'workApp/system/'

    async auth (): Promise<systemTypes.auth> {
        return await this.GET<systemTypes.auth>({
            url: this.route + 'auth'
        }).then((res) => {
            this.store.dispatch(systemAction.setAuth(res))
            return res
        })
    }

    async func (): Promise<systemTypes.func> {
        return await this.GET<systemTypes.func>({
            url: this.route + 'func'
        }).then((res) => {
            this.store.dispatch(systemAction.setFunc(res))
            return res
        })
    }
}
export default new systemApi()
export {
    type params as systemApiParams
}
