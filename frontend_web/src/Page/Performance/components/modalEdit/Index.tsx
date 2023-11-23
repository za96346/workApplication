import { DatePicker, Form, FormInstance, Input, InputNumber, Select } from 'antd'
import dayjs from 'dayjs'
import useRoleBanchUserList from 'hook/useRoleBanchUserList'
import React, { useMemo } from 'react'
import Btn from 'shared/Button'
import { Modal } from 'shared/Modal/Index'
import modal from 'shared/Modal/types'
import { modalTitle, modalType } from 'static'
import performanceTypes from 'types/performance'
import { funcCode, operationCode } from 'types/system'

interface modalInfo {
    type?: modalType
    onSave: (v: FormInstance<any>) => void
    data?: performanceTypes.TABLE
}

interface props {
    modalInfo: modal.modalInfoProps<modalInfo>
}

const ModalEdit = ({ modalInfo }: props): JSX.Element => {
    const { data, type } = modalInfo
    const roleBanchUserList = useRoleBanchUserList({
        funcCode: funcCode.performance,
        operationCode: operationCode?.[type]
    })

    const [form] = Form.useForm()

    const month = useMemo(() => (
        new Array(12).fill('').map((item, index) => (
            <Select.Option
                value={index + 1}
                key={index + 1}
            >
                {index + 1} 月
            </Select.Option>
        ))
    ), [])

    return (
        <>
            <Form
                name="validateOnly"
                autoComplete="off"
                form={form}
                className="row mt-4"
            >
                <Form.Item
                    name="Year"
                    initialValue={dayjs()}
                    className="col-6"
                    label="年度"
                    required
                >
                    <DatePicker picker='year' placeholder="選擇年度"/>
                </Form.Item>
                <Form.Item
                    name="Month"
                    initialValue={data?.Month}
                    required
                    className="col-6"
                    label="月份"
                >
                    <Select>
                        {month}
                    </Select>
                </Form.Item>
                <Form.Item
                    name="UserId"
                    className="col-md-6"
                    label="姓名"
                    initialValue={data?.UserId}
                    required
                >
                    <Select options={roleBanchUserList?.userSelectList} />
                </Form.Item>
                <Form.Item
                    name="Goal"
                    initialValue={data?.Goal || ''}
                    className=""
                    label="年度目標"
                    required
                >
                    <Input.TextArea autoSize placeholder="輸入年度目標"/>
                </Form.Item>
                <Form.Item
                    name="Attitude"
                    initialValue={data?.Attitude || 0}
                    className="col-4"
                    label="態度"
                >
                    <InputNumber min={0} max={100} placeholder="態度"/>
                </Form.Item>
                <Form.Item
                    name="Efficiency"
                    initialValue={data?.Efficiency || 0}
                    className="col-4"
                    label="效率"
                >
                    <InputNumber min={0} max={100} placeholder="效率"/>
                </Form.Item>
                <Form.Item
                    name="Professional"
                    initialValue={data?.Professional || 0}
                    className="col-4"
                    label="專業"
                >
                    <InputNumber min={0} max={100} placeholder="專業"/>
                </Form.Item>
                <Form.Item
                    name="BeLate"
                    initialValue={data?.BeLate || 0}
                    className="col-6"
                    label="遲到/早退"
                >
                    <InputNumber min={0} max={100} placeholder="次數"/>
                </Form.Item>
                <Form.Item
                    name="DayOffNotOnRule"
                    initialValue={data?.DayOffNotOnRule || 0}
                    className="col-6"
                    label="未依規定請假"
                >
                    <InputNumber min={0} max={100} placeholder="次數"/>
                </Form.Item>
                <Form.Item
                    name="Directions"
                    initialValue={data?.Directions || ''}
                    className=""
                    label="績效描述"
                >
                    <Input.TextArea autoSize placeholder="績效描述"/>
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
    title: (v) => modalTitle[v?.type] + '績效',
    width: (isLess) => isLess('md') ? '100vw' : '500px'
})
