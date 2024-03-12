import api from 'api/Index'
import ModalEdit from '../components/modalEdit/Index'
import { modalType } from 'static'
import performanceTypes from 'types/performance'
import ModalChangeBanch from '../components/modalChangeBanch/Index'
import PerformanceSession from './performanceSession'

const BtnEvent = ({ type, value }: BtnEventParams<performanceTypes.TABLE>): void => {
    const onClose = (): void => {
        ModalEdit.close({})
        ModalChangeBanch.close({})
        void api.performance.get(PerformanceSession.Instance.get().currentParams)
    }
    if (type === modalType.delete) {
        void api.performance.delete({ performanceId: value?.PerformanceId })
            .then(onClose)
        return
    }
    if (type === modalType.changeBanch) {
        ModalChangeBanch.open({
            onSave: (form) => {
                void api.performance.changeBanch({
                    PerformanceId: value?.PerformanceId,
                    ...form?.getFieldsValue()
                })
                    .then(onClose)
            }
        })
        return
    }
    ModalEdit.open({
        type,
        data: value,
        onSave: (form) => {
            const fields = form.getFieldsValue()
            const year = fields?.Year?.$y - 1911
            if (type === modalType.edit) {
                void api.performance.update({
                    ...fields,
                    Year: year,
                    PerformanceId: value?.PerformanceId
                })
                    .then(onClose)
            } else if (type === modalType.add) {
                void api.performance.add({
                    ...fields,
                    Year: year
                })
                    .then(onClose)
            } else if (type === modalType.copy) {
                void api.performance.copy({
                    ...fields,
                    Year: year
                })
                    .then(onClose)
            }
        }
    })
}
export default BtnEvent
