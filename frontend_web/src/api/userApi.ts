import apiAbstract from './apiAbstract'
import userTypes from 'types/user'

declare namespace params {
    interface deleted {
        UserId: number
    }

    interface getEmployee {
        BanchId?: number
        RoleId?: number
        UserName?: string
        EmployeeNumber?: string
    }

    interface updatePassword {
        OldPassword: string
        NewPassword: string
        NewPasswordAgain: string
        UserId: number
    }

    type add = Omit<
    userTypes.TABLE,
    'CreateTime' | 'LastModify' | 'DeleteTime' | 'DeleteFlag' | 'CompanyId'
    >

    interface updateMy {
        UserName: string
    }
}

class userApi extends apiAbstract {
    private readonly route = 'workApp/user/'

    async getMine (): Promise<userTypes.reducerType['mine']> {
        return await this.GET<userTypes.reducerType['mine']>({
            url: this.route + 'my'
        }).then((res) => {
            this.store.dispatch(this.action.user.setMine(res))
            return res
        })
    }

    async getEmployee (v: params.getEmployee): Promise<userTypes.reducerType['employee']> {
        return await this.GET<userTypes.reducerType['employee']>({
            url: this.route,
            data: v
        }).then((res) => {
            this.store.dispatch(this.action.user.setEmployee(res))
            return res
        })
    }

    async delete (v: params.deleted): Promise<void> {
        return await this.DELETE<null>({
            url: this.route,
            data: v,
            checkTitle: this.checkTitle.confirmDelete
        })
    }

    async update (v: params.add): Promise<void> {
        return await this.POST<null>({
            url: this.route,
            data: v,
            checkTitle: this.checkTitle.confirmUpdate
        })
    }

    async updateMine (v: params.updateMy): Promise<void> {
        return await this.POST<null>({
            url: this.route + 'my',
            data: v,
            checkTitle: this.checkTitle.confirmUpdate
        })
    }

    async add (v: params.add): Promise<void> {
        return await this.PUT<null>({
            url: this.route,
            data: v,
            checkTitle: this.checkTitle.confirmAdd
        })
    }

    async getSelector (v: params.getEmployee): Promise<userTypes.reducerType['selector']> {
        return await this.GET<userTypes.reducerType['selector']>({
            url: this.route + 'selector',
            data: v
        }).then((res) => {
            this.store.dispatch(this.action.user.setSelector(res))
            return res
        })
    }

    async updatePassword (v: params.updatePassword): Promise<void> {
        return await this.POST<null>({
            url: this.route + 'password',
            data: v
        })
    }
}
export default new userApi()
export {
    type params as userApiParams
}
