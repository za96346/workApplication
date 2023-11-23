import { type action, performanceType } from '../types'
import { performanceReducerType } from 'reducer/reducer/performanceReducer'

class performanceAction {
    setAll (v: performanceReducerType['all']): action {
        return {
            type: performanceType.SET_PERFORMANCE_ALL,
            payload: {
                all: v
            }
        }
    }
}
export default new performanceAction()
