export interface action {
    type: any
    payload: any
}
export const userType = {
    SET_TOKEN: 'SET_TOKEN',
    CLEAR_TOKEN: 'CLEAR_TOKEN'
}
export const companyType = {
    SET_BANCH: 'SET_BANCH',
    SET_EMPLOYEE: 'SET_EMPLOYEE'
}

export const statusType = {
    FETCH_BANCH_ON: 'FETCH_BANCH_ON',
    FETCH_BANCH_OFF: 'FETCH_BANCH_OFF',
    FETCH_USER_ALL_ON: 'FETCH_USER_ALL_ON',
    FETCH_USER_ALL_OFF: 'FETCH_USER_ALL_OFF'
}
