import React, { useMemo, useState } from 'react'
import { Table } from 'antd'

import useReduceing from 'Hook/useReducing'
import { column } from './method/column'
import SearchBar from './component/SearchBar'
import Btn from 'Share/Btn'
import statics from 'statics'
import ModalEdit from './component/modal/Index'
import { UserType } from 'type'
import dateHandle from 'method/dateHandle'
// import dateHandle from 'method/dateHandle'

interface modalType {
    open: boolean
    type: string
    value: UserType
}
const EmployeeManager = (): JSX.Element => {
    const { company } = useReduceing()
    const [modal, setModal] = useState<modalType>({
        open: false,
        type: '',
        value: null
    })
    const onClose = (): void => {
        setModal((prev) => ({ ...prev, open: false, value: null }))
    }
    const onOpen = (v: UserType): void => {
        setModal((prev) => ({ ...prev, open: true, value: v }))
    }
    const employee = useMemo(() => {
        return company.employee.map((item) => ({
            ...item,
            Permession: statics.permession[item.Permession],
            OnWorkDay: dateHandle.transferUtcFormat(item.OnWorkDay).substring(0, 10),
            action: (
                <div>
                    <Btn.Edit onClick={() => { onOpen(item) }} />
                </div>
            )
        }))
    }, [company.employee])
    return (
        <>
            <ModalEdit data={modal.value} open={modal.open} onClose={onClose} />
            <SearchBar reSearching={modal.open} />
            <Table
                sticky={{ offsetHeader: -20 }}
                columns={column}
                dataSource={employee}
            />
        </>
    )
}
export default EmployeeManager
