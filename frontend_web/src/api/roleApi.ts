import { roleReducerType } from 'reducer/reducer/roleReducer'
import apiAbstract from './apiAbstract'
import systemTypes from 'types/system'

declare namespace params {
    interface getSingle {
        RoleId: number
    }

    interface update {
        RoleId?: number
        RoleName: string
        StopFlag?: Flag
        Data: Partial<systemTypes.auth['permission']>
    }

    interface deleted {
        RoleId: number
    }
}

class roleApi extends apiAbstract {
    private readonly route = 'workApp/role/'

    async get (): Promise<void> {
        const [result] = this.makeFormData([{
            formName: 'roleManage',
            validCheck: true
        }])
        return await this.GET<null>({
            url: this.route,
            data: result
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
            data: v,
            checkTitle: this.checkTitle.confirmUpdate
        })
    }

    async add (v: params.update): Promise<void> {
        return await this.PUT<null>({
            url: this.route,
            data: v,
            checkTitle: this.checkTitle.confirmAdd
        })
    }

    async delete (v: params.deleted): Promise<void> {
        return await this.DELETE<null>({
            url: this.route,
            data: v,
            checkTitle: this.checkTitle.confirmDelete
        })
    }
}
export default new roleApi()
export {
    type params as roleApiParams
}
