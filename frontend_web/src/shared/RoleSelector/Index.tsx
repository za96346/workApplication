import { Button } from 'antd'
import React, { useEffect, useMemo, useState } from 'react'
import ModalEdit from './components/ModalEdit'
import RoleTags from './components/RoleTags'
import { v4 } from 'uuid'
import api from 'api/Index'
import roleTypes from 'types/role'
import { useAppSelector } from 'hook/redux'

interface props {
    subComponents: 'tag' | 'table'
    defaultValue?: number[] // role id []
    onChange?: (v: roleTypes.TABLE[]) => void
}

const Index = ({
    subComponents,
    defaultValue = [],
    onChange = () => {}
}: props): JSX.Element => {
    const selector = useAppSelector((v) => v?.role.selector)
    const defaultValueFilter = selector?.filter((item) => defaultValue?.includes(item?.RoleId))
    const [selected, setSelected] = useState<roleTypes.TABLE[]>(defaultValueFilter || [])

    const ModalComponent = useMemo(() => ModalEdit({ id: v4() }), [])

    useEffect(() => {
        void api.role.getSelector()
    }, [])

    useEffect(() => {
        onChange(selected)
    }, [selected])
    return (
        <>
            <ModalComponent />
            {
                subComponents === 'tag' && (
                    <RoleTags
                        selected={selected}
                        setSelected={setSelected}
                    />
                )
            }
            <Button
                style={{ width: '100px' }}
                onClick={() => {
                    ModalComponent.open({
                        defaultValue: selected?.map((i) => i?.RoleId),
                        onSave: (v) => {
                            setSelected(v)
                            ModalComponent.close({})
                        }
                    })
                }}
            >
                選擇角色
            </Button>
        </>
    )
}
export default Index
