import userApi from './userApi'
import entryApi from './entryApi'
import systemApi from './systemApi'
import companyApi from './companyApi'

const api = {
    entry: entryApi,
    system: systemApi,
    user: userApi,
    company: companyApi
}
export default api
