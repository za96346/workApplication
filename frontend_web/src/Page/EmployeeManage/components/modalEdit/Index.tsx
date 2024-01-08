import { Button, DatePicker, Form, FormInstance, Input, Select } from 'antd'
import dayjs from 'dayjs'
import useRoleBanchList from 'hook/useRoleBanchUserList'
import React from 'react'
import Btn from 'shared/Button'
import { Modal } from 'shared/Modal/Index'
import modal from 'shared/Modal/types'
import { modalTitle, modalType } from 'static'
import { funcCode, operationCode } from 'types/system'
import userTypes from 'types/user'
import ModelChangePassword from 'Page/SelfData/ModelChangePassword/Index'
import Switch from 'shared/AntdOverwrite/Switch'

interface modalInfo {
    type?: modalType
    onSave: (v: FormInstance<any>) => void
    data?: userTypes.TABLE
}

interface props {
    modalInfo: modal.modalInfoProps<modalInfo>
}

const ModalEdit = ({ modalInfo }: props): JSX.Element => {
    const roleBanchList = useRoleBanchList({
        funcCode: funcCode.employeeManage,
        operationCode: operationCode?.[modalInfo?.type]
    })

    const [form] = Form.useForm()
    return (
        <>
            <ModelChangePassword />
            <Form
                name="validateOnly"
                autoComplete="off"
                className='row'
                form={form}
                onFieldsChange={(_, v) => console.log('all ', v)}
            >
                <Form.Item
                    name="UserName"
                    label="姓名"
                    className='col-md-6'
                    initialValue={modalInfo?.data?.UserName || ''}
                    rules={[{ required: true }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="EmployeeNumber"
                    label="員工編號"
                    className='col-md-6'
                    initialValue={modalInfo?.data?.EmployeeNumber || ''}
                    rules={[{ required: true }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="Account"
                    label="帳號"
                    className='col-md-6'
                    initialValue={modalInfo?.data?.Account || ''}
                    rules={[{ required: true }]}
                >
                    <Input disabled={modalInfo?.type === modalType.edit} />
                </Form.Item>
                {
                    modalInfo?.type === modalType.add && (
                        <Form.Item
                            name="Password"
                            label="密碼"
                            className='col-md-6'
                            initialValue={modalInfo?.data?.Password || ''}
                            rules={[{ required: true }]}
                        >
                            <Input />
                        </Form.Item>
                    )
                }
                <Form.Item
                    name="OnWorkDay"
                    label="到職日"
                    className='col-md-6'
                    initialValue={dayjs(modalInfo?.data?.OnWorkDay || new Date())}
                    rules={[{ required: true }]}
                >
                    <DatePicker />
                </Form.Item>
                <Form.Item
                    name="BanchId"
                    label="部門"
                    className='col-md-6'
                    initialValue={modalInfo?.data?.BanchId || ''}
                    rules={[{ required: true }]}
                >
                    <Select options={roleBanchList?.banchSelectList} />
                </Form.Item>
                <Form.Item
                    name="RoleId"
                    label="角色"
                    className='col-md-6'
                    initialValue={modalInfo?.data?.RoleId || ''}
                    rules={[{ required: true }]}
                >
                    <Select options={roleBanchList?.roleSelectList} />
                </Form.Item>
                <Modal.Footer>
                    {
                        () => (
                            <div className='col-12 d-flex justify-content-between'>

                                <div>
                                    {
                                        modalInfo?.type !== modalType.add && (
                                            <Switch
                                                label="在職狀態"
                                                formInstance={form}
                                                defaultValue={modalInfo?.data?.QuitFlag}
                                                fieldName='QuitFlag'
                                                antdSwitchProps={{
                                                    onChange: (v) => {
                                                        console.log('change => ', v)
                                                        form.setFieldValue('QuitFlag', v ? 'Y' : 'N')
                                                    },
                                                    checkedChildren: '在職',
                                                    unCheckedChildren: '離職'
                                                }}
                                            />
                                        )
                                    }
                                </div>

                                <div>
                                    {
                                        modalInfo?.type !== modalType.add && (
                                            <Button
                                                onClick={() => {
                                                    ModelChangePassword.open({
                                                        UserId: modalInfo?.data?.UserId
                                                    })
                                                }}
                                            >
                                                        更換密碼
                                            </Button>
                                        )
                                    }
                                    <Btn.Cancel onClick={() => { void modalInfo.onClose() }} />

                                    <Btn.Save
                                        onClick={() => {
                                            modalInfo.onSave(form)
                                        }}
                                    />
                                </div>
                            </div>
                        )
                    }
                </Modal.Footer>
            </Form>

        </>
    )
}
export default Modal<modalInfo, any>({
    children: ModalEdit,
    title: (v) => `${modalTitle[v?.type]}員工`,
    width: (isLess) => isLess('md') ? '100vw' : '500px'
})
