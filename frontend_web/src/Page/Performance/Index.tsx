import { Table } from 'antd'
import React, { useEffect, useMemo } from 'react'
import columns from './methods/column'
import Searchbar from './components/Searchbar'
import { useAppSelector } from 'hook/redux'
import { v4 } from 'uuid'
import Btn from 'shared/Button'
import Dropdown from 'shared/Dropdown'
import ModalEdit from './components/modalEdit/Index'
import BtnEvent from './methods/BtnEvent'
import { modalType } from 'static'
import { usePermission } from 'hook/usePermission'
import { FuncCodeEnum } from 'types/system'
import { dropdownList } from './methods/dropdownList'
import ModalChangeBanch from './components/modalChangeBanch/Index'
import PerformanceSession from './methods/performanceSession'

interface propsType {
    userId?: number
    year?: number
}

const Index = ({ userId, year }: propsType): JSX.Element => {
    const performance = useAppSelector((v) => v?.performance?.all)
    const permission = usePermission({ funcCode: FuncCodeEnum.performance })

    const dataSource = useMemo(() => {
        return performance?.map((item) => ({
            ...item,
            key: v4(),
            action: (
                <Dropdown
                    menu={dropdownList(permission, item)}
                    onSelect={(v) => {
                        BtnEvent({
                            type: v,
                            value: item
                        })
                    }}
                />
            )
        }))
    }, [performance, permission])

    useEffect(() => {
        return () => {
            PerformanceSession.Instance.reset({})
        }
    }, [])

    return (
        <PerformanceSession.Provider>
            <ModalEdit />
            <ModalChangeBanch />
            {
                permission?.isAddable && (
                    <Btn.Add
                        onClick={() => {
                            BtnEvent({
                                type: modalType.add,
                                value: null
                            })
                        }}
                    />
                )
            }
            {
                permission?.isPrintable && (
                    <>
                        <Btn.Print
                            text='表單'
                            onClick={() => {
                                window.open(
                                    'performance/print/form',
                                    '績效評核',
                                    'height=800,width=800'
                                )
                            }}
                        />
                        <Btn.Print
                            text='清單'
                            onClick={() => {
                                window.open(
                                    'performance/print/list',
                                    '績效評核',
                                    'height=800,width=800'
                                )
                            }}
                        />
                    </>
                )
            }
            <Searchbar userId={userId} year={year} />
            <Table
                dataSource={dataSource}
                columns={columns}
                sticky={{ offsetHeader: -20 }}
                style={{
                    fontSize: '0.5rem',
                    width: 'fit-content'
                }}
                size='small'
                scroll={{
                    scrollToFirstRowOnChange: true,
                    x: 1500
                }}
            />
        </PerformanceSession.Provider>
    )
}

export default Index
