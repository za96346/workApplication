import { Form, FormInstance, Input } from 'antd'
import React from 'react'
import Btn from 'shared/Button'
import { Modal } from 'shared/Modal/Index'
import modal from 'shared/Modal/types'
import { modalType } from 'static'

interface modalInfo {
    type?: typeof modalType
    onSave: (v: FormInstance<any>) => void
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
                <Form.Item name="BanchName" label="部門名稱" rules={[{ required: true }]}>
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
    title: () => '編輯部門',
    width: (isLess) => isLess('md') ? '100vw' : '500px'
})
