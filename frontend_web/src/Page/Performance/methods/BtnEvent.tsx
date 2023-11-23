import api from 'api/Index'
import ModalEdit from '../components/modalEdit/Index'
import { modalType } from 'static'
import performanceTypes from 'types/performance'

const BtnEvent = ({ type, value }: BtnEventParams<performanceTypes.TABLE>): void => {
    const onClose = (): void => {
        ModalEdit.close({})
        void api.performance.get()
    }
    if (type === modalType.delete) {
        void api.performance.delete({ performanceId: value?.PerformanceId })
            .then(onClose)
        return
    }
    ModalEdit.open({
        type,
        data: value,
        onSave: (form) => {
            const fields = form.getFieldsValue()
            void form.validateFields()
                .then(() => {
                    console.log(fields)
                    if (type === modalType.edit) {
                        void api.performance.update({
                            ...fields,
                            Year: fields?.Year?.$y - 1911,
                            PerformanceId: value?.PerformanceId
                        })
                            .then(onClose)
                    } else if (type === modalType.add) {
                        void api.performance.add({
                            ...fields,
                            Year: fields?.Year?.$y - 1911
                        })
                            .then(onClose)
                    }
                })
        }
    })
}
export default BtnEvent
