import api from 'api/Index'
import ModalEdit from '../components/modalEdit/Index'

const BtnEvent = ({ type, reload, value }): void => {
    ModalEdit.open({
        type,
        onSave: (form) => {
            void form.validateFields()
                .then(() => {
                    void api.companyBanch.add(form.getFieldsValue())
                        .then(() => {
                            ModalEdit.open({})
                        })
                })
        }
    })
}
export default BtnEvent
