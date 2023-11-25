import { Table } from 'antd'
import React, { useEffect, useMemo } from 'react'
import columns from './methods/column'
import api from 'api/Index'
import Searchbar from './components/Searchbar'
import { useAppSelector } from 'hook/redux'
import { v4 } from 'uuid'
import Btn from 'shared/Button'
import Dropdown from 'shared/Dropdown'
import ModalEdit from './components/modalEdit/Index'
import BtnEvent from './methods/BtnEvent'
import { modalType } from 'static'
import { usePermission } from 'hook/usePermission'
import { funcCode } from 'types/system'
import { dropdownList } from './methods/dropdownList'

const Index = (): JSX.Element => {
    const role = useAppSelector((v) => v?.role?.all)
    const permission = usePermission({ funcCode: funcCode.roleManage })

    const dataSource = useMemo(() => {
        return role?.map((item) => ({
            ...item,
            key: v4(),
            action: (
                <Dropdown
                    menu={dropdownList(permission, item)}
                    onSelect={(v) => {
                        BtnEvent({
                            type: v,
                            value: item
                        })
                    }}
                />
            )
        }))
    }, [role, permission])

    useEffect(() => {
        void api.role.get()
        void api.system.func()
        void api.companyBanch.getSelector()
        void api.role.getSelector()
    }, [])
    return (
        <>
            <ModalEdit />
            {
                permission?.isAddable && (
                    <Btn.Add
                        onClick={() => {
                            BtnEvent({
                                type: modalType.add,
                                value: null
                            })
                        }}
                    />
                )
            }
            <Searchbar/>
            <Table
                dataSource={dataSource}
                columns={columns}
            />
        </>
    )
}

export default Index
