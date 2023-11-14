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
    const { session, setSession } = useSession({})

    // 此元件session 的值
    const currentValue = session()
        ?.[functionItem.funcCode]
        ?.[operationItem?.OperationCode]

    // 部門勾選
    const onBanchScopeChange = (e: RadioChangeEvent): void => {
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
                    scopeBanch: e.target.value === 'customize'
                        ? []
                        : e.target.value
                }
            }
        }))
    }

    // 角色勾選
    const onRoleScopeChange = (e: RadioChangeEvent): void => {
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
                    scopeRole: e.target.value === 'customize'
                        ? []
                        : e.target.value
                }
            }
        }))
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
                    <BanchSelector defaultValue={[]} subComponents='tag' />
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
                    <RoleSelector subComponents='tag' />
                )
            }
        </>
    )
}
export default RadioGroup
