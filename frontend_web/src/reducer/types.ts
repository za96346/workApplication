export interface action {
    type: any
    payload: any
}

export const loadingType = {
    SET_LOADING: 'SET_LOADING'
}

export const systemType = {
    SET_AUTH: 'SET_AUTH',
    SET_SIDEBAR: 'SET_SIDEBAR'
}

export const userType = {
    SET_MINE: 'SET_MINE',
    SET_EMPLOYEE: 'SET_EMPLOYEE'
}

export const companyType = {
    SET_MINE: 'SET_MINE'
}

export const companyBanchType = {
    SET_ALL: 'SET_ALL'
}
