import { Radio, RadioChangeEvent } from 'antd'
import { useSession } from 'hook/useSession'
import React from 'react'
import BanchSelector from 'shared/BanchSelector/Index'
import RoleSelector from 'shared/RoleSelector/Index'
import systemTypes from 'types/system'

const RadioGroup = (
    { operationItem, functionItem }: {
        operationItem: systemTypes.operationItemTable
        functionItem: systemTypes.functionItemTable
    }): JSX.Element => {
    const { session, setSession } = useSession<systemTypes.auth['permission']>({})

    // 此元件session 的值
    const currentValue = session()
        ?.[functionItem.funcCode]
        ?.[operationItem?.OperationCode]

    // 當前的 session位置
    const setCurrentSession = (v: Partial<systemTypes.permission>): void => {
        setSession((prev) => ({
            ...prev,
            [functionItem.funcCode]: {
                ...(prev?.[functionItem.funcCode] || {}),
                [operationItem?.OperationCode]: {
                    ...(
                        prev
                            ?.[functionItem.funcCode]
                            ?.[operationItem?.OperationCode] ||
                        {}
                    ),
                    ...v
                }
            }
        }))
    }

    // 部門勾選
    const onBanchScopeChange = (e: RadioChangeEvent): void => {
        setCurrentSession({
            scopeBanch: e.target.value === 'customize'
                ? []
                : e.target.value
        })
    }

    // 角色勾選
    const onRoleScopeChange = (e: RadioChangeEvent): void => {
        setCurrentSession({
            scopeRole: e.target.value === 'customize'
                ? []
                : e.target.value
        })
    }
    return (
        <>
            {/* 部門 */}
            <span className='text-secondary'>部門</span>
            <Radio.Group
                onChange={onBanchScopeChange}
                value={
                    Array.isArray(currentValue?.scopeBanch)
                        ? 'customize'
                        : currentValue?.scopeBanch
                }
            >
                <Radio value={'all'}>全部</Radio>
                <Radio value={'self'}>自己</Radio>
                <Radio value={'customize'}>自訂</Radio>
            </Radio.Group>
            {
                Array.isArray(currentValue?.scopeBanch) && (
                    <BanchSelector
                        defaultValue={currentValue?.scopeBanch || []}
                        subComponents='tag'
                        onChange={(v) => {
                            setCurrentSession({
                                scopeBanch: v?.map((item) => item?.BanchId)
                            })
                        }}
                    />
                )
            }

            {/* 角色 */}
            <span className='text-secondary'>角色</span>
            <Radio.Group
                onChange={onRoleScopeChange}
                value={
                    Array.isArray(currentValue?.scopeRole)
                        ? 'customize'
                        : currentValue?.scopeRole
                }
            >
                <Radio value={'all'}>全部</Radio>
                <Radio value={'self'}>自己</Radio>
                <Radio value={'customize'}>自訂</Radio>
            </Radio.Group>
            {
                Array.isArray(currentValue?.scopeRole) && (
                    <RoleSelector
                        defaultValue={currentValue?.scopeRole || []}
                        subComponents='tag'
                        onChange={(v) => {
                            setCurrentSession({
                                scopeRole: v?.map((item) => item?.RoleId)
                            })
                        }}
                    />
                )
            }
        </>
    )
}
export default RadioGroup
