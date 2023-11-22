import { funcCode, operationCode } from 'types/system'
import { useAppSelector } from './redux'
import { useMemo } from 'react'
import roleTypes from 'types/role'
import companyBanchTypes from 'types/companyBanch'
import { DefaultOptionType } from 'antd/es/select'

interface props {
    funcCode: funcCode | 'all'
    operationCode: operationCode | 'all'
}

interface returnType {
    banchList: companyBanchTypes.TABLE[]
    roleList: roleTypes.TABLE[]
    banchSelectList: DefaultOptionType[]
    roleSelectList: DefaultOptionType[]
    banchObject: Record<companyBanchTypes.TABLE['BanchId'], companyBanchTypes.TABLE>
    roleObject: Record<roleTypes.TABLE['RoleId'], roleTypes.TABLE>
}

const useRoleBanchList = (
    { funcCode, operationCode }: props
): returnType => {
    const roleBanchList = useAppSelector((v) => v?.system?.roleBanchList)

    const {
        availableBanch,
        availableRole,
        scopeBanch,
        scopeRole
    } = roleBanchList

    const banchList = useMemo(() => {
        const banchIdArray = scopeBanch?.[funcCode]?.[operationCode]
        if (funcCode === 'all') {
            return availableBanch
        }
        return availableBanch?.filter((item) => banchIdArray.includes(item?.BanchId))
    }, [roleBanchList])

    const roleList = useMemo(() => {
        const roleIdArray = scopeRole?.[funcCode]?.[operationCode]
        if (funcCode === 'all') {
            return availableRole
        }
        return availableRole?.filter((item) => roleIdArray.includes(item?.RoleId))
    }, [roleBanchList])

    const banchSelectList = useMemo(() => {
        return banchList?.map((item) => ({
            value: item?.BanchId,
            label: item?.BanchName
        }))
    }, [banchList])

    const roleSelectList = useMemo(() => {
        return roleList?.map((item) => ({
            value: item?.RoleId,
            label: item?.RoleName
        }))
    }, [roleList])

    const banchObject = useMemo(() => {
        return availableBanch?.reduce((accr, item) => {
            accr[item?.BanchId] = item
            return accr
        }, {})
    }, [availableBanch])

    const roleObject = useMemo(() => {
        return availableRole?.reduce((accr, item) => {
            accr[item?.RoleId] = item
            return accr
        }, {})
    }, [availableRole])

    return {
        banchList,
        roleList,
        roleSelectList,
        banchSelectList,

        // array to object
        banchObject,
        roleObject
    }
}
export default useRoleBanchList
