import { Form, FormInstance, Input, InputNumber } from 'antd'
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
                className='row'
                form={form}
                onFinish={() => {
                    modalInfo.onSave(form)
                }}
            >
                <Form.Item
                    name="BanchName"
                    className='col-md-6'
                    label="部門名稱"
                    rules={[{ required: true }]}
                    initialValue={modalInfo?.data?.BanchName || ''}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="Sort"
                    className='col-md-6'
                    label="排序"
                    initialValue={modalInfo?.data?.Sort || 0}
                >
                    <InputNumber />
                </Form.Item>
                <Modal.Footer>
                    {
                        () => (
                            <>
                                <Btn.Cancel onClick={() => { void modalInfo.onClose() }} />
                                <Btn.Save />
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
