import api from 'api/Index'
import ModalEdit from '../components/modalEdit/Index'
import { modalType } from 'static'
import Session from './session'
import roleTypes from 'types/role'

const BtnEvent = ({ type, value }: BtnEventParams<roleTypes.TABLE>): void => {
    const onClose = (): void => {
        ModalEdit.close({})
        void api.role.get()
    }
    const onOpen = (): void => {
        ModalEdit.open({
            type,
            data: value,
            onSave: (form) => {
                const permission = Session.Instance.get()
                const fields = form.getFieldsValue()
                void form.validateFields()
                    .then(() => {
                        if (type === modalType.edit) {
                            void api.role.update({
                                Data: { ...permission },
                                RoleId: value.RoleId,
                                RoleName: fields.RoleName,
                                StopFlag: 'N'
                            }).then(onClose)
                        } else if (type === modalType.add) {
                            void api.role.add({
                                Data: { ...permission },
                                RoleName: fields.RoleName,
                                StopFlag: 'N'
                            }).then(onClose)
                        }
                    })
            }
        })
    }
    if (type === modalType.delete) {
        console.log()
    } else if (type === modalType.add) {
        onOpen()
    } else if (type === modalType.edit) {
        void api.role.getSingle({
            RoleId: value?.RoleId
        }).then(onOpen)
    }
}
export default BtnEvent
