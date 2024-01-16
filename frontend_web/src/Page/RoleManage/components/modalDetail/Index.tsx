import React from 'react'
import Btn from 'shared/Button'
import { Modal } from 'shared/Modal/Index'
import { Divider } from 'antd'
import modal from 'shared/Modal/types'

import { useSession } from 'hook/useSession'
import systemTypes from 'types/system'
import RadioGroup from './components/RadioGroup'
import { useAppSelector } from 'hook/redux'

interface modalInfo {
    onSave: (v: any) => void
    functionItem?: systemTypes.functionItemTable
    operationItemArray: systemTypes.operationItemTable[]
}

interface props {
    modalInfo: modal.modalInfoProps<modalInfo>
}

const ModalDetail = ({ modalInfo }: props): JSX.Element => {
    const { session } = useSession<systemTypes.auth['permission']>({})

    const {
        functionItem,
        operationItemArray = []
    } = modalInfo
    const relation = useAppSelector((v) => (
        v.system.func
            ?.functionRoleBanchRelation
            ?.[functionItem?.FuncCode]
    ))

    const findOperationItem = (v: string): systemTypes.operationItemTable => {
        return operationItemArray?.find((item) => item?.OperationCode === v)
    }

    return (
        <>
            <div className='row'>
                {
                    Object.keys(session()?.[functionItem?.FuncCode] || {})
                        ?.map((item) => {
                            const operationItem = findOperationItem(item)
                            return (
                                <React.Fragment key={operationItem.OperationCode}>
                                    <div className='col-5'>
                                        {operationItem?.OperationName}：
                                    </div>
                                    <div className='col-7 d-flex flex-column'>
                                        <RadioGroup
                                            operationItem={operationItem}
                                            functionItem={functionItem}
                                            scopeLimit={relation?.[operationItem?.OperationCode]}
                                        />
                                    </div>
                                    <Divider/>
                                </React.Fragment>
                            )
                        })
                }
            </div>
            <Modal.Footer>
                {
                    () => (
                        <>
                            <Btn.Close
                                onClick={() => {
                                    void modalInfo.onClose()
                                }}
                            />
                        </>
                    )
                }
            </Modal.Footer>
        </>
    )
}
export default Modal<modalInfo, any>({
    children: ModalDetail,
    title: () => '編輯細項',
    width: () => '100vw'
})
