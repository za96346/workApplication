import { funcCode } from 'types/system'
import { useAppSelector } from './redux'
import { usePermissionProps } from './types'
import { useCallback } from 'react'

const usePermission = ({ funcCode }: { funcCode: funcCode }): usePermissionProps.returnType => {
    const permission = useAppSelector((v) => v?.system?.auth?.permission?.[funcCode])
    const roleBanchList = useAppSelector((v) => v?.system?.roleBanchList)

    const isEditable = useCallback(({ banchId = null, roleId = null }): boolean => {
        let resultFlag = 'edit' in permission
        if (banchId !== null) {
            resultFlag = resultFlag && roleBanchList
                ?.scopeBanch
                ?.[funcCode]
                ?.edit
                ?.includes(banchId)
        }
        if (roleId !== null) {
            resultFlag = resultFlag && roleBanchList
                ?.scopeRole
                ?.[funcCode]
                ?.edit
                ?.includes(roleId)
        }
        return resultFlag
    }, [funcCode, roleBanchList])

    const isDeleteable = useCallback(({ banchId = null, roleId = null }): boolean => {
        let resultFlag = 'delete' in permission
        if (banchId !== null) {
            resultFlag = resultFlag && roleBanchList
                ?.scopeBanch
                ?.[funcCode]
                ?.delete
                ?.includes(banchId)
        }
        if (roleId !== null) {
            resultFlag = resultFlag && roleBanchList
                ?.scopeRole
                ?.[funcCode]
                ?.delete
                ?.includes(roleId)
        }
        return resultFlag
    }, [funcCode, roleBanchList])

    return {
        // 此兩個不需要看 role banch scope
        isInquirable: 'inquire' in permission,
        isAddable: 'add' in permission,

        // 此兩個才需要看role banch scope
        isEditable,
        isDeleteable,

        //
        isPrintable: 'print' in permission,
        isCopyable: 'copy' in permission
    }
}
export {
    usePermission,
    type usePermissionProps as usePermissionTypes
}
