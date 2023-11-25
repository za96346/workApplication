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
    const companyBanch = useAppSelector((v) => v?.companyBanch?.all)
    const permission = usePermission({ funcCode: funcCode.banchManage })

    const dataSource = useMemo(() => {
        return companyBanch?.map((item) => ({
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
    }, [companyBanch, permission])

    useEffect(() => {
        void api.companyBanch.get()
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
