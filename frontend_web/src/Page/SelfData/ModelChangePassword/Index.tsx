import { Form, Input } from 'antd'
import React from 'react'
import Btn from 'shared/Button'
import { Modal } from 'shared/Modal/Index'
import api from 'api/Index'

const Index = ({ modalInfo }: any): JSX.Element => {
    return (
        <>
            <Form
                onFinish={(v) => {
                    void api.user.updatePassword({
                        ...v,
                        UserId: modalInfo?.UserId
                    }).then(() => {
                        modalInfo?.onClose()
                    })
                }}
                name="validateOnly"
                className='row'
                autoComplete="off"
            >
                <Form.Item
                    name="OldPassword"
                    label="舊密碼"
                    rules={[{ required: true }]}
                >
                    <Input.Password />
                </Form.Item>
                <Form.Item
                    name="NewPassword"
                    label="新密碼"
                    rules={[{ required: true }]}
                >
                    <Input.Password />
                </Form.Item>
                <Form.Item
                    name="NewPasswordAgain"
                    label="再次輸入新密碼"
                    rules={[{ required: true }]}
                >
                    <Input.Password />
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
export default Modal<any, any>({
    children: Index,
    title: () => '更換密碼'
})
