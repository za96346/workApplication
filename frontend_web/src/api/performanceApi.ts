import apiAbstract from './apiAbstract'
import performanceTypes from 'types/performance'

declare namespace params {
    interface get {
        UserName: string
        RoleId: number
        BanchId: number
    }
    interface deleted {
        performanceId: performanceTypes.TABLE['PerformanceId']
    }
    type add = Omit<
    performanceTypes.TABLE,
    'LastModify'
    | 'CreateTime'
    | 'DeleteTime'
    | 'DeleteFlag'
    | 'CompanyId'
    >
}

class performanceApi extends apiAbstract {
    private readonly route = 'performance/'

    async get (v?: params.get): Promise<void> {
        return await this.GET<performanceTypes.reducerType['all']>({
            url: this.route,
            data: v
        })
            .then((v) => {
                this.store.dispatch(this.action.performance.setAll(v))
            })
    }

    async getYear (v?: params.get): Promise<void> {
        return await this.GET<performanceTypes.reducerType['year']>({
            url: this.route + 'year',
            data: v
        })
            .then((v) => {
                this.store.dispatch(this.action.performance.setYear(v))
            })
    }

    async add (v: params.add): Promise<void> {
        return await this.PUT<null>({
            url: this.route,
            data: v,
            checkTitle: this.checkTitle.confirmAdd
        })
    }

    async copy (v: params.add): Promise<void> {
        return await this.PUT<null>({
            url: this.route + 'copy',
            data: v,
            checkTitle: this.checkTitle.confirmAdd
        })
    }

    async update (v: params.add): Promise<void> {
        return await this.POST<null>({
            url: this.route,
            data: v,
            checkTitle: this.checkTitle.confirmUpdate
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
export default new performanceApi()
export {
    type params as performanceParams
}
