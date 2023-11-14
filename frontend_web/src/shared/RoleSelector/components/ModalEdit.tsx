import { Table } from 'antd'
import api from 'api/Index'
import { useAppSelector } from 'hook/redux'
import React, { useEffect, useMemo } from 'react'
import { Modal } from 'shared/Modal/Index'
import columns from '../method/columns'
import modal from 'shared/Modal/types'
import Btn from 'shared/Button'

interface modalInfo {
    onSave: (v: any) => void
}

interface props {
    modalInfo: modal.modalInfoProps<modalInfo>
}

const RoleSelector = ({ modalInfo }: props): JSX.Element => {
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

    useEffect(() => {
        void api.role.getSelector()
    }, [])
    return (
        <>
            <Table
                dataSource={dataSource}
                columns={columns}
                rowSelection={{
                    type: 'checkbox'
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
                                    // modalInfo.onSave(form)
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
