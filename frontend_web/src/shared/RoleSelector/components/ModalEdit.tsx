import { Table } from 'antd'
import { useAppSelector } from 'hook/redux'
import React, { useEffect, useMemo, useRef } from 'react'
import { Modal } from 'shared/Modal/Index'
import columns from '../method/columns'
import modal from 'shared/Modal/types'
import Btn from 'shared/Button'

interface modalInfo {
    onSave: (v: any) => void
    defaultValue: number[]
}

interface props {
    modalInfo: modal.modalInfoProps<modalInfo>
}

const RoleSelector = ({ modalInfo }: props): JSX.Element => {
    const dataRef = useRef([])
    const selector = useAppSelector((v) => v?.role?.selector)

    const dataSource = useMemo(() => {
        return selector?.map((item) => ({
            ...item,
            key: item?.RoleId,
            Status: (
                <span className='text-danger'>
                    {item?.DeleteFlag === 'Y' ? '已刪除' : ''}
                </span>
            )
        }))
    }, [selector])

    // 設定預設值
    useEffect(() => {
        dataRef.current = dataSource?.filter((item) => (
            modalInfo
                ?.defaultValue
                ?.includes(item?.RoleId)
        ))
    }, [modalInfo?.defaultValue])
    return (
        <>
            <Table
                dataSource={dataSource}
                columns={columns}
                rowSelection={{
                    type: 'checkbox',
                    defaultSelectedRowKeys: modalInfo?.defaultValue || [],
                    onSelect: (v, isSelected) => {
                        if (isSelected) {
                            dataRef.current.push(v)
                        } else {
                            dataRef.current = dataRef.current
                                ?.filter((item) => (
                                    item?.RoleId !== v?.RoleId
                                ))
                        }
                    }
                }}
            />
            <Modal.Footer>
                {
                    () => (
                        <>
                            <Btn.Cancel
                                onClick={() => {
                                    void modalInfo.onClose()
                                }}
                            />
                            <Btn.Save
                                onClick={() => {
                                    modalInfo.onSave(dataRef.current)
                                }}
                            />
                        </>
                    )
                }
            </Modal.Footer>
        </>
    )
}
export default ({ id }): any => Modal<modalInfo, any>({
    children: RoleSelector,
    title: () => '角色選擇',
    width: (isLess) => isLess('md') ? '100vw' : '500px',
    uid: id
})
