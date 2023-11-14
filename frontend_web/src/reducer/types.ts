export interface action {
    type: any
    payload: any
}

export const loadingType = {
    SET_LOADING: 'SET_LOADING'
}

export const systemType = {
    SET_AUTH: 'SET_AUTH',
    SET_SIDEBAR: 'SET_SIDEBAR',
    SET_FUNC: 'SET_FUNC'
}

export const userType = {
    SET_USER_MINE: 'SET_USER_MINE',
    SET_EMPLOYEE: 'SET_EMPLOYEE'
}

export const companyType = {
    SET_COMPANY_MINE: 'SET_COMPANY_MINE'
}

export const companyBanchType = {
    SET_BANCH_ALL: 'SET_BANCH_ALL',
    SET_BANCH_SELECTOR: 'SET_BANCH_SELECTOR'
}

export const roleType = {
    SET_ROLE_ALL: 'SET_ROLE_ALL',
    SET_ROLE_SINGLE: 'SET_ROLE_SINGLE',
    SET_ROLE_SELECTOR: 'SET_ROLE_SELECTOR'
}
