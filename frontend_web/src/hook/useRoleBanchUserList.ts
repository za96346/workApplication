import { funcCode, operationCode } from 'types/system'
import { useAppSelector } from './redux'
import { useMemo } from 'react'
import roleTypes from 'types/role'
import companyBanchTypes from 'types/companyBanch'
import { DefaultOptionType } from 'antd/es/select'
import userTypes from 'types/user'

interface props {
    funcCode: funcCode | 'all'
    operationCode: operationCode | 'all'
}

interface returnType {
    banchList: companyBanchTypes.TABLE[]
    roleList: roleTypes.TABLE[]
    userList: userTypes.TABLE[]
    banchSelectList: DefaultOptionType[]
    roleSelectList: DefaultOptionType[]
    userSelectList: DefaultOptionType[]
    banchObject: Record<companyBanchTypes.TABLE['BanchId'], companyBanchTypes.TABLE>
    roleObject: Record<roleTypes.TABLE['RoleId'], roleTypes.TABLE>
    userObject: Record<userTypes.TABLE['UserId'], userTypes.TABLE>
}

const useRoleBanchUserList = (
    { funcCode, operationCode }: props
): returnType => {
    const roleBanchList = useAppSelector((v) => v?.system?.roleBanchList)

    const {
        availableBanch,
        availableRole,
        availableUser,
        scopeBanch,
        scopeRole,
        scopeUser
    } = roleBanchList

    /* 部門 */
    const banchList = useMemo(() => {
        const banchIdArray = scopeBanch?.[funcCode]?.[operationCode]
        if (funcCode === 'all') {
            return availableBanch
        }
        return availableBanch?.filter((item) => banchIdArray?.includes(item?.BanchId))
    }, [roleBanchList])

    const banchSelectList = useMemo(() => {
        return banchList?.map((item) => ({
            value: item?.BanchId,
            label: item?.BanchName
        }))
    }, [banchList])

    const banchObject = useMemo(() => {
        return availableBanch?.reduce((accr, item) => {
            accr[item?.BanchId] = item
            return accr
        }, {})
    }, [availableBanch])

    /* 角色 */

    const roleList = useMemo(() => {
        const roleIdArray = scopeRole?.[funcCode]?.[operationCode]
        if (funcCode === 'all') {
            return availableRole
        }
        return availableRole?.filter((item) => roleIdArray?.includes(item?.RoleId))
    }, [roleBanchList])

    const roleSelectList = useMemo(() => {
        return roleList?.map((item) => ({
            value: item?.RoleId,
            label: item?.RoleName
        }))
    }, [roleList])

    const roleObject = useMemo(() => {
        return availableRole?.reduce((accr, item) => {
            accr[item?.RoleId] = item
            return accr
        }, {})
    }, [availableRole])

    /* 使用者 */

    const userList = useMemo(() => {
        const userIdArray = scopeUser?.[funcCode]?.[operationCode]
        if (funcCode === 'all') {
            return availableUser
        }
        return availableUser?.filter((item) => userIdArray?.includes(item?.UserId))
    }, [roleBanchList])

    const userSelectList = useMemo(() => {
        return userList?.map((item) => ({
            value: item?.UserId,
            label: item?.UserName
        }))
    }, [userList])

    const userObject = useMemo(() => {
        return availableUser?.reduce((accr, item) => {
            accr[item?.UserId] = item
            return accr
        }, {})
    }, [availableUser])

    return {
        banchList,
        roleList,
        userList,
        roleSelectList,
        banchSelectList,
        userSelectList,

        // array to object
        banchObject,
        roleObject,
        userObject
    }
}
export default useRoleBanchUserList
