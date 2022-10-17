export interface action {
    type: any
    payload: any
}
export const userType = {
    SET_TOKEN: 'SET_TOKEN',
    SET_SELF_DATA: 'SET_SELF_DATA',
    CLEAR_TOKEN: 'CLEAR_TOKEN',
    CLEAR_USER_ALL: 'CLEAR_USER_ALL'
}
export const companyType = {
    SET_BANCH: 'SET_BANCH',
    SET_EMPLOYEE: 'SET_EMPLOYEE',
    SET_BANCH_STYLE: ' SET_BANCH_STYLE',
    CLEAR_COMPANY_ALL: 'CLEAR_COMPANY_ALL'
}

export const statusType = {
    FETCH_BANCH_ON: 'FETCH_BANCH_ON',
    FETCH_BANCH_OFF: 'FETCH_BANCH_OFF',
    FETCH_USER_ALL_ON: 'FETCH_USER_ALL_ON',
    FETCH_USER_ALL_OFF: 'FETCH_USER_ALL_OFF',
    FETCH_SELF_DATA_ON: 'FETCH_SELF_INFO_ON',
    FETCH_SELF_DATA_OFF: 'FETCH_SELF_INFO_OFF',
    FETCH_BANCH_STYLE_ON: 'FETCH_BANCH_STYLE_ON',
    FETCH_BANCH_STYLE_OFF: 'FETCH_BANCH_STYLE_OFF',
    CLEAR_STATUS_ALL: 'CLEAR_STATUS_ALL'
}
