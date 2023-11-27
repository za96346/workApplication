import { Form, FormInstance, Select } from 'antd'
import useRoleBanchUserList from 'hook/useRoleBanchUserList'
import React from 'react'
import Btn from 'shared/Button'
import { Modal } from 'shared/Modal/Index'
import modal from 'shared/Modal/types'
import performanceTypes from 'types/performance'
import { funcCode, operationCode } from 'types/system'

interface modalInfo {
    onSave: (v: FormInstance<any>) => void
    data?: performanceTypes.TABLE
}

interface props {
    modalInfo: modal.modalInfoProps<modalInfo>
}

const ModalChangeBanch = ({ modalInfo }: props): JSX.Element => {
    const { data } = modalInfo
    const roleBanchUserList = useRoleBanchUserList({
        funcCode: funcCode.performance,
        operationCode: operationCode.changeBanch
    })

    const [form] = Form.useForm()

    return (
        <>
            <Form
                name="validateOnly"
                autoComplete="off"
                form={form}
                className="row mt-4"
            >
                <Form.Item
                    name="BanchId"
                    initialValue={data?.BanchId}
                    required
                    className="col-6"
                    label="部門"
                >
                    <Select options={roleBanchUserList.banchSelectList} />
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
    children: ModalChangeBanch,
    title: (v) => '更換績效部門',
    width: (isLess) => isLess('md') ? '100vw' : '500px'
})
