export interface action {
    type: any
    payload: any
}
export const userType = {
    SET_TOKEN: 'SET_TOKEN',
    CLEAR_TOKEN: 'CLEAR_TOKEN',
    CLEAR_USER_ALL: 'CLEAR_USER_ALL'
}
export const companyType = {
    SET_BANCH: 'SET_BANCH',
    SET_EMPLOYEE: 'SET_EMPLOYEE',
    CLEAR_COMPANY_ALL: 'CLEAR_COMPANY_ALL'
}

export const statusType = {
    FETCH_BANCH_ON: 'FETCH_BANCH_ON',
    FETCH_BANCH_OFF: 'FETCH_BANCH_OFF',
    FETCH_USER_ALL_ON: 'FETCH_USER_ALL_ON',
    FETCH_USER_ALL_OFF: 'FETCH_USER_ALL_OFF',
    CLEAR_STATUS_ALL: 'CLEAR_STATUS_ALL'
}
