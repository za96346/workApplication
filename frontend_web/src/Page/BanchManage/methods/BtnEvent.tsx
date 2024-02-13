import api from 'api/Index'
import ModalEdit from '../components/modalEdit/Index'
import { modalType } from 'static'
import companyBanchTypes from 'types/companyBanch'

const BtnEvent = ({ type, value }: BtnEventParams<companyBanchTypes.TABLE>): void => {
    const onClose = (): void => {
        ModalEdit.close({})
        void api.companyBanch.get()
    }
    if (type === modalType.delete) {
        void api.companyBanch.delete({ BanchId: value?.BanchId })
            .then(onClose)
        return
    }
    ModalEdit.open({
        type,
        data: value,
        onSave: (form) => {
            const fields = form.getFieldsValue()
            if (type === modalType.edit) {
                void api.companyBanch.update({
                    ...fields,
                    BanchId: value?.BanchId
                })
                    .then(onClose)
            } else if (type === modalType.add) {
                void api.companyBanch.add(fields)
                    .then(onClose)
            }
        }
    })
}
export default BtnEvent
