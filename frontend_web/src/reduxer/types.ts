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
    SET_BANCH_RULE: 'SET_BANCH_RULE',
    SET_COMPANY: 'SET_COMPANY',
    SET_WEEKEND_SETTING: 'SET_WEEKEND_SETTING',
    SET_WAIT_REPLY: 'SET_WAIT_REPLY',
    SET_WORK_TIME: 'SET_WORK_TIME',
    SET_YEAR_PERFORMANCE: 'SET_YEAR_PERFORMANCE',
    CLEAR_COMPANY_ALL: 'CLEAR_COMPANY_ALL',
    SET_PERFORMANCE: 'SET_PERFORMANCE'
}

export const statusType = {
    ENTRY_ON: 'ENTRY_ON',
    ENTRY_OFF: 'ENTRY_OFF',

    FETCH_COMPANY_ON: 'FETCH_COMPANY_ON',
    FETCH_COMPANY_OFF: 'FETCH_COMPANY_OFF',

    UPDATE_COMPANY_ON: 'UPDATE_COMPANY_ON',
    UPDATE_COMPANY_OFF: 'UPDATE_COMPANY_OFF',

    FETCH_BANCH_ON: 'FETCH_BANCH_ON',
    FETCH_BANCH_OFF: 'FETCH_BANCH_OFF',
    UPDATE_BANCH_ON: 'UPDATE_BANCH_ON',
    UPDATE_BANCH_OFF: 'UPDATE_BANCH_OFF',
    CREATE_BANCH_ON: 'CREATE_BANCH_ON',
    CREATE_BANCH_OFF: 'CREATE_BANCH_OFF',
    DELETE_BANCH_ON: 'DELETE_BANCH_ON',
    DELETE_BANCH_OFF: 'DELETE_BANCH_OFF',

    FETCH_USER_ALL_ON: 'FETCH_USER_ALL_ON',
    FETCH_USER_ALL_OFF: 'FETCH_USER_ALL_OFF',
    UPDATE_USER_ON: 'UPDATE_USER_ON',
    UPDATE_USER_OFF: 'UPDATE_USER_OFF',
    CREATE_USER_ON: 'CREATE_USER_ON',
    CREATE_USER_OFF: 'CREATE_USER_OFF',

    FETCH_SELF_DATA_ON: 'FETCH_SELF_INFO_ON',
    FETCH_SELF_DATA_OFF: 'FETCH_SELF_INFO_OFF',
    UPDATE_SELF_DATA_ON: 'UPDATE_SELF_DATA_ON',
    UPDATE_SELF_DATA_OFF: 'UPDATE_SELF_DATA_OFF',

    FETCH_BANCH_STYLE_ON: 'FETCH_BANCH_STYLE_ON',
    FETCH_BANCH_STYLE_OFF: 'FETCH_BANCH_STYLE_OFF',
    UPDATE_BANCH_STYLE_ON: 'UPDATE_BANCH_STYLE_ON',
    UPDATE_BANCH_STYLE_OFF: 'UPDATE_BANCH_STYLE_OFF',
    CREATE_BANCH_STYLE_ON: 'CREATE_BANCH_STYLE_ON',
    CREATE_BANCH_STYLE_OFF: 'CREATE_BANCH_STYLE_OFF',
    DELETE_BANCH_STYLE_ON: 'DELETE_BANCH_STYLE_ON',
    DELETE_BANCH_STYLE_OFF: 'DELETE_BANCH_STYLE_OFF',

    FETCH_BANCH_RULE_ON: 'FETCH_BANCH_RULE_ON',
    FETCH_BANCH_RULE_OFF: 'FETCH_BANCH_RULE_OFF',
    UPDATE_BANCH_RULE_ON: 'UPDATE_BANCH_RULE_ON',
    UPDATE_BANCH_RULE_OFF: 'UPDATE_BANCH_RULE_OFF',
    CREATE_BANCH_RULE_ON: 'CREATE_BANCH_RULE_ON',
    CREATE_BANCH_RULE_OFF: 'CREATE_BANCH_RULE_OFF',
    DELETE_BANCH_RULE_ON: 'DELETE_BANCH_RULE_ON',
    DELETE_BANCH_RULE_OFF: 'DELETE_BANCH_RULE_OFF',

    FETCH_WEEKEND_SETTING_ON: 'FETCH_WEEKEND_SETTING_ON',
    FETCH_WEEKEND_SETTING_OFF: 'FETCH_WEEKEND_SETTING_OFF',
    CREATE_WEEKEND_SETTING_ON: 'CREATE_WEEKEND_SETTING_ON',
    CREATE_WEEKEND_SETTING_OFF: 'CREATE_WEEKEND_SETTING_OFF',
    DELETE_WEEKEND_SETTING_ON: 'DELETE_WEEKEND_SETTING_ON',
    DELETE_WEEKEND_SETTING_OFF: 'DELETE_WEEKEND_SETTING_OFF',

    FETCH_WAIT_REPLY_ON: 'FETCH_WAIT_REPLY_ON',
    FETCH_WAIT_REPLY_OFF: 'FETCH_WAIT_REPLY_OFF',
    CREATE_WAIT_REPLY_ON: 'CREATE_WAIT_REPLY_ON',
    CREATE_WAIT_REPLY_OFF: 'CREATE_WAIT_REPLY_OFF',
    UPDATE_WAIT_REPLY_ON: 'UPDATE_WAIT_REPLY_ON',
    UPDATE_WAIT_REPLY_OFF: 'UPDATE_WAIT_REPLY_OFF',

    FETCH_WORK_TIME_ON: 'FETCH_WORK_TIME_ON',
    FETCH_WORK_TIME_OFF: 'FETCH_WORK_TIME_OFF',
    UPDATE_WORK_TIME_ON: 'UPDATE_WORK_TIME_ON',
    UPDATE_WORK_TIME_OFF: 'UPDATE_WORK_TIME_OFF',
    CREATE_WORK_TIME_ON: 'CREATE_WORK_TIME_ON',
    CREATE_WORK_TIME_OFF: 'CREATE_WORK_TIME_OFF',

    FETCH_PERFORMANCE_ON: 'FETCH_PERFORMANCE_ON',
    FETCH_PERFORMANCE_OFF: 'FETCH_PERFORMANCE_OFF',
    UPDATE_PERFORMANCE_ON: 'UPDATE_PERFORMANCE_ON',
    UPDATE_PERFORMANCE_OFF: 'UPDATE_PERFORMANCE_OFF',
    CREATE_PERFORMANCE_ON: 'CREATE_PERFORMANCE_ON',
    CREATE_PERFORMANCE_OFF: 'CREATE_PERFORMANCE_OFF',
    DELETE_PERFORMANCE_ON: 'DELETE_PERFORMANCE_ON',
    DELETE_PERFORMANCE_OFF: 'DELETE_PERFORMANCE_OFF',
    COPY_PERFORMANCE_ON: 'COPY_PERFORMANCE_ON',
    COPY_PERFORMANCE_OFF: 'COPY_PERFORMANCE_OFF',
    FETCH_YEAR_PERFORMANCE_ON: 'FETCH_YEAR_PERFORMANCE_ON',
    FETCH_YEAR_PERFORMANCE_OFF: 'FETCH_YEAR_PERFORMANCE_OFF',

    FETCH_SHIFT_MONTH_ON: 'FETCH_SHIFT_MONTH_ON',
    FETCH_SHIFT_MONTH_OFF: 'FETCH_SHIFT_MONTH_OFF',

    FETCH_SHIFT_TOTAL_ON: 'FETCH_SHIFT_TOTAL_ON',
    FETCH_SHIFT_TOTAL_OFF: 'FETCH_SHIFT_TOTAL_OFF',

    CLEAR_STATUS_ALL: 'CLEAR_STATUS_ALL'
}

export const shiftEditType = {
    SET_SHIFT: 'SET_SHIFT',
    SET_SHIFT_STATUS: 'SET_SHIFT_STATUS',
    CLEAR_SHIFT_ALL: 'CLEAR_SHIFT_ALL'
}

export const shiftType = {
    SET_SHIFT: 'SET_SHIFT',
    CLEAR_SHIFT_ALL: 'CLEAR_SHIFT_ALL'
}
