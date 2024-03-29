import userApi, { userApiParams } from './userApi'
import entryApi, { entryApiParams } from './entryApi'
import systemApi, { systemApiParams } from './systemApi'
import companyApi, { companyApiParams } from './companyApi'
import companyBanchApi, { companyBanchApiParams } from './companyBanchApi'
import roleApi, { roleApiParams } from './roleApi'
import performanceApi, { performanceParams } from './performanceApi'

const api = {
    entry: entryApi,
    system: systemApi,
    user: userApi,
    company: companyApi,
    companyBanch: companyBanchApi,
    role: roleApi,
    performance: performanceApi
}
export default api
export {
    type userApiParams,
    type entryApiParams,
    type companyApiParams,
    type companyBanchApiParams,
    type systemApiParams,
    type roleApiParams,
    type performanceParams
}
