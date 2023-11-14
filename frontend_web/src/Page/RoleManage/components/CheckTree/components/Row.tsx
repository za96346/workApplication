import React, { useMemo, useState } from 'react'
import { Button, Checkbox } from 'antd'
import type { CheckboxChangeEvent } from 'antd/es/checkbox'
import type { CheckboxValueType } from 'antd/es/checkbox/Group'
import systemTypes from 'types/system'
import ModalDetail from '../../modalDetail/Index'
import { useSession } from 'hook/useSession'

const CheckboxGroup = Checkbox.Group

const defaultCheckedList = ['Apple', 'Orange']

const Row = ({
    functionItem,
    operationItemArray = []
}: {
    functionItem: systemTypes.functionItemTable
    operationItemArray: systemTypes.operationItemTable[]
}): JSX.Element => {
    const { setSession } = useSession<systemTypes.auth['permission']>({})
    const [checkedList, setCheckedList] = useState<CheckboxValueType[]>(defaultCheckedList)

    const checkAll = operationItemArray.length === checkedList.length
    const indeterminate = checkedList.length > 0 && checkedList.length < operationItemArray.length

    const option = useMemo(() => {
        return operationItemArray?.map((item) => item?.OperationName)
    }, [operationItemArray])

    // 單一勾選
    const onChange = (list: CheckboxValueType[]): void => {
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
            accr[op.OperationCode] = {
                scopeBanch: 'all',
                scopeRole: 'all'
            }
            return accr
        }, {})

        setSession((prev) => ({
            ...prev,
            [functionItem?.funcCode]: transListName
        }))
        setCheckedList(list)
    }

    // 勾選全部
    const onCheckAllChange = (e: CheckboxChangeEvent): void => {
        console.log('onchange all =>', e.target.checked)
        const data = operationItemArray.reduce((accr, item) => {
            accr[item?.OperationCode] = {
                scopeBanch: 'all',
                scopeRole: 'all'
            }
            return accr
        }, {})
        setSession((prev) => ({
            ...prev,
            [functionItem?.funcCode]: data
        }))
        setCheckedList(e.target.checked ? option : [])
    }

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
        </div>
    )
}

export default Row
