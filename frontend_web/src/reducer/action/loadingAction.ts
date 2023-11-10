import { type action, loadingType } from '../types'

class statusAcion {
    onLoading (s: boolean): action {
        return {
            type: loadingType.SET_LOADING,
            payload: {
                loading: s
            }
        }
    }
}
export default new statusAcion()
