import { Table } from 'antd'
import React, { useEffect, useMemo } from 'react'
import columns from './methods/column'
import Searchbar from './components/Searchbar'
import { useAppSelector } from 'hook/redux'
import ModalPerformance from './components/ModalPeroformance/Index'
import { dropdownList } from './methods/dropdownList'
import Dropdown from 'shared/Dropdown'
import { usePermission } from 'hook/usePermission'
import { FuncCodeEnum } from 'types/system'
import BtnEvent from './methods/BtnEvent'
import YearPerformanceSession from './methods/yearPerformanceSession'

const Index = (): JSX.Element => {
    const performance = useAppSelector((v) => v?.performance?.year)
    const permission = usePermission({ funcCode: FuncCodeEnum.performance })

    const data = useMemo(() => (
        performance?.map((item) => ({
            ...item,
            action: (
                <Dropdown
                    menu={dropdownList(permission)}
                    onSelect={(v) => {
                        BtnEvent({
                            type: v,
                            value: item
                        })
                    }}
                />
            )
        }))
    ), [performance])

    useEffect(() => {
        return () => {
            YearPerformanceSession.Instance.reset({})
        }
    }, [])

    return (
        <YearPerformanceSession.Provider>
            <ModalPerformance />
            <Searchbar/>
            <Table
                dataSource={data}
                columns={columns}
                sticky={{ offsetHeader: -20 }}
                style={{
                    fontSize: '0.5rem'
                    // width: 'fit-content'
                }}
                size='small'
            />
        </YearPerformanceSession.Provider>
    )
}

export default Index
