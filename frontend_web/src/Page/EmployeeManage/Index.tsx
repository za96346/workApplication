import { Table } from 'antd'
import React, { useEffect, useMemo } from 'react'
import columns from './methods/column'
import api from 'api/Index'
import Searchbar from './components/Searchbar'
import { useAppSelector } from 'hook/redux'
import { v4 } from 'uuid'
import Btn from 'shared/Button'
import Dropdown from 'shared/Dropdown'
import { usePermission } from 'hook/usePermission'
import { FuncCodeEnum, OperationCodeEnum } from 'types/system'
import BtnEvent from './methods/BtnEvent'
import { modalType } from 'static'
import ModalEdit from './components/modalEdit/Index'
import { dropdownList } from './methods/dropdownList'
import useRoleBanchList from 'hook/useRoleBanchUserList'
import useFlag from 'hook/useFlag'
import dayjs from 'dayjs'

const Index = (): JSX.Element => {
    const permission = usePermission({ funcCode: FuncCodeEnum.employeeManage })
    const roleBanchList = useRoleBanchList({
        funcCode: FuncCodeEnum.employeeManage,
        operationCode: OperationCodeEnum.inquire
    })
    const employee = useAppSelector((v) => v?.user?.employee)
    const { flagToDom } = useFlag()

    const dataSource = useMemo(() => {
        return employee?.map((item) => ({
            ...item,
            key: v4(),
            BanchId: roleBanchList.banchObject?.[item?.BanchId]?.BanchName,
            RoleId: roleBanchList.roleObject?.[item?.RoleId]?.RoleName,
            OnWorkDay: dayjs(item?.OnWorkDay).format('YYYY-MM-DD'),
            QuitFlag: flagToDom({
                flag: item.QuitFlag,
                flagNText: '在職',
                flagYText: '離職'
            }),
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
    }, [employee, permission])

    useEffect(() => {
        void api.user.getEmployee({})
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
            <Searchbar />
            <Table
                dataSource={dataSource}
                columns={columns}
            />
        </>
    )
}

export default Index
