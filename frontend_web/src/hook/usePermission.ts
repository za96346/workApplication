import { funcCode } from 'types/system'
import { useAppSelector } from './redux'

interface returnType {
    isEditable: boolean
    isDeleteable: boolean
    isInquirable: boolean
    isAddable: boolean
}

const usePermission = ({ funcCode }: { funcCode: funcCode }): returnType => {
    const permission = useAppSelector((v) => v?.system?.auth?.permission?.[funcCode])

    return {
        isEditable: 'edit' in permission,
        isDeleteable: 'delete' in permission,
        isInquirable: 'inquire' in permission,
        isAddable: 'add' in permission
    }
}
export default usePermission
