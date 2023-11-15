import { Button } from 'antd'
import React, { useEffect, useMemo, useState } from 'react'
import ModalEdit from './components/ModalEdit'
import RoleTags from './components/RoleTags'
import { v4 } from 'uuid'
import api from 'api/Index'
import roleTypes from 'types/role'

interface props {
    subComponents: 'tag' | 'table'
    defaultValue?: roleTypes.TABLE[]
    onChange?: (v: roleTypes.TABLE[]) => void
}

const Index = ({
    subComponents,
    defaultValue = [],
    onChange = () => {}
}: props): JSX.Element => {
    const [selected, setSelected] = useState<roleTypes.TABLE[]>(defaultValue)

    const ModalComponent = useMemo(() => ModalEdit({ id: v4() }), [])

    useEffect(() => {
        onChange(selected)
    }, [selected])

    useEffect(() => {
        void api.role.getSelector()
    }, [])
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
