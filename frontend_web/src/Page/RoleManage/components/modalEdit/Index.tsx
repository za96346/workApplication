import Session from '../../methods/session'
import { Form, FormInstance, Input, InputNumber } from 'antd'
import React, { useEffect } from 'react'
import Btn from 'shared/Button'
import { Modal } from 'shared/Modal/Index'
import modal from 'shared/Modal/types'
import { modalTitle, modalType } from 'static'
import roleTypes from 'types/role'
import CheckTree from 'Page/RoleManage/components/CheckTree/Index'
import { useAppDispatch } from 'hook/redux'

interface modalInfo {
    type?: modalType
    onSave: (v: FormInstance<any>) => void
    data?: roleTypes.TABLE
}

interface props {
    modalInfo: modal.modalInfoProps<modalInfo>
}

const ModalEdit = ({ modalInfo }: props): JSX.Element => {
    const { dispatch, action } = useAppDispatch()
    // const { type } = modalInfo
    const [form] = Form.useForm()
    console.log('Session => ', Session.Instance.get())

    useEffect(() => () => {
        Session.Instance.set({})
        dispatch(action.role.clearSingle())
    }, [])
    return (
        <Session.Provider>
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
                    name="RoleName"
                    label="角色名稱"
                    className='col-md-6'
                    rules={[{ required: true }]}
                    initialValue={modalInfo?.data?.RoleName || ''}
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
                <CheckTree />
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
        </Session.Provider>
    )
}
export default Modal<modalInfo, any>({
    children: ModalEdit,
    title: (v) => `${modalTitle[v?.type]}角色`,
    width: () => '100vw'
})
