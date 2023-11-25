import performanceTypes from 'types/performance'
import { type action, performanceType } from '../types'

class performanceAction {
    setAll (v: performanceTypes.reducerType['all']): action {
        return {
            type: performanceType.SET_PERFORMANCE_ALL,
            payload: {
                all: v
            }
        }
    }

    setYear (v: performanceTypes.reducerType['year']): action {
        return {
            type: performanceType.SET_PERFORMANCE_YEAR,
            payload: {
                year: v
            }
        }
    }
}
export default new performanceAction()
