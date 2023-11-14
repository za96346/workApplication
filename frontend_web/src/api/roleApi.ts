import { roleReducerType } from 'reducer/reducer/roleReducer'
import apiAbstract from './apiAbstract'
import systemTypes from 'types/system'

declare namespace params {
    interface getSingle {
        RoleId: number
    }

    interface update {
        RoleId: number
        RoleName: string
        StopFlag: Flag
        Type: string
        Data: systemTypes.auth['permission']
    }
}

class roleApi extends apiAbstract {
    private readonly route = 'workApp/role/'

    async get (): Promise<void> {
        return await this.GET<null>({
            url: this.route,
            data: {}
        })
            .then((v) => {
                this.store.dispatch(this.action.role.setAll(v))
            })
    }

    async getSingle (v: params.getSingle): Promise<void> {
        return await this.GET<roleReducerType['single']>({
            url: this.route + 'single',
            data: v
        })
            .then((v) => {
                this.store.dispatch(this.action.role.setSingle(v))
            })
    }

    async getSelector (): Promise<void> {
        return await this.GET<null>({
            url: this.route + 'selector',
            data: {}
        })
            .then((v) => {
                this.store.dispatch(this.action.role.setSelector(v))
            })
    }

    async update (v: params.update): Promise<void> {
        return await this.POST<null>({
            url: this.route,
            data: v
        })
    }
}
export default new roleApi()
export {
    type params as roleApiParams
}
