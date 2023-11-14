import { Button } from 'antd'
import React, { useEffect, useMemo, useState } from 'react'
import ModalEdit from './components/ModalEdit'
import BanchTags from './components/BanchTags'
import companyBanchTypes from 'types/companyBanch'
import api from 'api/Index'
import { v4 } from 'uuid'

interface props {
    subComponents: 'tag' | 'table'
    defaultValue: companyBanchTypes.TABLE[]
}

const Index = ({
    subComponents,
    defaultValue = []
}: props): JSX.Element => {
    const [selected, setSelected] = useState<companyBanchTypes.TABLE[]>([])

    const ModalComponent = useMemo(() => ModalEdit({ id: v4() }), [])

    useEffect(() => {
        setSelected(defaultValue)
    }, [defaultValue])

    useEffect(() => {
        void api.companyBanch.getSelector()
    }, [])
    return (
        <>
            <ModalComponent />
            {
                subComponents === 'tag' && (
                    <BanchTags
                        selected={selected}
                        setSelected={setSelected}
                    />
                )
            }
            <Button
                onClick={() => {
                    ModalComponent.open({
                        defaultValue: selected?.map((i) => i?.BanchId),
                        onSave: (v) => {
                            setSelected(v)
                            ModalComponent.close({})
                        }
                    })
                }}
            >
                選擇部門
            </Button>
        </>
    )
}
export default Index
