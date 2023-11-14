import { Button } from 'antd'
import React, { useEffect, useMemo } from 'react'
import ModalEdit from './components/ModalEdit'
import RoleTags from './components/RoleTags'
import { v4 } from 'uuid'
import api from 'api/Index'

interface props {
    subComponents: 'tag' | 'table'
}

const Index = ({
    subComponents
}: props): JSX.Element => {
    const ModalComponent = useMemo(() => ModalEdit({ id: v4() }), [])

    useEffect(() => {
        void api.role.getSelector()
    }, [])
    return (
        <>
            <ModalComponent />
            {
                subComponents === 'tag' && (
                    <RoleTags />
                )
            }
            <Button
                onClick={() => {
                    ModalComponent.open({})
                }}
            >
                選擇角色
            </Button>
        </>
    )
}
export default Index
