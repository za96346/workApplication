import api from 'api/Index'
import ModalEdit from '../components/modalEdit/Index'
import { modalType } from 'static'
import userTypes from 'types/user'

const BtnEvent = ({ type, value }: BtnEventParams<userTypes.TABLE>): void => {
    const onClose = (): void => {
        ModalEdit.close({})
        void api.user.getEmployee({})
    }
    if (type === modalType.delete) {
        void api.user.delete({ UserId: value?.UserId })
            .then(onClose)
        return
    }

    ModalEdit.open({
        type,
        data: value,
        onSave: (form) => {
            const fields = form.getFieldsValue()
            if (type === modalType.edit) {
                void api.user.update({
                    ...fields,
                    UserId: value?.UserId
                })
                    .then(onClose)
            } else if (type === modalType.add) {
                void api.user.add(fields)
                    .then(onClose)
            }
        }
    })
}
export default BtnEvent
