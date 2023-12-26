import { Form, FormInstance, Input } from 'antd'
import React from 'react'
import Btn from 'shared/Button'
import { Modal } from 'shared/Modal/Index'
import modal from 'shared/Modal/types'
import { modalTitle, modalType } from 'static'
import companyBanchTypes from 'types/companyBanch'

interface modalInfo {
    type?: modalType
    onSave: (v: FormInstance<any>) => void
    data?: companyBanchTypes.TABLE
}

interface props {
    modalInfo: modal.modalInfoProps<modalInfo>
}

const ModalEdit = ({ modalInfo }: props): JSX.Element => {
    // const { type } = modalInfo
    const [form] = Form.useForm()
    return (
        <>
            <Form
                name="validateOnly"
                autoComplete="off"
                form={form}
            >
                <Form.Item
                    name="BanchName"
                    label="部門名稱"
                    rules={[{ required: true }]}
                    initialValue={modalInfo?.data?.BanchName || ''}
                >
                    <Input />
                </Form.Item>
                <Modal.Footer>
                    {
                        () => (
                            <>
                                <Btn.Cancel onClick={() => { void modalInfo.onClose() }} />
                                <Btn.Save
                                    onClick={() => {
                                        modalInfo.onSave(form)
                                    }}
                                />
                            </>
                        )
                    }
                </Modal.Footer>
            </Form>

        </>
    )
}
export default Modal<modalInfo, any>({
    children: ModalEdit,
    title: (v) => `${modalTitle[v?.type]}部門`,
    width: (isLess) => isLess('md') ? '100vw' : '500px'
})
