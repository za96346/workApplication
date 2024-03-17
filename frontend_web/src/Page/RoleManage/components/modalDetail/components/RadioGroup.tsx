import { Radio, RadioChangeEvent } from 'antd'
import { useSession } from 'hook/useSession'
import React, { useMemo } from 'react'
import BanchSelector from 'shared/BanchSelector/Index'
import RoleSelector from 'shared/RoleSelector/Index'
import { ScopeNameEnum } from 'static'
import systemTypes, { ScopeEnum } from 'types/system'
import UserSelector from 'shared/UserSelector/Index'

const RadioGroup = (
    { operationItem, functionItem, scopeLimit }: {
        operationItem: systemTypes.operationItemTable
        functionItem: systemTypes.functionItemTable
        scopeLimit?: systemTypes.scope
    }): JSX.Element => {
    // scope 預處理 ( 把 empty string in array 的 移除 )
    scopeLimit = useMemo(() => ({
        scopeBanch: scopeLimit?.scopeBanch?.filter((i) => i),
        scopeRole: scopeLimit?.scopeRole?.filter((i) => i),
        scopeUser: scopeLimit?.scopeUser?.filter((i) => i)
    }), [scopeLimit])

    const { session, setSession } = useSession<systemTypes.auth['permission']>({})

    // 此元件session 的值
    const currentValue = session()
        ?.[functionItem.FuncCode]
        ?.[operationItem?.OperationCode]

    // 當前的 session位置
    const setCurrentSession = (v: Partial<systemTypes.permission>): void => {
        setSession((prev) => ({
            ...prev,
            [functionItem.FuncCode]: {
                ...(prev?.[functionItem.FuncCode] || {}),
                [operationItem?.OperationCode]: {
                    ...(
                        prev
                            ?.[functionItem.FuncCode]
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

    // 使用者勾選
    const onUserScopeChange = (e: RadioChangeEvent): void => {
        setCurrentSession({
            scopeUser: e.target.value === 'customize'
                ? []
                : e.target.value
        })
    }

    return (
        <>
            {/* 部門 */}
            {
                scopeLimit.scopeBanch?.length > 0 && (
                    <>
                        <span className='text-secondary'>部門</span>
                        <Radio.Group
                            onChange={onBanchScopeChange}
                            value={
                                Array.isArray(currentValue?.scopeBanch)
                                    ? 'customize'
                                    : currentValue?.scopeBanch
                            }
                        >
                            {
                                scopeLimit?.scopeBanch?.map((item) => (
                                    <Radio key={item} value={item}>{ScopeNameEnum[item]}</Radio>
                                ))
                            }
                        </Radio.Group>
                    </>
                )
            }
            {
                (
                    Array.isArray(currentValue?.scopeBanch) &&
                    scopeLimit?.scopeBanch?.includes(ScopeEnum.customize)
                ) && (
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
            {
                scopeLimit?.scopeRole?.length > 0 && (
                    <>
                        <span className='text-secondary'>角色</span>
                        <Radio.Group
                            onChange={onRoleScopeChange}
                            value={
                                Array.isArray(currentValue?.scopeRole)
                                    ? 'customize'
                                    : currentValue?.scopeRole
                            }
                        >
                            {
                                scopeLimit?.scopeRole?.map((item) => (
                                    <Radio key={item} value={item}>{ScopeNameEnum[item]}</Radio>
                                ))
                            }
                        </Radio.Group>
                    </>
                )
            }

            {
                (
                    Array.isArray(currentValue?.scopeRole) &&
                    scopeLimit?.scopeRole?.includes(ScopeEnum.customize)
                ) && (
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

            {/* 使用者 */}
            {
                scopeLimit?.scopeUser?.length > 0 && (
                    <>
                        <span className='text-secondary'>使用者</span>
                        <Radio.Group
                            onChange={onUserScopeChange}
                            value={
                                Array.isArray(currentValue?.scopeUser)
                                    ? 'customize'
                                    : currentValue?.scopeUser
                            }
                        >
                            {
                                scopeLimit?.scopeUser?.map((item) => (
                                    <Radio key={item} value={item}>{ScopeNameEnum[item]}</Radio>
                                ))
                            }
                        </Radio.Group>
                    </>
                )
            }

            {
                (
                    Array.isArray(currentValue?.scopeUser) &&
                    scopeLimit?.scopeRole?.includes(ScopeEnum.customize)
                ) && (
                    <UserSelector
                        defaultValue={currentValue?.scopeUser || []}
                        subComponents='tag'
                        onChange={(v) => {
                            setCurrentSession({
                                scopeUser: v?.map((item) => item?.UserId)
                            })
                        }}
                    />
                )
            }
        </>
    )
}
export default RadioGroup
