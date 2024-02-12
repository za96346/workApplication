import React, { useEffect, useMemo, useState } from 'react'
import { Button, Checkbox } from 'antd'
import type { CheckboxChangeEvent } from 'antd/es/checkbox'
import type { CheckboxValueType } from 'antd/es/checkbox/Group'
import systemTypes from 'types/system'
import ModalDetail from '../../modalDetail/Index'
import { useSession } from 'hook/useSession'

interface props {
    functionItem: systemTypes.functionItemTable
    operationItemArray: systemTypes.operationItemTable[]
}

const CheckboxGroup = Checkbox.Group

const Row = ({
    functionItem,
    operationItemArray = []
}: props): JSX.Element => {
    const { setSession, session } = useSession<systemTypes.auth['permission']>({})
    const [checkedList, setCheckedList] = useState<CheckboxValueType[]>([])

    const checkAll = operationItemArray.length === checkedList.length && operationItemArray.length > 0
    const indeterminate = checkedList.length > 0 && checkedList.length < operationItemArray.length

    const option = useMemo(() => {
        return operationItemArray?.map((item) => item?.OperationName)
    }, [operationItemArray])

    // 單一勾選
    const onChange = (list: CheckboxValueType[]): void => {
        const currentSession = session()?.[functionItem?.FuncCode] || {}

        /**
         * 轉變 list 中文變英文
         * [func code]: {
         *      [operationCode]: {
         *          scopeBanch:,
         *          scopeRole:,
         *      }
         * }
        */
        const transListName = list.reduce((accr, item) => {
            const op = operationItemArray?.find((i) => i?.OperationName === item)

            // 如果有舊有的資料就拿舊的
            if (op?.OperationCode in currentSession) {
                accr[op.OperationCode] = currentSession[op.OperationCode]
            } else {
                // 預設都是 全部
                accr[op.OperationCode] = {
                    scopeBanch: 'all',
                    scopeRole: 'all'
                }
            }

            return accr
        }, {})

        setSession((prev) => ({
            ...prev,
            [functionItem?.FuncCode]: transListName
        }))
        setCheckedList(list)
    }

    // 勾選全部
    const onCheckAllChange = (e: CheckboxChangeEvent): void => {
        const isCheckAll = e.target.checked
        const data = operationItemArray.reduce((accr, item) => {
            accr[item?.OperationCode] = {
                scopeBanch: 'all',
                scopeRole: 'all'
            }
            return accr
        }, {})
        setSession((prev) => ({
            ...prev,
            [functionItem?.FuncCode]: isCheckAll
                ? data
                : {}
        }))
        setCheckedList(isCheckAll ? option : [])
    }

    // 預設值
    useEffect(() => {
        const operationCode = Object.keys(session()?.[functionItem?.FuncCode] || {})

        const transListName = operationCode.map((item) => {
            const op = operationItemArray?.find((i) => i?.OperationCode === item)
            return op?.OperationName
        })

        setCheckedList(transListName)
    }, [])

    return (
        <div className='row col-12 my-2'>
            <Checkbox
                indeterminate={indeterminate}
                onChange={onCheckAllChange}
                checked={checkAll}
                className='col-4'
            >
                {functionItem?.FuncName}
            </Checkbox>
            <CheckboxGroup
                options={option}
                value={checkedList}
                onChange={onChange}
                className='col-6'
            />
            {
                checkedList?.length > 0 && (
                    <Button
                        className='col-2'
                        onClick={() => {
                            ModalDetail.open({
                                functionItem,
                                operationItemArray
                            })
                        }}
                    >
                        編輯細項
                    </Button>
                )
            }
        </div>
    )
}

export default Row
