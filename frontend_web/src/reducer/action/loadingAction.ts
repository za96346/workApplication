import { type action, loadingType } from '../types'

class loadingAction {
    onLoading (s: Record<string, boolean>): action {
        return {
            type: loadingType.SET_LOADING,
            payload: s
        }
    }
}
export default new loadingAction()
