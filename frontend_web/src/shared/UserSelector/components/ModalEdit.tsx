import { Table } from 'antd'
import { useAppSelector } from 'hook/redux'
import React, { useEffect, useMemo, useRef } from 'react'
import { Modal } from 'shared/Modal/Index'
import columns from '../method/columns'
import modal from 'shared/Modal/types'
import Btn from 'shared/Button'
import userTypes from 'types/user'
import Searchbar from './Searchbar'
import { RowSelectionType } from 'antd/es/table/interface'

interface modalInfo {
    onSave: (v: any) => void
    defaultValue: number[]
    type: RowSelectionType
}

interface props {
    modalInfo: modal.modalInfoProps<modalInfo>
}

const ModalEdit = ({ modalInfo }: props): JSX.Element => {
    const dataRef = useRef<userTypes.TABLE[]>([])
    const selector = useAppSelector((v) => v?.user?.selector)

    const dataSource = useMemo(() => {
        return selector?.map((item) => ({
            ...item,
            key: item?.UserId,
            Status: (
                <span className='text-danger'>
                    {item?.DeleteFlag === 'Y' ? '已刪除' : ''}
                </span>
            )
        }))
    }, [selector])

    // 被選擇事件
    const onSelect = (v: userTypes.TABLE, isSelected: boolean): void => {
        if (modalInfo?.type === 'radio') {
            if (isSelected) {
                dataRef.current = [v]
            } else {
                dataRef.current = []
            }
        } else {
            if (isSelected) {
                dataRef.current.push(v)
            } else {
                dataRef.current = dataRef.current
                    ?.filter((item) => (
                        item?.UserId !== v?.UserId
                    ))
            }
        }
    }

    // 設定預設值
    useEffect(() => {
        dataRef.current = dataSource?.filter((item) => (
            modalInfo
                ?.defaultValue
                ?.includes(item?.UserId)
        ))
    }, [modalInfo?.defaultValue])
    return (
        <>
            <Searchbar />
            <Table
                dataSource={dataSource}
                columns={columns}
                rowSelection={{
                    type: modalInfo?.type || 'checkbox',
                    defaultSelectedRowKeys: modalInfo?.defaultValue || [],
                    onSelect
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
    children: ModalEdit,
    title: () => '使用者選擇',
    width: (isLess) => isLess('md') ? '100vw' : '500px',
    uid: id
})
