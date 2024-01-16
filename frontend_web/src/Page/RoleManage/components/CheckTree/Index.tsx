import Row from './components/Row'
import React, { useCallback } from 'react'
import { useAppSelector } from 'hook/redux'
import { v4 } from 'uuid'
import ModalDetail from '../modalDetail/Index'
import { FuncCodeEnum } from 'types/system'

const Index = (): JSX.Element => {
    const func = useAppSelector((v) => v?.system?.func)

    // 尋找 operation item
    const getOperationItemArray = useCallback((funcCode: FuncCodeEnum) => {
        const qperationKeyArray = Object.keys(func.functionRoleBanchRelation?.[funcCode] || {})
        return func.operationItem.filter((item) => qperationKeyArray?.includes(item?.OperationCode))
    }, [func])

    return (
        <>
            <ModalDetail />
            <div className="row">
                <div className="col-12 d-flex justify-content-between mt-1">
                    <div>
                        {' '}
                    </div>
                    {/* <div>
                        <Btn.AllNotChecked disabled={disabled} onClick={() => { onAllCheckButtonClick(false) }} />
                        <Btn.AllChecked disabled={disabled} onClick={() => { onAllCheckButtonClick(true) }} />
                    </div> */}
                </div>
                <div className="col-4 checkTreeTittle overflow-x-auto">
                    功能
                </div>
                <div className="col-8 checkTreeTittle overflow-x-auto">
                    選取細項
                </div>
                {
                    func?.functionItem?.map((item) => (
                        <Row
                            key={v4()}
                            functionItem={item}
                            operationItemArray={getOperationItemArray(item?.FuncCode)}
                        />
                    ))
                }

            </div>
        </>
    )
}
export default Index
