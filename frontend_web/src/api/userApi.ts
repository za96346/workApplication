import { userReducerType } from 'reducer/reducer/userReducer'
import apiAbstract from './apiAbstract'
import userTypes from 'types/user'

declare namespace params {
    interface deleted {
        UserId: number
    }
    type add = Omit<
    userTypes.TABLE,
    'CreateTime' | 'LastModify' | 'DeleteTime' | 'DeleteFlag' | 'CompanyId'
    >
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

    async delete (v: params.deleted): Promise<void> {
        return await this.DELETE<null>({
            url: this.route,
            data: v
        })
    }

    async update (v: params.add): Promise<void> {
        return await this.POST<null>({
            url: this.route,
            data: v
        })
    }

    async add (v: params.add): Promise<void> {
        return await this.PUT<null>({
            url: this.route,
            data: v
        })
    }
}
export default new userApi()
export {
    type params as userApiParams
}
