import { modalType } from 'static'
import performanceTypes from 'types/performance'
import ModalPerformance from '../components/ModalPeroformance/Index'
import YearPerformanceSession from './yearPerformanceSession'
import api from 'api/Index'

const BtnEvent = ({ type, value }: BtnEventParams<performanceTypes.year>): void => {
    if (type === modalType.inquire) {
        ModalPerformance.open({
            userId: value?.UserId,
            year: value?.Year,
            userName: value?.UserName,
            onDestroy: () => {
                void api.performance.getYear(
                    YearPerformanceSession.Instance.get().currentParams
                )
            }
        })
    }
}
export default BtnEvent
