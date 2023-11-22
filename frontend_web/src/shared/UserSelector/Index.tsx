import { Button } from 'antd'
import React, { useEffect, useMemo, useState } from 'react'
import ModalEdit from './components/ModalEdit'
import RoleTags from './components/UserTags'
import { v4 } from 'uuid'
import api from 'api/Index'
import { useAppSelector } from 'hook/redux'
import userTypes from 'types/user'
import { RowSelectionType } from 'antd/es/table/interface'

interface props {
    subComponents: 'tag' | 'table'
    defaultValue?: number[] // role id []
    onChange?: (v: userTypes.TABLE[]) => void
    type?: RowSelectionType
}

const Index = ({
    subComponents,
    defaultValue = [],
    onChange = () => {},
    type = 'checkbox'
}: props): JSX.Element => {
    const selector = useAppSelector((v) => v?.user.selector)
    const defaultValueFilter = selector?.filter((item) => defaultValue?.includes(item?.UserId))
    const [selected, setSelected] = useState<userTypes.TABLE[]>(defaultValueFilter || [])

    const ModalComponent = useMemo(() => ModalEdit({ id: v4() }), [])

    useEffect(() => {
        void api.user.getSelector({})
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
                        type,
                        defaultValue: selected?.map((i) => i?.RoleId),
                        onSave: (v) => {
                            setSelected(v)
                            ModalComponent.close({})
                        }
                    })
                }}
            >
                選擇員工
            </Button>
        </>
    )
}
export default Index
