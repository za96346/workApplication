import Row from './components/Row'
import React from 'react'
import { useAppSelector } from 'hook/redux'
import { v4 } from 'uuid'
import ModalDetail from '../modalDetail/Index'

const Index = (): JSX.Element => {
    const func = useAppSelector((v) => v?.system?.func)
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
                            operationItemArray={func.operationItem}
                        />
                    ))
                }

            </div>
        </>
    )
}
export default Index
