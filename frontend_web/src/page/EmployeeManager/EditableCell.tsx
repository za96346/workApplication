import { Form } from 'antd'
import React, { ReactNode } from 'react'
import { EmpManagerCellType } from '../../type'

interface EditableCellProps extends React.HTMLAttributes<HTMLElement> {
    editing: boolean
    dataIndex: string
    title: any
    inputType: ReactNode
    record: EmpManagerCellType
    index: number
    children: React.ReactNode
}

const EditableCell: React.FC<EditableCellProps> = ({
    editing,
    dataIndex,
    title,
    inputType,
    record,
    index,
    children,
    ...restProps
}) => {
    return (
        <td {...restProps}>
            {editing
                ? (
                    <>
                        <Form.Item
                            name={dataIndex}
                            style={{ margin: 0 }}
                            rules={[
                                {
                                    required: true,
                                    message: `請輸入${title}!`
                                }
                            ]}
                        >
                            {inputType}
                        </Form.Item>
                        <Form.Item initialValue="" name="BanchId" />
                        <Form.Item initialValue="" name="CompanyCode" />
                        <Form.Item initialValue="" name="PermessionId"/>
                    </>
                )
                : (
                    children
                )}
        </td>
    )
}
export default EditableCell
