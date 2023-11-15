import { Button } from 'antd'
import React, { useEffect, useMemo, useState } from 'react'
import ModalEdit from './components/ModalEdit'
import BanchTags from './components/BanchTags'
import companyBanchTypes from 'types/companyBanch'
import api from 'api/Index'
import { v4 } from 'uuid'
import { useAppSelector } from 'hook/redux'

interface props {
    subComponents: 'tag' | 'table'
    defaultValue?: number[] // banch id []
    onChange?: (v: companyBanchTypes.TABLE[]) => void
}

const Index = ({
    subComponents,
    defaultValue = [],
    onChange = () => {}
}: props): JSX.Element => {
    const selector = useAppSelector((v) => v?.companyBanch.selector)
    const defaultValueFilter = selector?.filter((item) => defaultValue?.includes(item?.BanchId))
    const [selected, setSelected] = useState<companyBanchTypes.TABLE[]>(defaultValueFilter)

    const ModalComponent = useMemo(() => ModalEdit({ id: v4() }), [])

    useEffect(() => {
        void api.companyBanch.getSelector()
    }, [])

    useEffect(() => {
        onChange(selected)
    }, [selected])
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
                style={{ width: '100px' }}
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
