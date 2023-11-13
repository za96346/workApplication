import { funcCode } from 'types/system'
import { useAppSelector } from './redux'
import { usePermissionProps } from './types'

const usePermission = ({ funcCode }: { funcCode: funcCode }): usePermissionProps.returnType => {
    const permission = useAppSelector((v) => v?.system?.auth?.permission?.[funcCode])

    return {
        isEditable: 'edit' in permission,
        isDeleteable: 'delete' in permission,
        isInquirable: 'inquire' in permission,
        isAddable: 'add' in permission
    }
}
export {
    usePermission,
    type usePermissionProps as usePermissionTypes
}
