import React, { useMemo, useState } from 'react'
import { Button, Table } from 'antd'

import useReduceing from 'Hook/useReducing'
import { column } from './method/column'
import SearchBar from './component/SearchBar'
import Btn from 'Share/Btn'
import statics from 'statics'
import ModalEdit from './component/modal/Index'
import { UserType } from 'type'
import dateHandle from 'method/dateHandle'
import ModalCreate from './component/modalCreate/Index'
// import dateHandle from 'method/dateHandle'

interface modalType {
    open: boolean
    type: string
    value: UserType
}
const EmployeeManager = (): JSX.Element => {
    const { company, user } = useReduceing()
    const [modal, setModal] = useState<modalType>({
        open: false,
        type: '',
        value: null
    })
    // 權限欄位是否可更改
    const permissionEditable = user.selfData?.Permession === 100
    // 部門欄位是否可更改
    const banchEditable = user.selfData?.Permession === 100 ||
        !(user.selfData.UserId === modal.value?.UserId)

    // 關閉事件
    const onClose = (): void => {
        setModal((prev) => ({ ...prev, open: false, value: null }))
    }

    // 打開事件
    const onOpen = (v: UserType, type: string): void => {
        setModal((prev) => ({ ...prev, open: true, value: v, type }))
    }
    const employee = useMemo(() => {
        return company.employee.map((item) => ({
            ...item,
            Permession: statics.permession[item.Permession],
            OnWorkDay: dateHandle.transferUtcFormat(item.OnWorkDay).substring(0, 10),
            action: (
                <div>
                    <Btn.Edit
                        onClick={() => { onOpen(item, statics.type.edit) }}
                    />
                </div>
            )
        }))
    }, [company.employee])
    return (
        <>
            {
                modal.type === statics.type.edit && (
                    <ModalEdit
                        data={modal.value}
                        open={modal.open}
                        onClose={onClose}
                        permissionEditable={permissionEditable}
                        banchEditable={banchEditable}
                    />
                )
            }
            {
                modal.type === statics.type.create && (
                    <ModalCreate
                        permissionEditable={permissionEditable}
                        banchEditable={user.selfData?.Permession === 100}
                        open={modal.open}
                        onClose={onClose}
                    />
                )
            }
            <div className={window.styles.empManagerFilter}>
                <Button
                    onClick={() => {
                        onOpen(null, statics.type.create)
                    }}>
                    新增員工
                </Button>
                <SearchBar reSearching={modal.open} />
            </div>
            <Table
                sticky={{ offsetHeader: -20 }}
                columns={column}
                dataSource={employee}
                scroll={{
                    scrollToFirstRowOnChange: true,
                    x: 1000
                }}
            />
        </>
    )
}
export default EmployeeManager
