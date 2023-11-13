import { Table } from 'antd'
import React, { useEffect, useMemo } from 'react'
import columns from './methods/column'
// import modalType from 'static'
import api from 'api/Index'
import Searchbar from './components/Searchbar'
import { useAppSelector } from 'hook/redux'
import { v4 } from 'uuid'
import Btn from 'shared/Button'
import Dropdown from 'shared/Dropdown'
import ModalEdit from './components/modalEdit/Index'
import BtnEvent from './methods/BtnEvent'
import { modalType } from 'static'

const Index = (): JSX.Element => {
    const companyBanch = useAppSelector((v) => v?.companyBanch?.all)
    const dataSource = useMemo(() => {
        return companyBanch?.map((item) => ({
            ...item,
            key: v4(),
            action: <Dropdown menu={[]}/>
        }))
    }, [companyBanch])

    useEffect(() => {
        void api.companyBanch.get()
    }, [])
    return (
        <>
            <ModalEdit />
            <Btn.Add
                onClick={() => {
                    BtnEvent({
                        type: modalType.add,
                        value: null,
                        reload: () => { api.companyBanch.get() }
                    })
                }}
            />
            <Searchbar/>
            <Table
                dataSource={dataSource}
                columns={columns}
            />
        </>
    )
}

export default Index
